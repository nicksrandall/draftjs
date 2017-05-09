package draftjs

import "fmt"

type EntityDecorator interface {
	RenderBefore(data map[string]interface{}) string
	RenderAfter(data map[string]interface{}) string
}

type LinkDecorator struct{}

func (d *LinkDecorator) RenderBefore(data map[string]interface{}) string {
	if title, ok := data["title"]; ok {
		return fmt.Sprintf("<a href=\"%s\" target=\"_blank\" title=\"%s\">", data["url"], title)
	}
	return fmt.Sprintf("<a href=\"%s\" target=\"_blank\">", data["url"])
}

func (d *LinkDecorator) RenderAfter(data map[string]interface{}) string {
	return "</a>"
}

type PlainTextLinkDecorator struct{}

func (d *PlainTextLinkDecorator) RenderBefore(data map[string]interface{}) string {
	return "["
}

func (d *PlainTextLinkDecorator) RenderAfter(data map[string]interface{}) string {
	return fmt.Sprintf("](%s)", data["url"])
}
