package main

import (
	"errors"
)

// BlacklistService provides operations on the Blacklist map
type BlacklistService interface {
	GetIPDetails(string) (*connectInfo, error)
	GetIPCount(string) int
	GetIPsActiveSince(int64) []*connectInfo
}

// blacklistService contains a loaded blacklist table, the "key" is the IP.
type blacklistService struct {
	blacklistTable map[string]*connectInfo
}

// GetIPDetails returns the connectionInfo for an IP.
func (b blacklistService) GetIPDetails(ip string) (*connectInfo, error) {
	if ip == "" {
		return nil, errEmptyIP
	}
	ipEntry, ipExists := b.blacklistTable[ip]
	if !ipExists {
		return nil, errors.New("IP does not exist in the blacklist table")
	} else {
		return ipEntry, nil
	}
}

// GetIPCount retuns the count inside of a connectionInfo
func (b blacklistService) GetIPCount(ip string) int {
	ipData, err := b.GetIPDetails(ip)
	if err == nil {
		return ipData.count
	}
	return 0
}

// GetIPsActiveSince returns an array with the IPs that have been active after a specific Epoch
func (b blacklistService) GetIPsActiveSince(epoch int64) []*connectInfo {
	res := make([]*connectInfo, 0)
	for _, ip := range b.blacklistTable {
		if ip.lastSeen.Unix() > epoch {
			res = append(res, ip)
		}
	}
	return res
}

var errEmptyIP = errors.New("Empty IP parameter")
