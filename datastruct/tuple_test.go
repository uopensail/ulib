package datastruct

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestTuple(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := make(TupleList[int, int], 0, 10)
	for i := 0; i < 10; i++ {
		list = append(list, Tuple[int, int]{r.Int(), r.Int()})
	}
	list.Print()
	sort.Sort(list)
	list.Print()
}
