package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           BlacklistService
}

func (mw instrumentingMiddleware) GetIPDetails(s string) (output *connectInfo, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getIPDetails", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetIPDetails(s)
	return
}

func (mw instrumentingMiddleware) GetIPCount(ip string) (count int) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getIPCount", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		mw.countResult.Observe(float64(count))
	}(time.Now())

	count = mw.next.GetIPCount(ip)
	return
}

func (mw instrumentingMiddleware) GetIPsActiveSince(epoch int64) (output []*connectInfo) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getIPsActiveSince", "error", "false"}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
		mw.countResult.Observe(float64(len(output)))
	}(time.Now())

	output = mw.next.GetIPsActiveSince(epoch)
	return
}
