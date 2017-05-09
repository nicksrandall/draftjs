package draftjs

type blockIterator struct {
	block        *ContentBlock
	index        int
	contentState *ContentState
}

func newBlockIterator(contentState *ContentState) *blockIterator {
	bi := new(blockIterator)
	bi.contentState = contentState
	bi.index = 0
	if len(contentState.Blocks) > 0 {
		bi.block = contentState.Blocks[0]
	}
	return bi
}

func (bi *blockIterator) HasNext() bool {
	return len(bi.contentState.Blocks) != 0 && bi.index+1 < len(bi.contentState.Blocks)
}

func (bi *blockIterator) StepNext() *blockIterator {
	if bi.HasNext() {
		bi.index++
		bi.block = bi.contentState.Blocks[bi.index]
		return bi
	}
	bi.block = nil
	return nil
}

func (bi blockIterator) NextBlock() *ContentBlock {
	if bi.HasNext() {
		return bi.contentState.Blocks[bi.index+1]
	}
	return nil
}
