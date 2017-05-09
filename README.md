# draftjs
[![GoDoc](https://godoc.org/github.com/nicksrandall/draftjs?status.svg)](https://godoc.org/github.com/nicksrandall/draftjs)
[![Build Status](https://travis-ci.org/nicksrandall/draftjs.svg?branch=master)](https://travis-ci.org/nicksrandall/draftjs)

This package can take a [raw](https://draftjs.org/docs/api-reference-data-conversion.html#converttoraw) [draftjs](https://draftjs.org) [contentState](https://draftjs.org/docs/api-reference-content-state.html#content) and export it as HTML, Markdown, and Plain Text.

## Usage

```go
func Export(rawContentState []byte) (string, error) {
  block := draftjs.ContentState{}

  if err := json.Unmarshall(rawContentState, &block); err != nil {
    return "", err
  }

  config := draftjs.NewHTMLConfig() // Export HTML
  // config := draftjs.NewMarkdownConfig() // Export Markdown
  // config := draftjs.NewPlainTextConfig() // Export Plain Text

  html := draftjs.Render(&block, config)
  return html, nil
}

```
