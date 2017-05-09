package draftjs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExoportHTML(t *testing.T) {
	blocks := []ContentState{}

	if err := json.Unmarshal([]byte(TestString), &blocks); err != nil {
		t.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := NewHTMLConfig()
	for i, block := range blocks {
		html := Render(&block, config)
		assert.Equal(t, ExpectedHTML[i], html, "Results should be equal")
	}
}

func TestExoportPlainText(t *testing.T) {
	blocks := []ContentState{}

	if err := json.Unmarshal([]byte(TestString), &blocks); err != nil {
		t.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := NewPlainTextConfig()
	for i, block := range blocks {
		text := Render(&block, config)
		assert.Equal(t, ExpectedPlainText[i], text, "Results should be equal")
	}
}

func TestExoportMarkdown(t *testing.T) {
	blocks := []ContentState{}

	if err := json.Unmarshal([]byte(TestString), &blocks); err != nil {
		t.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := NewMarkdownConfig()
	for i, block := range blocks {
		markdown := Render(&block, config)
		assert.Equal(t, ExpectedMarkdown[i], markdown, "Results should be equal")
	}
}

const TestString = `[
	{"entityMap":{"0":{"type":"VARIABLE","mutability":"IMMUTABLE","data":{"name":"Alert name","value":"ALERT_NAME","meta":{"type":"STRING"},"options":[]}}},"blocks":[{"key":"4fv7b","text":"This is a header","type":"header-one","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"9k238","text":"This is a {ALERT_NAME} message.","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[{"offset":10,"length":12,"key":0}],"data":{}},{"key":"71hbn","text":"Here is some bold text.","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":13,"length":4,"style":"BOLD"}],"entityRanges":[],"data":{}},{"key":"30ngu","text":"","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"e18bk","text":"Here","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"96q01","text":"is a","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"cqce8","text":"bullet","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"ae3au","text":"Thanks!","type":"blockquote","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"1m210","text":"hi underline","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":3,"length":9,"style":"UNDERLINE"}],"entityRanges":[],"data":{}}]},
	{"entityMap":{},"blocks":[{"key":"33nh8","text":"a","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[]}]},
	{"entityMap":{},"blocks":[{"key":"99n0j","text":"asdf","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":3,"length":1,"style":"BOLD"}],"entityRanges":[]}]},
	{"entityMap":{},"blocks":[{"key":"9nc73","text":"BoldItalic","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":10,"style":"BOLD"},{"offset":0,"length":10,"style":"ITALIC"}],"entityRanges":[]}]},
	{"entityMap":{},"blocks":[{"key":"9nc73","text":"BoldItalic","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":4,"length":6,"style":"BOLD"},{"offset":0,"length":4,"style":"ITALIC"}],"entityRanges":[]}]},
	{"entityMap":{"0":{"type":"LINK","mutability":"MUTABLE","data":{"url":"/","rel":null,"title":"hi","extra":"foo"}}},"blocks":[{"key":"8r91j","text":"a","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":1,"style":"ITALIC"}],"entityRanges":[{"offset":0,"length":1,"key":0}]}]},
	{"entityMap":{"0":{"type":"LINK","mutability":"MUTABLE","data":{"url":"/","rel":null,"title":"hi","extra":"foo","data-id":42,"data-mutability":"mutable","data-False":"bad","data-":"no"}}},"blocks":[{"key":"8r91j","text":"a","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":1,"style":"ITALIC"}],"entityRanges":[{"offset":0,"length":1,"key":0}]}]},
	{"entityMap":{"0":{"type":"LINK","mutability":"MUTABLE","data":{"url":"/"}}},"blocks":[{"key":"8r91j","text":"a","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":0,"length":1,"style":"ITALIC"}],"entityRanges":[{"offset":0,"length":1,"key":0}]}]},
	{"entityMap":{},"blocks":[{"key":"33nh8","text":"An ordered list:","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[]},{"key":"8kinl","text":"One","type":"ordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[]},{"key":"ekll4","text":"Two","type":"ordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[]}]},
	{"entityMap":{"0":{"type":"VARIABLE","mutability":"IMMUTABLE","data":{"name":"Amount","value":"Amount","meta":{"type":"NUMBER"},"options":[]}}},"blocks":[{"key":"8e2a6","text":"This is a title","type":"header-two","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"7sfk","text":"This is a {Amount} message.","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[{"offset":10,"length":8,"key":0}],"data":{}},{"key":"3l9si","text":"","type":"unstyled","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"akhhp","text":"Nested","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"ek85q","text":"List","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"b77b2","text":"inner","type":"unordered-list-item","depth":1,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"cejl4","text":"Back","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"4mjkg","text":"inner ","type":"unordered-list-item","depth":1,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"dace2","text":"deep","type":"unordered-list-item","depth":2,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"594kf","text":"Back again","type":"unordered-list-item","depth":0,"inlineStyleRanges":[],"entityRanges":[],"data":{}},{"key":"3grre","text":" Bold italic underine","type":"unstyled","depth":0,"inlineStyleRanges":[{"offset":1,"length":4,"style":"BOLD"},{"offset":6,"length":6,"style":"ITALIC"},{"offset":13,"length":8,"style":"UNDERLINE"}],"entityRanges":[],"data":{}}]}
]`

var ExpectedHTML = []string{
	"<h1>This is a header</h1>\n<p>This is a {ALERT_NAME} message.</p>\n<p>Here is some <strong>bold</strong> text.</p>\n<p></p><ul>\n<li>Here</li>\n<li>is a</li>\n<li>bullet</li></ul>\n\n<blockquote>Thanks!</blockquote>\n<p>hi <ins>underline</ins></p>",
	"<p>a</p>",
	"<p>asd<strong>f</strong></p>",
	"<p><strong><em>BoldItalic</em></strong></p>",
	"<p><em>Bold</em><strong>Italic</strong></p>",
	`<p><em><a href="/" target="_blank" title="hi">a</a></em></p>`,
	`<p><em><a href="/" target="_blank" title="hi">a</a></em></p>`,
	`<p><em><a href="/" target="_blank">a</a></em></p>`,
	"<p>An ordered list:</p><ol>\n<li>One</li>\n<li>Two</li></ol>",
	"<h2>This is a title</h2>\n<p>This is a {Amount} message.</p>\n<p></p><ul>\n<li>Nested</li>\n<li>List  <ul>\n  <li>inner</li>  </ul></li>\n<li>Back  <ul>\n  <li>inner     <ul>\n    <li>deep</li>    </ul></li>    </ul></li>\n<li>Back again</li></ul>\n\n<p> <strong>Bold</strong> <em>italic</em> <ins>underine</ins></p>",
}
var ExpectedPlainText = []string{
	`This is a header
This is a {ALERT_NAME} message.
Here is some *bold* text.

- Here
- is a
- bullet

| Thanks!
hi underline`,
	`a`,
	`asd*f*`,
	`*BoldItalic*`,
	`Bold*Italic*`,
	`[a](/)`,
	`[a](/)`,
	`[a](/)`,
	`An ordered list:
- One
- Two`,
	`This is a title
This is a {Amount} message.

- Nested
- List  
  - inner  
- Back  
  - inner     
    - deep        
- Back again

 *Bold* italic underine`,
}
var ExpectedMarkdown = []string{
	`# This is a header
This is a {ALERT_NAME} message.
Here is some **bold** text.

- Here
- is a
- bullet

> Thanks!
hi underline`,
	`a`,
	`asd**f**`,
	`**_BoldItalic_**`,
	`_Bold_**Italic**`,
	`_[a](/)_`,
	`_[a](/)_`,
	`_[a](/)_`,
	`An ordered list:
1. One
1. Two`,
	`## This is a title
This is a {Amount} message.

- Nested
- List  
  - inner  
- Back  
  - inner     
    - deep        
- Back again

 **Bold** _italic_ underine`,
}
