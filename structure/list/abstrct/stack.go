package abstrct

type IStack interface {
	Push(data interface{})
	Pop() (interface{}, error)
	Length() int
	IsEmpty() bool
	Clear()
}
