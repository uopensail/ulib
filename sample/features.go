package sample

type DataType byte

const (
	Int64Type DataType = iota
	Float32Type
	StringType
	Int64sType
	Float32sType
	StringsType
	ErrorType DataType = 255
)

type Features interface {
	GetType(string) DataType
	Keys() []string
	GetInt64(string) (int64, error)
	GetFloat32(string) (float32, error)
	GetString(string) (string, error)
	GetInt64s(string) ([]int64, error)
	GetFloat32s(string) ([]float32, error)
	GetStrings(string) ([]string, error)
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
