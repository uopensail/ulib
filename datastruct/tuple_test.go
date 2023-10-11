package datastruct

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTupleMarshal(t *testing.T) {
	v := Tuple[string, int]{First: "string", Second: 1}
	data, _ := json.Marshal(v)
	fmt.Println(string(data))
	var v1 Tuple[string, int]
	json.Unmarshal(data, &v1)
	v1.Print()
}
