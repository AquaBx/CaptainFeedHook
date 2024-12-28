# CaptainFeedHook

CaptainFeedHook is an RSS feed management application that automatically posts updates from various RSS sources to Discord Webhooks. It’s perfect for keeping your communities informed of the latest news, videos, or articles.

## Features

- **Multi-source support**: Configure multiple RSS feeds with their own Discord Webhooks.
- **Customizable update intervals**: Set the update frequency for each RSS feed.
- **Easy setup**: Uses a simple `config.json` file to manage feeds and webhooks.
- **Lightweight and fast**: Built with Go for optimal performance.

## Requirements

- Go SDK or Docker installed on your machine.
- Valid JSON for the `config.json` configuration file.

## Feed Configuration

In the `config.json` file, add an entry for each RSS feed. Example:

```json
{
  "example": {
    "rss": "https://www.example.com/.rss",
    "webhook": "https://discord.com/api/webhooks/123456789/example",
    "interval": 600
  }
}
```

| Field        | Description                                                                  |
|--------------|------------------------------------------------------------------------------|
| `rss`        | URL of the RSS feed.                                                         |
| `webhook`    | Discord Webhook URL to post the updates.                                     |
| `interval`   | Interval in seconds between update checks.                                   |

## Usage (Go)

1. Clone the repository:
   ```bash
   git clone https://github.com/AquaBx/CaptainFeedHook.git
   cd CaptainFeedHook
   ```

2. Create a `config` folder and add a `config.json` file inside:

    See [Feed Configuration](#feed-configuration) to create this file.

3. Build the application:
    ```bash
    go build
    ```

4. Run the application:
   ```bash
   ./CaptainFeedHook
   ```

## Usage (Docker)

1. Create a `config.json` file in your directory:

    See [Feed Configuration](#feed-configuration) to create this file.

2. Create a Dockerfile:
   ```Dockerfile
   FROM ghcr.io/aquabx/captainfeedhook:latest
   COPY config.json /config/config.json
   ```

3. Build the Docker image:
   ```bash
   docker build -t captainfeedhook .
   ```

4. Run the container:
   ```bash
   docker run -d --name captainfeedhook captainfeedhook
   ```

## Contributing

Contributions are welcome! If you’d like to add a feature or fix a bug, feel free to open an issue or a pull request.

## License

This project is licensed under the GNU GPL. See the [LICENSE](LICENSE) file for more details.