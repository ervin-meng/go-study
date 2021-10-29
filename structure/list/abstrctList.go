package list

type List interface {
	Clear()
	IsEmpty() bool
	Length() int
	Append(value interface{})
	Insert(index int, value interface{}) error
	Delete(index int) error
	Update(index int, value interface{}) error
	Get(index int) (interface{}, error)
	//Sort() error
}
