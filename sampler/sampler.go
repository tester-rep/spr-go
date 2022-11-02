package sampler

type ISampler interface {
	NewSamplerStart(key string) *ISampler
	SampleEnd(result bool)
}
