# github-status-notifier

Post GitHub commit status event to Slack.

# Description

github-status-notifier is simple web application that receives webhooks from GitHub.

# Usage

1. Prepare config.json (see also Configuration section)
2. Launch github-status-notifier
3. Add a webhook to your repository
  * Set `content type` to `application/json`
  * Activate a `Status` event

```
github-status-notifier --port 8000 --config config.json
```

# Configuration

config.json:

```json
{
  "Notifications": {
    "Slack": {
      "WebhookURL": "https://hooks.slack.com/services/...",
      "Channel": "#general",
      "Username": "github",
      "IconEmoji": ":robot:"
    }
  }
}
```
