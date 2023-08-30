package sample

type DataType byte

const (
	Int64Type DataType = iota
	Float32Type
	StringType
	Int64ArrayType
	Float32ArrayType
	StringArrayType
	ErrorType DataType = 255
)

type Features interface {
	GetInt64(string) (int64, error)
	GetFloat32(string) (float32, error)
	GetString(string) (string, error)
	GetInt64Array(string) ([]int64, error)
	GetFloat32Array(string) ([]float32, error)
	GetStringArray(string) ([]string, error)
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
