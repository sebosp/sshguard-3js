package main

import (
	"testing"

	geoip2 "github.com/oschwald/geoip2-golang"
)

func TestParseSSHFailedPasswordLine(t *testing.T) {
	blacklist := map[string]*connectInfo{}
	db, _ := geoip2.Open("GeoLite2-City.mmdb")
	defer db.Close()
	lessEntriesLogLine := "Not enough records line\n"
	errorLessRecords := parseSSHFailedPasswordLine(lessEntriesLogLine, blacklist, db)
	if errorLessRecords == nil {
		t.Errorf("Expected errors on less records in Log Line")
	}
	moreEntriesLogLine := "Too many records is this line we expect error here\n"
	errorMoreRecords := parseSSHFailedPasswordLine(moreEntriesLogLine, blacklist, db)
	if errorMoreRecords == nil {
		t.Errorf("Expected errors on more records in Log Line")
	}
	badIPLogLine := "Failed password for root from X.144.172.42 port 6721 ssh2\n"
	errorBadIP := parseSSHFailedPasswordLine(badIPLogLine, blacklist, db)
	if errorBadIP == nil {
		t.Errorf("Expected errors on Bad IP")
	}
	rootLogLine := "Failed password for root from 59.144.172.42 port 6721 ssh2\n"
	parseSSHFailedPasswordLine(rootLogLine, blacklist, db)
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
	parseSSHFailedPasswordLine(rootLogLine, blacklist, db)
	parseSSHFailedPasswordLine(rootLogLine, blacklist, db)
	nonRootLogLine := "Failed password for nonroot from 59.144.172.42 port 6721 ssh2\n"
	parseSSHFailedPasswordLine(nonRootLogLine, blacklist, db)
	ipEntry, _ = blacklist["59.144.172.42"]
	if ipEntry.count != 4 {
		t.Errorf("IP count not incremented")
	}
	rootUserEntry, _ := ipEntry.targetUser["root"]
	if rootUserEntry != 3 {
		t.Errorf("Root user count not incremented")
	}
	userEntry, _ := ipEntry.targetUser["nonroot"]
	if userEntry != 1 {
		t.Errorf("Non-Root count not incremented")
	}
}
func TestParseSSHLogLine(t *testing.T) {
	blacklist := map[string]*connectInfo{}
	db, _ := geoip2.Open("GeoLite2-City.mmdb")
	defer db.Close()
	testCases := []struct {
		line    string
		isError bool
	}{
		{"Failed password for Not enough records line\n", true},
		{"Failed password for 1 2 3 4 5 6 7 8 9\n", true},
		{"Failed password for root from 59.144.172.42 port 6721 ssh2\n", false},
		{"Failed password for nonroot from 59.144.172.42 port 6721 ssh2\n", false},
	}
	for _, tc := range testCases {
		err := parseSSHLogLine(tc.line, blacklist, db)
		if (err != nil) != tc.isError {
			t.Errorf("Expected %v", err)
		}
	}
}
