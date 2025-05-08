# KBot

A Telegram bot that provides Lorem Ipsum functionality and basic interaction capabilities.

## Features

- Telegram bot integration
- Basic command handling
- Counter functionality
- Simple message responses

## Prerequisites

- Go 1.24.2 or higher
- Telegram Bot Token (obtained from [@BotFather](https://t.me/botfather))

## Installation

1. Clone the repository:

```bash
git clone https://github.com/i474232898/kbot.git
cd kbot
```

2. Install dependencies:

```bash
go mod download
```

3. Set up your Telegram bot token as an environment variable:

```bash
export TELE_TOKEN="your_telegram_bot_token"
```

## Usage

1. Start the bot:

```bash
go run main.go kbot
```

2. Interact with the bot on Telegram:
   - Find the bot at [@kbotlorem_bot](https://t.me/kbotlorem_bot)
   - Send `/s hello` to get a response with a counter

## Available Commands

- `/s hello` - Returns a greeting message with a counter

## License

This project is open source and available under the MIT License.
