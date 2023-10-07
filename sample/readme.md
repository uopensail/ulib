# sample

## feature interface
This is the interface of feature, defining the following functions:
1. Type() DataType                  // return the data type of this feature
2. GetInt64() (int64, error)        
3. GetFloat32() (float32, error)
4. GetString() (string, error)
5. GetInt64s() ([]int64, error)
6. GetFloat32s() ([]float32, error)
7. GetStrings() ([]string, error)

## features interface
This is a set of features encapsulated with a map, defining the following functions:
1. Keys() []string                  // return all the keys of features
2. Get(string) Feature
3. MarshalJSON() ([]byte, error)
4. UnmarshalJSON(data []byte) error


## immutable features
This is the generated features by reading from string and cannot be changed. This structure declutters the memory used to store the features, reduces memory fragmentation, and thus lowers the GC. The memory structure of each type is as follows:

### int64
1. 4byte: data type
2. 4byte: not used
3. 8byte: int64 value

### int64 array
1. 4byte: data type
2. 4byte: not used
3. 24byte: reflect.SliceHeader, reflect.SliceHeader.Data point to the data part
4. data: int64 array data

### float32
1. 4byte: data type
2. 4byte: float32 value

### float32 array
1. 4byte: data type
2. 4byte: not used
3. 24byte: reflect.SliceHeader, reflect.SliceHeader.Data point to the data part
4. data: float32 array data


### string
1. 4byte: data type
2. 4byte: not used
3. 16byte: reflect.StringHeader, reflect.StringHeader.Data point to the data part
4. data: string value data, due to memory alignment, the length is divisible by 8

### string array
1. 4byte: data type
2. 4byte: not used
3. 24byte: reflect.SliceHeader, reflect.SliceHeader.Data point to the string header part
4. string heads: reflect.StringHeader array
5. data: byte list values, due to memory alignment, the length is divisible by 8


## mutable features
This is a changeable features struct, stored in the map. Nothing special to say. Immutable Features can be converted to mutable features.