package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type StringMap map[string]string

func (s *StringMap) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), s)
	case []byte:
		return json.Unmarshal(val, s)
	default:
		return errors.New("not support")
	}

}

func (s StringMap) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}
func (s StringMap) GormDataType() string {
	return "json"
}

type IntSlice []int

func (s *IntSlice) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), s)
	case []byte:
		return json.Unmarshal(val, s)
	default:
		return errors.New("not support")
	}

}

func (s IntSlice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func (s IntSlice) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (IntSlice) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	case "sqlserver":
		return "NVARCHAR(MAX)"
	}
	return ""
}

type Float32Slice []float32

func (s *Float32Slice) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), s)
	case []byte:
		return json.Unmarshal(val, s)
	default:
		return errors.New("not support")
	}

}

func (s Float32Slice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func (s Float32Slice) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (Float32Slice) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	case "sqlserver":
		return "NVARCHAR(MAX)"
	}
	return ""
}

type Int64Slice []int64

func (s *Int64Slice) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), s)
	case []byte:
		return json.Unmarshal(val, s)
	default:
		return errors.New("not support")
	}

}

func (s Int64Slice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func (s Int64Slice) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (Int64Slice) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	case "sqlserver":
		return "NVARCHAR(MAX)"
	}
	return ""
}

type Float64Slice []float64

func (s *Float64Slice) Scan(val interface{}) error {
	switch val := val.(type) {
	case string:
		return json.Unmarshal([]byte(val), s)
	case []byte:
		return json.Unmarshal(val, s)
	default:
		return errors.New("not support")
	}

}

func (s Float64Slice) Value() (driver.Value, error) {
	bytes, err := json.Marshal(s)
	return string(bytes), err
}

func (s Float64Slice) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (Float64Slice) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	case "sqlserver":
		return "NVARCHAR(MAX)"
	}
	return ""
}
