package exprs

/*
#cgo LDFLAGS: -L/usr/local/lib -lcminia

#include <stdlib.h>
#include "cminia.h"
*/
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/uopensail/ulib/sample"
)

type Feature struct {
	dtype sample.DataType
	data  unsafe.Pointer
}

func (f *Feature) Type() sample.DataType {
	return f.dtype
}

func (f *Feature) GetInt64() (int64, error) {
	if f.dtype != sample.Int64Type {
		return 0, fmt.Errorf("type mismatch: expected %d got %d", sample.Int64Type, f.dtype)
	}
	return *(*int64)(f.data), nil
}

func (f *Feature) GetFloat32() (float32, error) {
	if f.dtype != sample.Float32Type {
		return 0.0, fmt.Errorf("type mismatch: expected %d got %d", sample.Float32Type, f.dtype)
	}
	return *(*float32)(f.data), nil
}

func (f *Feature) GetString() (string, error) {
	if f.dtype != sample.StringType {
		return "", fmt.Errorf("type mismatch: expected %d got %d", sample.StringType, f.dtype)
	}
	return *(*string)(f.data), nil
}

func (f *Feature) GetInt64s() ([]int64, error) {
	if f.dtype != sample.Int64sType {
		return nil, fmt.Errorf("type mismatch: expected %d got %d", sample.Int64sType, f.dtype)
	}
	return *(*[]int64)(f.data), nil
}

func (f *Feature) GetFloat32s() ([]float32, error) {
	if f.dtype != sample.Float32sType {
		return nil, fmt.Errorf("type mismatch: expected %d got %d", sample.Float32sType, f.dtype)
	}
	return *(*[]float32)(f.data), nil
}

func (f *Feature) GetStrings() ([]string, error) {
	if f.dtype != sample.StringsType {
		return nil, fmt.Errorf("type mismatch: expected %d got %d", sample.StringsType, f.dtype)
	}
	return *(*[]string)(f.data), nil
}

func (f *Feature) Release() {
	C.minia_del_feature(f.data, C.int32_t(f.dtype))
}

type Features struct {
	instancePtr unsafe.Pointer
	features    map[string]Feature
}

func NewFeatures(ptr unsafe.Pointer, featureNames []string) *Features {
	featureMap := make(map[string]Feature, len(featureNames))

	for _, name := range featureNames {
		cName := C.CString(name)
		defer C.free(unsafe.Pointer(cName))

		var dtype sample.DataType
		cFeature := C.minia_get_feature(
			ptr,
			cName,
			(*C.int32_t)(unsafe.Pointer(&dtype)),
		)

		featureMap[name] = Feature{
			dtype: dtype,
			data:  cFeature,
		}
	}

	return &Features{
		instancePtr: ptr,
		features:    featureMap,
	}
}

func (fs *Features) Keys() []string {
	keys := make([]string, 0, len(fs.features))
	for key := range fs.features {
		keys = append(keys, key)
	}
	return keys
}

func (fs *Features) Len() int {
	return len(fs.features)
}

func (fs *Features) Get(key string) sample.Feature {
	if fea, ok := fs.features[key]; ok {
		return &fea
	}
	return nil
}

func (fs *Features) UnmarshalJSON(data []byte) error {
	return fmt.Errorf("not implement error")
}

func (fs *Features) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("not implement error")
}

func (fs *Features) MapAny() map[string]any {
	return nil
}

func (fs *Features) Release() {
	for _, feature := range fs.features {
		feature.Release()
	}
	C.minia_del_features(fs.instancePtr)
}

type ExprsHandler struct {
	instancePtr     unsafe.Pointer
	featuresListPtr unsafe.Pointer
	featureNames    []string
}

func NewExprsHandler(configPath string) *ExprsHandler {
	cConfig := C.CString(configPath)
	defer C.free(unsafe.Pointer(cConfig))

	instance := C.minia_create(cConfig)
	features := C.minia_features(instance)

	return &ExprsHandler{
		instancePtr:     instance,
		featuresListPtr: features,
		featureNames:    *(*[]string)(features),
	}
}

func (h *ExprsHandler) Call(input string) *Features {
	cInput := C.CString(input)
	defer C.free(unsafe.Pointer(cInput))

	return NewFeatures(
		C.minia_call(h.instancePtr, cInput),
		h.featureNames,
	)
}

func (h *ExprsHandler) Release() {
	C.minia_del_feature(h.featuresListPtr, C.int32_t(sample.StringsType))
	C.minia_release(h.instancePtr)
}
