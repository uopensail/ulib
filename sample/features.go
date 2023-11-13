package sample

type DataType int32

const (
	Int64Type DataType = iota
	Float32Type
	StringType
	Int64sType
	Float32sType
	StringsType
	ErrorType DataType = 127
)

type Feature interface {
	Type() DataType
	GetInt64() (int64, error)
	GetFloat32() (float32, error)
	GetString() (string, error)
	GetInt64s() ([]int64, error)
	GetFloat32s() ([]float32, error)
	GetStrings() ([]string, error)
}

type Features interface {
	Keys() []string
	Len() int
	Get(string) Feature
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
