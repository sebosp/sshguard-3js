package main

import (
	"testing"
)

func TestParseSSHFailedPasswordLine(t *testing.T) {
	blacklist := map[string]*connectInfo{}
	logLine := "Failed password for root from 59.144.172.42 port 6721 ssh2\n"
	parseSSHFailedPasswordLine(logLine, blacklist)
	val, ok := blacklist["59.144.172.42"]
	if !ok {
		t.Errorf("IP should have been blacklisted")
	} else {
		if val.count != 1 {
			t.Errorf("IP count not incremented")
		}
	}
}
