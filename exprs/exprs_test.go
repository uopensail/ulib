package exprs

import (
	"fmt"
	"testing"
)

func TestFeatures(t *testing.T) {
	handler := NewExprsHandler("/tmp/config.toml")
	defer handler.Release()

	features := handler.Call(`{"h": {"type": 3, "value": [4, 8]}}`)
	defer features.Release()

	// Iterate over the features in the FeatureSet
	for key, feature := range features.features {
		// Attempt to retrieve the string slice from the feature
		if values, err := feature.GetStrings(); err == nil {
			fmt.Printf("Key: %s, Value: %v\n", key, values)
		} else {
			fmt.Printf("Key: %s, Error: %v\n", key, err)
		}
	}
}
