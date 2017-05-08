package draftjs

type Config struct {
	blockRenderers  map[string]*BlockOption
	styleRenderers  map[string]*BlockOption
	entityRenderers map[string]EntityDecorator
}

func NewHTMLConfig() *Config {
	blockRenderers := map[string]*BlockOption{
		"header-one":          {Before: "<h1>", After: "</h1>"},
		"header-two":          {Before: "<h2>", After: "</h2>"},
		"header-three":        {Before: "<h3>", After: "</h3>"},
		"header-four":         {Before: "<h4>", After: "</h4>"},
		"header-five":         {Before: "<h5>", After: "</h5>"},
		"header-six":          {Before: "<h6>", After: "</h6>"},
		"unstyled":            {Before: "<p>", After: "</p>"},
		"code-block":          {Before: "<pre>", After: "</pre>"},
		"blockquote":          {Before: "<blockquote>", After: "</blockquote>"},
		"atomic":              {Before: "<figure>", After: "</figure>"},
		"unordered-list-item": {Before: "<li>", After: "</li>", ParentBefore: "<ul>", ParentAfter: "</ul>"},
		"ordered-list-item":   {Before: "<li>", After: "</li>", ParentBefore: "<ol>", ParentAfter: "</ol>"},
	}
	styleRenderers := map[string]*BlockOption{
		"BOLD":          {Before: "<strong>", After: "</strong>"},
		"UNDERLINE":     {Before: "<ins>", After: "</ins>"},
		"ITALIC":        {Before: "<em>", After: "</em>"},
		"CODE":          {Before: "<code>", After: "</code>"},
		"STRIKETHROUGH": {Before: "<del>", After: "</del>"},
	}
	entityRenderers := map[string]EntityDecorator{
		"LINK": &LinkDecorator{},
	}
	c := Config{blockRenderers, styleRenderers, entityRenderers}
	return &c
}
