package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// errUnsupportedScanType is returned by Scan when the database driver
// provides a value type that is neither string nor []byte.
var errUnsupportedScanType = errors.New("unsupported scan type: expected string or []byte")

// jsonScan unmarshals a SQL column value (string or []byte) into dst.
// It is the shared implementation used by all custom GORM types below.
func jsonScan(dst any, val any) error {
	switch v := val.(type) {
	case string:
		return json.Unmarshal([]byte(v), dst)
	case []byte:
		return json.Unmarshal(v, dst)
	default:
		return errUnsupportedScanType
	}
}

// jsonValue serialises src to a JSON string for storage in a SQL column.
func jsonValue(src any) (driver.Value, error) {
	b, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

// gormDBDataType returns the appropriate column type for common SQL dialects.
// postgres uses JSONB for efficient binary-JSON indexing; all others use JSON.
func gormDBDataType(db *gorm.DB) string {
	switch db.Dialector.Name() {
	case "postgres":
		return "JSONB"
	case "sqlserver":
		return "NVARCHAR(MAX)"
	default:
		// sqlite, mysql and unknown dialects
		return "JSON"
	}
}

// StringMap is a map[string]string that is persisted as a JSON column by GORM.
type StringMap map[string]string

// Scan implements sql.Scanner. It deserialises the column value into the map.
func (s *StringMap) Scan(val any) error { return jsonScan(s, val) }

// Value implements driver.Valuer. It serialises the map to a JSON string.
func (s StringMap) Value() (driver.Value, error) { return jsonValue(s) }

// GormDataType returns the logical GORM data type.
func (s StringMap) GormDataType() string { return "json" }

// IntSlice is a []int that is persisted as a JSON column by GORM.
type IntSlice []int

// Scan implements sql.Scanner.
func (s *IntSlice) Scan(val any) error { return jsonScan(s, val) }

// Value implements driver.Valuer.
func (s IntSlice) Value() (driver.Value, error) { return jsonValue(s) }

// GormDataType returns the logical GORM data type.
func (s IntSlice) GormDataType() string { return "json" }

// GormDBDataType returns the physical column type for the target database.
func (IntSlice) GormDBDataType(db *gorm.DB, _ *schema.Field) string { return gormDBDataType(db) }

// Float32Slice is a []float32 that is persisted as a JSON column by GORM.
type Float32Slice []float32

// Scan implements sql.Scanner.
func (s *Float32Slice) Scan(val any) error { return jsonScan(s, val) }

// Value implements driver.Valuer.
func (s Float32Slice) Value() (driver.Value, error) { return jsonValue(s) }

// GormDataType returns the logical GORM data type.
func (s Float32Slice) GormDataType() string { return "json" }

// GormDBDataType returns the physical column type for the target database.
func (Float32Slice) GormDBDataType(db *gorm.DB, _ *schema.Field) string { return gormDBDataType(db) }

// Int64Slice is a []int64 that is persisted as a JSON column by GORM.
type Int64Slice []int64

// Scan implements sql.Scanner.
func (s *Int64Slice) Scan(val any) error { return jsonScan(s, val) }

// Value implements driver.Valuer.
func (s Int64Slice) Value() (driver.Value, error) { return jsonValue(s) }

// GormDataType returns the logical GORM data type.
func (s Int64Slice) GormDataType() string { return "json" }

// GormDBDataType returns the physical column type for the target database.
func (Int64Slice) GormDBDataType(db *gorm.DB, _ *schema.Field) string { return gormDBDataType(db) }

// Float64Slice is a []float64 that is persisted as a JSON column by GORM.
type Float64Slice []float64

// Scan implements sql.Scanner.
func (s *Float64Slice) Scan(val any) error { return jsonScan(s, val) }

// Value implements driver.Valuer.
func (s Float64Slice) Value() (driver.Value, error) { return jsonValue(s) }

// GormDataType returns the logical GORM data type.
func (s Float64Slice) GormDataType() string { return "json" }

// GormDBDataType returns the physical column type for the target database.
func (Float64Slice) GormDBDataType(db *gorm.DB, _ *schema.Field) string { return gormDBDataType(db) }
