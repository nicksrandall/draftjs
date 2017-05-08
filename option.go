package draftjs

import (
	"html/template"
	"strings"
)

var funcMap = template.FuncMap{
	"repeat": strings.Repeat,
}

type BlockOption struct {
	Before       string
	After        string
	ParentBefore string
	ParentAfter  string
}

// func (d *BlockOption) GetBefore(wr io.Writer, block *ContentBlock) {
// d.lock.RLock()
// if d.templateMap == nil {
// d.lock.RUnlock()
// d.lock.Lock()
// d.templateMap = map[string]*template.Template{}
// d.lock.Unlock()
// d.lock.RLock()
// }
// t, ok := d.templateMap["Before"]
// d.lock.RUnlock()
// if !ok {
// t = template.Must(template.New(d.Type).Funcs(funcMap).Parse(d.Before))
// d.lock.Lock()
// d.templateMap["Before"] = t
// d.lock.Unlock()
// }
// t.Execute(wr, block)
// }
// func (d *BlockOption) GetAfter(wr io.Writer, block *ContentBlock) {
// d.lock.RLock()
// if d.templateMap == nil {
// d.lock.RUnlock()
// d.lock.Lock()
// d.templateMap = map[string]*template.Template{}
// d.lock.Unlock()
// d.lock.RLock()
// }
// t, ok := d.templateMap["After"]
// d.lock.RUnlock()
// if !ok {
// t = template.Must(template.New(d.Type).Funcs(funcMap).Parse(d.After))
// d.lock.Lock()
// d.templateMap["After"] = t
// d.lock.Unlock()
// }
// t.Execute(wr, block)
// }
// func (d *BlockOption) GetParentBefore(wr io.Writer, block *ContentBlock) {
// d.lock.RLock()
// if d.templateMap == nil {
// d.lock.RUnlock()
// d.lock.Lock()
// d.templateMap = map[string]*template.Template{}
// d.lock.Unlock()
// d.lock.RLock()
// }
// t, ok := d.templateMap["ParentBefore"]
// d.lock.RUnlock()
// if !ok {
// t = template.Must(template.New(d.Type).Funcs(funcMap).Parse(d.ParentBefore))
// d.lock.Lock()
// d.templateMap["ParentBefore"] = t
// d.lock.Unlock()
// }
// t.Execute(wr, block)
// }
// func (d *BlockOption) GetParentAfter(wr io.Writer, block *ContentBlock) {
// d.lock.RLock()
// if d.templateMap == nil {
// d.lock.RUnlock()
// d.lock.Lock()
// d.templateMap = map[string]*template.Template{}
// d.lock.Unlock()
// d.lock.RLock()
// }
// t, ok := d.templateMap["ParentAfter"]
// d.lock.RUnlock()
// if !ok {
// t = template.Must(template.New(d.Type).Funcs(funcMap).Parse(d.ParentAfter))
// d.lock.Lock()
// d.templateMap["ParentAfter"] = t
// d.lock.Unlock()
// }
// t.Execute(wr, block)
// }
