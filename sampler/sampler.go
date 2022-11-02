package sampler

import "github.com/tester-rep/spr-go/abstract"

type ISampler interface {
	NewSamplerStart(key string, robot abstract.IRobot) ISampler
	SampleEnd(result bool)
}
