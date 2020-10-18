# Adguard Telegram Bot

This telegram bot sends you a message on Telegram with your daily statistics as an Ascii Graph and a Pie Chart with your Blocked Domains.

## Prerequisites

- Go
- AdGuard Home Installed
- Adguard Home's `AGH_SESSION_TOKEN`. You can get that via this way

![AGH_Session-GIF](https://file.coffee/u/nuAQBNHR_M.gif)

- `Telegram Bot Token` and `Chat ID`, for getting them refer [Telegram Bot Documnetation](https://core.telegram.org/bots#6-botfather)

- **Save these as environment variables with keys, `AGH_SESSION` for for agh_session_token, `BOT_TOKEN` and `CHAT_ID` for telegram bot token and Telegram chat id respectively.**

## Using It

You can clone the repo and generate a binary by `$ go build main.go`. 


## Sample Output

```text
Total DNS Queries: 5273

DNS Queries Blocked: 927

-----

Per cent of Queries Blocked: 17.58%

-----

DNS Query Graph :

 888 ┼╭╮
 799 ┤││               ╭╮   ╭
 710 ┤││               ││   │
 622 ┤│╰╮             ╭╯│   │
 533 ┤│ │ ╭╮          │ │   │
 444 ┤│ │ ││          │ │   │
 355 ┤│ │ │╰╮        ╭╯ │   │
 266 ┤│ │ │ │        │  │   │
 178 ┤│ │╭╯ │        │  │   │
  89 ┤│ ╰╯  │        │  │   │
   0 ┼╯     ╰────────╯  ╰───╯

    Number of DNS Queries

-----

Blocked Graph:

 213 ┼                      ╭
 192 ┤                      │
 170 ┤                 ╭╮   │
 149 ┤╭╮   ╭╮          ││   │
 128 ┤││   ││          ││   │
 106 ┤││   ││          ││   │
  85 ┤││  ╭╯│         ╭╯│   │
  64 ┤││  │ │         │ │   │
  43 ┤│╰╮ │ │         │ │   │
  21 ┤│ │ │ │        ╭╯ │   │
   0 ┼╯ ╰─╯ ╰────────╯  ╰───╯

    Number of Blocked Queries

```

![PieGraph](https://file.coffee/u/9ItgsPR2Xu.jpeg)

## License

[MIT](/LICENSE)