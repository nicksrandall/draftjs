package draftjs

import "fmt"

// EntityDecorator is the interface used to configure how to render a custom entity
type EntityDecorator interface {
	// RenderBefore will render returned string before entity text.
	RenderBefore(data map[string]interface{}) string
	// RenderAfter will render returned string after entity text.
	RenderAfter(data map[string]interface{}) string
}

// LinkDecorator implements EntityDecorator and can be used to render an LINK entity as HTML
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

// PlainTextLinkDecorator implements EntityDecorator and can be used to render an LINK entity as Markdown (or plain text)
type PlainTextLinkDecorator struct{}

func (d *PlainTextLinkDecorator) RenderBefore(data map[string]interface{}) string {
	return "["
}

func (d *PlainTextLinkDecorator) RenderAfter(data map[string]interface{}) string {
	return fmt.Sprintf("](%s)", data["url"])
}
