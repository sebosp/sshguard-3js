package main

import "testing"

func TestParseSSHLogLine(t *testing.T) {
	var blacklist map[string]sshConnectInfo
	logLine := "Failed password for root from 59.144.172.42 port 6721 ssh2\n"
	val, ok := blacklist["59.144.172.42"]
	if !ok {
		t.Errorf("IP should have been blacklisted")
	}
	if val != 1 {
		t.Errorf("IP count not incremented")
	}
}
