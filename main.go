package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/guptarohit/asciigraph"
	"github.com/wcharczuk/go-chart"
	"github.com/yanzay/tbot"
)

var (
	aghSession = os.Getenv("AGH_SESSION")
	botToken   = os.Getenv("BOT_TOKEN")
	chatId     = os.Getenv("CHAT_ID")
)

// Stats represents the data from Adguard as JSON
type Stats struct {
	// Average Processing Time for DNS Queries
	ProcessingTime float64 `json:"avg_processing_time"`
	// When the Filters did the Job. The slice consists of 24 elements(24 hours)
	// Each item represents the Blockings of Each Hour
	BlockedFilter []float64 `json:"blocked_filtering"`
	// Number of DNSQueries received by Adguard.
	//The slice consists of 24 elements(24 hours)
	// Each element is the queries received in an hour
	DNSQueries        []float64            `json:"dns_queries"`
	BlockedFilterNum  float64              `json:"num_blocked_filtering"`
	DNSQueriesNum     float64              `json:"num_dns_queries"`
	PT                float64              `json:"num_replaced_parental"`
	NumSB             float64              `json:"num_replaced_safebrowsing"`
	NumSS             float64              `json:"num_replaced_safesearch"`
	ReplacedPT        []float64            `json:"replaced_parental"`
	ReplacedSB        []float64            `json:"replaced_safebrowsing"`
	TimeUnits         string               `json:"time_units"`
	TopBlockedDomains []map[string]float64 `json:"top_blocked_domains"`
	TopClients        []map[string]float64 `json:"top_clients"`
	TopQueriedDomains []map[string]float64 `json:"top_queried_domains"`
}

func main() {

	url := "http://127.0.0.1/control/stats"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cookie", fmt.Sprintf("agh_session=%s", aghSession))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var stats Stats

	if err := json.Unmarshal(body, &stats); err != nil {
		log.Println(err)
	}
	stats.sendTGMessage()
}
func (s *Stats) generateGraph(tp string) string {
	defer log.Println("Graph Generated")
	graphof := map[string][]float64{
		"DNS": s.DNSQueries,
		"BLK": s.BlockedFilter,
	}
	caption := map[string]string{
		"DNS": "Number of DNS Queries",
		"BLK": "Number of Blocked Queries",
	}
	graph := asciigraph.Plot(graphof[tp], asciigraph.Height(10), asciigraph.Caption(caption[tp]))
	return graph
}

func (s *Stats) sendTGMessage() {
	percent := (s.BlockedFilterNum / s.DNSQueriesNum) * 100
	defer log.Println("Message Sent...")
	message := fmt.Sprintf("Total DNS Queries : %.f\n\nDNS Queries Blocked : %.f\n\n-----\n\nPercent of Queries Blocked: %.2f%%\n\n-----\n\nDNS Query Graph :\n\n`%s`\n\n-----\n\nBlocked Graph:\n\n`%s`\n", s.DNSQueriesNum, s.BlockedFilterNum, percent, s.generateGraph("DNS"), s.generateGraph("BLK"))
	c := tbot.NewClient(botToken, http.DefaultClient, "https://api.telegram.org")
	if _, err := c.SendMessage(chatId, message, tbot.OptParseModeMarkdown); err != nil {
		log.Printf("unable to send message: %v", err)
	}
	s.pieGrph()
	if _, err := c.SendPhotoFile(chatId, "output.png", tbot.OptCaption("PieGraph of Blocked Domains")); err != nil {
		log.Printf("unable to send image: %v", err)
	}
}

func (s *Stats) pieGrph() {
	defer log.Println("Pie Graph Generated and Image Send")
	var chartValues []chart.Value
	for i := range s.TopBlockedDomains {
		for k, v := range s.TopBlockedDomains[i] {
			values := chart.Value{Label: fmt.Sprintf("%.f:%s", v, k), Value: v}
			chartValues = append(chartValues, values)
		}
	}
	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: chartValues,
	}
	f, _ := os.Create("output.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
