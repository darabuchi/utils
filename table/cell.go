package table

type Cell interface {
	DrawImg()
	MinSize() *Size
	setSize(*Size)
	Size() *Size
}
