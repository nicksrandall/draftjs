package draftjs

import (
	"bytes"
	"strings"
)

const tab = "  "

type wrapperBlockStack struct {
	blocks []*ContentBlock
	length int
}

func (s *wrapperBlockStack) Push(block *ContentBlock) *wrapperBlockStack {
	s.blocks = append(s.blocks, block)
	s.length = s.length + 1
	return s
}

func (s *wrapperBlockStack) CurrentBlock() *ContentBlock {
	if !s.isEmpty() {
		return s.blocks[s.length-1]
	}
	return nil
}

func (s *wrapperBlockStack) Pop() *ContentBlock {
	if block := s.CurrentBlock(); block != nil {
		s.blocks = s.blocks[:s.length-1]
		s.length = s.length - 1
		return block
	}
	return nil
}
func (s *wrapperBlockStack) isEmpty() bool {
	return s.length < 1
}

func renderBlock(contentState *ContentState, BlockIterator *blockIterator, stack *wrapperBlockStack, config *Config, buf *bytes.Buffer) {
	for BlockIterator.block != nil {
		if !stack.isEmpty() {
			if stack.CurrentBlock().Type != BlockIterator.block.Type {
				wrapperBlock := stack.Pop()
				buf.WriteString(strings.Repeat(tab, wrapperBlock.Depth) + getBlockParentAfter(wrapperBlock, config))
				if wrapperBlock.Type == "unordered-list-item" || wrapperBlock.Type == "ordered-list-item" {
					buf.WriteRune('\n')
				}
				if BlockIterator.block.Type == "unordered-list-item" || BlockIterator.block.Type == "ordered-list-item" {
					buf.WriteString(strings.Repeat(tab, stack.length) + getBlockParentBefore(BlockIterator.block, config))
					stack.Push(BlockIterator.block)
				}
			} else if previousBlock := stack.CurrentBlock(); previousBlock.Depth < BlockIterator.block.Depth {
				if BlockIterator.block.Type == "unordered-list-item" || BlockIterator.block.Type == "ordered-list-item" {
					buf.WriteString(strings.Repeat(tab, stack.length) + getBlockParentBefore(previousBlock, config))
					stack.Push(BlockIterator.block)
				}
			}
		} else {
			if BlockIterator.block.Type == "unordered-list-item" || BlockIterator.block.Type == "ordered-list-item" {
				buf.WriteString(strings.Repeat(tab, BlockIterator.block.Depth) + getBlockParentBefore(BlockIterator.block, config))
				stack.Push(BlockIterator.block)
			}
		}
		currentBlock := BlockIterator.block
		buf.WriteString("\n" + strings.Repeat(tab, currentBlock.Depth) + getBlockBefore(currentBlock, config))

		renderInlineStylesAndEntities(contentState, currentBlock, config, buf)
		if BlockIterator.HasNext() && BlockIterator.NextBlock().Depth > currentBlock.Depth {
			renderBlock(contentState, BlockIterator.StepNext(), stack, config, buf)
		}
		buf.WriteString(getBlockAfter(currentBlock, config))

		if BlockIterator.HasNext() && BlockIterator.NextBlock().Depth < currentBlock.Depth {
			if BlockIterator.block.Type == "unordered-list-item" || BlockIterator.block.Type == "ordered-list-item" {
				buf.WriteString(strings.Repeat(tab, BlockIterator.block.Depth) + getBlockParentAfter(BlockIterator.block, config))
				stack.Pop()
			}
			break
		}
		BlockIterator.StepNext()
	}
}

// Render consumes the raw contentState from draftjs and returns the rendered string based on the config passed in
func Render(contentState *ContentState, config *Config) string {
	var stack wrapperBlockStack
	var buf bytes.Buffer

	renderBlock(contentState, newBlockIterator(contentState), &stack, config, &buf)

	if !stack.isEmpty() {
		currentWapperBlock := stack.Pop()
		buf.WriteString(strings.Repeat("  ", currentWapperBlock.Depth) + getBlockParentAfter(currentWapperBlock, config))
	}

	return strings.TrimPrefix(buf.String(), "\n")
}
