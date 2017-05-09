package draftjs

import (
	"bytes"
	"html/template"
	"sort"
	"strconv"
	"unicode/utf8"
)

func getBlockBefore(block *ContentBlock, config *Config) string {
	if option, ok := config.blockRenderers[block.Type]; ok {
		return option.Before
	}
	return ""
}

func getBlockAfter(block *ContentBlock, config *Config) string {
	if option, ok := config.blockRenderers[block.Type]; ok {
		return option.After
	}
	return ""
}

func getBlockParentBefore(block *ContentBlock, config *Config) string {
	if option, ok := config.blockRenderers[block.Type]; ok {
		return option.ParentBefore
	}
	return ""
}

func getBlockParentAfter(block *ContentBlock, config *Config) string {
	if option, ok := config.blockRenderers[block.Type]; ok {
		return option.ParentAfter
	}
	return ""
}

func getStyleBefore(style *InlineStyleRange, config *Config) string {
	if option, ok := config.styleRenderers[style.Style]; ok {
		return option.Before
	}
	return ""
}

func getStyleAfter(style *InlineStyleRange, config *Config) string {
	if option, ok := config.styleRenderers[style.Style]; ok {
		return option.After
	}
	return ""
}

func getEntityBefore(entity *Entity, config *Config) string {
	if entity == nil {
		return ""
	}
	if decorator, ok := config.entityRenderers[entity.Type]; ok {
		return decorator.RenderBefore(entity.Data)
	}
	return ""
}

func getEntityAfter(entity *Entity, config *Config) string {
	if entity == nil {
		return ""
	}
	if decorator, ok := config.entityRenderers[entity.Type]; ok {
		return decorator.RenderAfter(entity.Data)
	}
	return ""
}

func getEntity(contentState *ContentState, entityRange *EntityRange) *Entity {
	if entity, ok := contentState.EntityMap[strconv.Itoa(entityRange.Key)]; ok {
		return entity
	}
	return nil
}

func substring(s string, start int, end int) string {
	start_str_idx := 0
	i := 0
	for j := range s {
		if i == start {
			start_str_idx = j
		}
		if i == end {
			return s[start_str_idx:j]
		}
		i++
	}
	return s[start_str_idx:]
}

func renderInlineStylesAndEntities(content *ContentState, block *ContentBlock, config *Config, buf *bytes.Buffer) {
	ranges, noStyles := getRanges(block)
	if noStyles {
		buf.WriteString(template.HTMLEscapeString(block.Text))
		return
	}

	for _, rng := range ranges {
		styles := GetStyleForRange(rng, block)
		entities := getEntityForRange(rng, block)
		for i := 0; i < len(styles); i++ {
			buf.WriteString(getStyleBefore(styles[i], config))
		}
		for i := 0; i < len(entities); i++ {
			buf.WriteString(getEntityBefore(getEntity(content, entities[i]), config))
		}
		buf.WriteString(template.HTMLEscapeString(substring(block.Text, rng.Offset, rng.Offset+rng.Length)))
		for i := len(entities) - 1; i >= 0; i-- {
			buf.WriteString(getEntityAfter(getEntity(content, entities[i]), config))
		}
		for i := len(styles) - 1; i >= 0; i-- {
			buf.WriteString(getStyleAfter(styles[i], config))
		}
	}

}

func getEntityForRange(r *Range, block *ContentBlock) []*EntityRange {
	if block.EntityRanges == nil || len(block.EntityRanges) == 0 {
		return nil
	}
	res := make([]*EntityRange, 0, 0)
	for _, entityRange := range block.EntityRanges {
		if r.Offset >= entityRange.Offset && r.Offset+r.Length <= entityRange.Offset+entityRange.Length {
			res = append(res, entityRange)
		}
	}
	return res
}

func GetStyleForRange(r *Range, block *ContentBlock) []*InlineStyleRange {

	if block.InlineStyleRanges == nil || len(block.InlineStyleRanges) == 0 {
		return nil
	}
	res := make([]*InlineStyleRange, 0, 0)
	for _, styleRange := range block.InlineStyleRanges {
		if r.Offset >= styleRange.Offset && r.Offset+r.Length <= styleRange.Offset+styleRange.Length {
			res = append(res, styleRange)
		}
	}
	return res
}

// bool == fullstring (no styles)
func getRanges(block *ContentBlock) ([]*Range, bool) {
	if len(block.InlineStyleRanges)+len(block.EntityRanges) == 0 {
		return nil, true
	}

	breakPoints, runeCount := getBreakPoints(block)
	prev := 0
	res := make([]*Range, 0, 0)
	var lastRange *Range
	for _, v := range breakPoints {
		if v == prev {
			continue
		}
		t := new(Range)
		t.Offset = prev
		t.Length = v - prev
		prev = v
		res = append(res, t)
		lastRange = t
	}
	if lastRange != nil {
		if lastRange.Length+lastRange.Offset < runeCount {
			t := new(Range)
			t.Offset = lastRange.Offset + lastRange.Length
			t.Length = utf8.RuneCountInString(block.Text) - t.Offset
			res = append(res, t)
		}
	}
	return res, false
}

func getBreakPoints(block *ContentBlock) ([]int, int) {
	runeCount := utf8.RuneCountInString(block.Text)
	breakPoints := make([]int, runeCount+1, runeCount+1)

	inArray := func(v int, arr []int) bool {
		for i := len(arr) - 1; i >= 0; i-- {
			if v == arr[i] {
				return true
			}
		}
		return false
	}

	ranges := make([]*Range, 0, len(block.InlineStyleRanges)+len(block.EntityRanges))
	for _, styleRange := range block.InlineStyleRanges {
		ranges = append(ranges, &styleRange.Range)
	}
	for _, entityRange := range block.EntityRanges {
		ranges = append(ranges, &entityRange.Range)
	}

	breakPointsCount := 0
	for _, styleRange := range ranges {
		if !inArray(styleRange.Offset, breakPoints[:breakPointsCount]) {
			breakPoints[breakPointsCount] = styleRange.Offset
			breakPointsCount++
		}
		if !inArray(styleRange.Offset+styleRange.Length, breakPoints[:breakPointsCount]) {
			breakPoints[breakPointsCount] = styleRange.Offset + styleRange.Length
			breakPointsCount++
		}
	}

	breakPoints = breakPoints[:breakPointsCount]
	sort.Ints(breakPoints)

	return breakPoints, runeCount
}
