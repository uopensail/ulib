package prome

type CostBucket struct {
	sampleSum float64
}

func (h *CostBucket) Observe(v float64) {
	h.sampleSum += v

}
