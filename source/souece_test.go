package source

import (
	"fmt"
	"testing"

	"github.com/uopensail/ulib/sample"
)

func TestSource(t *testing.T) {
	s, err := NewSource("/tmp/pool.txt", "id")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer s.Release()

	// test get by key
	fea := s.GetByKey("key")

	// test collection

	c1 := "v1[float32] > 1000.0"
	s.BuildCollection("c1", c1)
	r1 := s.GetCollection("c1")
	for i := 0; i < len(r1) && i < 100; i++ {
		fmt.Println(r1[i])
		fea = s.GetById(r1[i])
		d_dy_4_imppv, _ := fea.Get("v1").GetFloat32()
		fmt.Println(r1[i], ":", d_dy_4_imppv)
	}

	// test condition

	cd := "(v1[float32] > 1000.0) and (param[float32] > v2[float32])"
	fmt.Println(cd)
	s.BuildCondition("c2", cd)
	cd1 := s.GetCondition("c2")
	mfea := sample.NewMutableFeatures()
	mfea.Set("param", &sample.Float32{Value: 400.0})
	r2 := cd1.CheckAll(mfea)
	for i := 0; i < len(r2) && i < 100; i++ {
		fea = s.GetById(r2[i])
		d_dy_use_dur, _ := fea.Get("v2").GetFloat32()
		fmt.Println(r2[i], ":", d_dy_use_dur)
	}
}
