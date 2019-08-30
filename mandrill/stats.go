package mandrill

import (
	"time"

	"github.com/lusis/gochimp/mandrill/api"
)

type Stats = api.StatsResponse

type TimeSeries struct {
	Time         time.Time
	Sent         int32
	HardBounces  int32
	SoftBounces  int32
	Rejects      int32
	Complaints   int32
	Unsubs       int32
	Opens        int32
	UniqueOpens  int32
	Clicks       int32
	UniqueClicks int32
	Reputation   int32
}
