package sampler

import (
	"github.com/tester-rep/spr-go/abstract"
	"time"
)

type ISampler interface {
	NewSamplerStart(key string, robot abstract.IRobot) ISampler
	SampleEnd(result bool)
	NewSamplerStartWithStartTime(key string, startTime time.Duration, robot abstract.IRobot) ISampler
	SampleEndWithEndTime(result bool, endTime time.Duration)
}
