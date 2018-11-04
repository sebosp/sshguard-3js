package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   BlacklistService
}

func (mw loggingMiddleware) GetIPDetails(s string) (output *connectInfo, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "getIPDetails",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetIPDetails(s)
	return
}

func (mw loggingMiddleware) GetIPCount(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "getIPCount",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.next.GetIPCount(s)
	return
}

func (mw loggingMiddleware) GetIPsActiveSince(i int64) (output []*connectInfo) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "getIPCount",
			"input", i,
			"n", output,
			"took", time.Since(begin),
		)
	}(time.Now())

	output = mw.next.GetIPsActiveSince(i)
	return
}
