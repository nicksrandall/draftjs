package draftjs

import (
	"encoding/json"
	"testing"
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

		if html != ExpectedResult[i] {
			t.Errorf("\n%s\n", html)
			t.Errorf("\n%s\n", ExpectedResult[i])
		}
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

var ExpectedResult = []string{
	`<h1>This is a header</h1>
<p>This is a {ALERT_NAME} message.</p>
<p>Here is some <strong>bold</strong> text.</p>
<p></p>
<ul>
  <li>Here</li>
  <li>is a</li>
  <li>bullet</li>
</ul>
<blockquote>Thanks!</blockquote>
<p>hi <ins>underline</ins></p>`,
	`<p>a</p>`,
	`<p>asd<strong>f</strong></p>`,
	`<p><strong><em>BoldItalic</em></strong></p>`,
	`<p><em>Bold</em><strong>Italic</strong></p>`,
	`<p><em><a href="/" target="_blank" title="hi">a</a></em></p>`,
	`<p><em><a href="/" target="_blank" title="hi">a</a></em></p>`,
	`<p><em><a href="/" target="_blank">a</a></em></p>`,
	`<p>An ordered list:</p>
<ol>
  <li>One</li>
  <li>Two</li>
</ol>`,
	`<h2>This is a title</h2>
<p>This is a {Amount} message.</p>
<p></p>
<ul>
  <li>Nested</li>
  <li>List  <ul>
    <li>inner</li>
  </ul></li>
  <li>Back  <ul>
    <li>inner     <ul>
      <li>deep</li>
    </ul></li>
    </ul></li>
  <li>Back again</li>
</ul>
<p> <strong>Bold</strong> <em>italic</em> <ins>underine</ins></p>`,
}
