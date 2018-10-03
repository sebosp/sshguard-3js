package main

import "testing"

func TestParseSSHLogLine(t *testing.T) {
	var blacklist map[string]int
	logLine := "Failed password for root from 59.144.172.42 port 6721 ssh2\n"
	assert.Equal(t, blacklist["59.144.172.42"], 1, "Adds simple IP to blacklist")
}
