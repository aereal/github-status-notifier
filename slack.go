package main

import (
	"fmt"
	"strings"
)

type SlackIncomingWebhookPost struct {
	Text        string            `json:"text"`
	IconEmoji   string            `json:"icon_emoji"`
	Username    string            `json:"username"`
	Channel     string            `json:"channel"`
	Attachments []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Color     string `json:"color"`
	Title     string `json:"title"`
	TitleLink string `json:"title_link"`
	Fallback  string `json:"fallback"`
}

func GitHubStatusEventAsPost(event GitHubStatusEvent) SlackIncomingWebhookPost {
	title := fmt.Sprintf("%s: %s", event.Context, strings.ToTitle(event.State))
	body := fmt.Sprintf("%s at %s", event.Description, event.Repository.Name)
	attachment := SlackAttachment{
		Color:     colorForState(event.State),
		Title:     title,
		TitleLink: event.TargetURL,
		Fallback:  body,
	}
	post := SlackIncomingWebhookPost{
		Text: body,
		Attachments: []SlackAttachment{
			attachment,
		},
	}
	return post
}

func colorForState(state string) string {
	switch state {
	case "pending":
		return "#cea61b"
	case "success":
		return "#55a532"
	case "failure", "error":
		return "#bd2c00"
	default:
		return "#cccccc"
	}
}
