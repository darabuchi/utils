package table

type Text struct {
	text string

	size *Size
}

func (p *Text) setSize(size *Size) {
	p.size = size
}

func (p *Text) Size() *Size {
	if p.size != nil {
		return p.size
	}
	return p.MinSize()
}

func (p *Text) DrawImg() {
	// TODO implement me
	panic("implement me")
}

func (p *Text) MinSize() *Size {
	return NewSize(int32(len(p.text)*12), 12)
}

func NewText(text string) *Text {
	return &Text{text: text}
}
