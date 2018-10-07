package main

import (
	"testing"
)

func TestParseSSHFailedPasswordLine(t *testing.T) {
	blacklist := map[string]*connectInfo{}
	logLine := "Failed password for root from 59.144.172.42 port 6721 ssh2\n"
	parseSSHFailedPasswordLine(logLine, blacklist)
	ipEntry, ipExists := blacklist["59.144.172.42"]
	if !ipExists {
		t.Errorf("IP should have been blacklisted")
	} else {
		if ipEntry.count != 1 {
			t.Errorf("IP count not initialized")
		}
		userEntry, userExists := ipEntry.targetUser["root"]
		if !userExists {
			t.Errorf("User not added to the blacklist map.")
		} else {
			if userEntry != 1 {
				t.Errorf("User count not initialized")
			}
		}
	}
	parseSSHFailedPasswordLine(logLine, blacklist)
	parseSSHFailedPasswordLine(logLine, blacklist)
	ipEntry, _ = blacklist["59.144.172.42"]
	if ipEntry.count != 3 {
		t.Errorf("IP count not incremented")
	}
	userEntry, _ := ipEntry.targetUser["root"]
	if userEntry != 3 {
		t.Errorf("User count not incremented")
	}
}
