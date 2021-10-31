package abstrct

type IStack interface {
	Push(data interface{})
	Pop() interface{}
	Length() int
	IsEmpty() bool
	Clear()
}
