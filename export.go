package draftjs

import (
	"bytes"
	"strings"
)

const TAB = "  "

type WrapperBlockStack struct {
	blocks []*ContentBlock
	length int
}

func (s *WrapperBlockStack) Push(block *ContentBlock) *WrapperBlockStack {
	s.blocks = append(s.blocks, block)
	s.length = s.length + 1
	return s
}

func (s *WrapperBlockStack) CurrentBlock() *ContentBlock {
	if !s.isEmpty() {
		return s.blocks[s.length-1]
	}
	return nil
}

func (s *WrapperBlockStack) Pop() *ContentBlock {
	if block := s.CurrentBlock(); block != nil {
		s.blocks = s.blocks[:s.length-1]
		s.length = s.length - 1
		return block
	}
	return nil
}
func (s *WrapperBlockStack) isEmpty() bool {
	return s.length < 1
}

func renderBlock(contentState *ContentState, BlockIterator *BlockIterator, stack *WrapperBlockStack, config *Config, buf *bytes.Buffer) {
	for BlockIterator.block != nil {
		if !stack.isEmpty() {
			if stack.CurrentBlock().Type != BlockIterator.block.Type {
				wrapperBlock := stack.Pop()
				buf.WriteString(strings.Repeat(TAB, wrapperBlock.Depth) + GetBlockParentAfter(wrapperBlock, config))
				if wrapperBlock.Type == "unordered-list-item" || wrapperBlock.Type == "ordered-list-item" {
					buf.WriteRune('\n')
				}
				if BlockIterator.block.Type == "unordered-list-item" || BlockIterator.block.Type == "ordered-list-item" {
					buf.WriteString(strings.Repeat(TAB, stack.length) + GetBlockParentBefore(BlockIterator.block, config))
					stack.Push(BlockIterator.block)
				}
			} else if previousBlock := stack.CurrentBlock(); previousBlock.Depth < BlockIterator.block.Depth {
				if BlockIterator.block.Type == "unordered-list-item" || BlockIterator.block.Type == "ordered-list-item" {
					buf.WriteString(strings.Repeat(TAB, stack.length) + GetBlockParentBefore(previousBlock, config))
					stack.Push(BlockIterator.block)
				}
			}
		} else {
			if BlockIterator.block.Type == "unordered-list-item" || BlockIterator.block.Type == "ordered-list-item" {
				buf.WriteString(strings.Repeat(TAB, BlockIterator.block.Depth) + GetBlockParentBefore(BlockIterator.block, config))
				stack.Push(BlockIterator.block)
			}
		}
		currentBlock := BlockIterator.block
		buf.WriteString("\n" + strings.Repeat(TAB, currentBlock.Depth) + GetBlockBefore(currentBlock, config))

		RenderInlineStylesAndEntities(contentState, currentBlock, config, buf)
		if BlockIterator.HasNext() && BlockIterator.NextBlock().Depth > currentBlock.Depth {
			renderBlock(contentState, BlockIterator.StepNext(), stack, config, buf)
		}
		buf.WriteString(GetBlockAfter(currentBlock, config))

		if BlockIterator.HasNext() && BlockIterator.NextBlock().Depth < currentBlock.Depth {
			if BlockIterator.block.Type == "unordered-list-item" || BlockIterator.block.Type == "ordered-list-item" {
				buf.WriteString(strings.Repeat(TAB, BlockIterator.block.Depth) + GetBlockParentAfter(BlockIterator.block, config))
				stack.Pop()
			}
			break
		}
		BlockIterator.StepNext()
	}
}

func Render(contentState *ContentState, config *Config) string {
	var stack WrapperBlockStack
	var buf bytes.Buffer

	renderBlock(contentState, NewBlockIterator(contentState), &stack, config, &buf)

	if !stack.isEmpty() {
		currentWapperBlock := stack.Pop()
		buf.WriteString(strings.Repeat("  ", currentWapperBlock.Depth) + GetBlockParentAfter(currentWapperBlock, config))
	}

	return strings.TrimPrefix(buf.String(), "\n")
}
