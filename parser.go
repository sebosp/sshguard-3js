package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// connectInfo could potentially be used for other connection attempts, not just SSH
type connectInfo struct {
	targetUser map[string]int
	count      int
	lastSeen   time.Time
}

// parseSSHFailedPasswordLine checks SSH failed password and feeds the blacklist struct
func parseSSHFailedPasswordLine(line string, blacklist map[string]*connectInfo) error {
	// Example of error line:
	// "Failed password for root from WW.XX.YY.ZZ port 6721 ssh2"
	lineContent := strings.Split(line, " ")
	if len(lineContent) != 9 {
		return errors.New("number of fields in Failed Password not 9")
	}
	sourceIP := lineContent[5]
	targetUser := lineContent[3]
	_, ok := blacklist[sourceIP]
	if ok {
		blacklist[sourceIP].count++
		_, ok := blacklist[sourceIP].targetUser[targetUser]
		if ok {
			blacklist[sourceIP].targetUser[targetUser]++
		} else {
			blacklist[sourceIP].targetUser[targetUser] = 1
		}
	} else {
		blacklist[sourceIP] = &connectInfo{
			targetUser: map[string]int{targetUser: 1},
			count:      1,
			lastSeen:   time.Now(),
		}
	}
	return nil
}

// parseSSHLogLine should cover main ssh log lines related to errors.
func parseSSHLogLine(line string, blacklist map[string]*connectInfo) error {
	if strings.HasPrefix(line, "Failed password for ") {
		if err := parseSSHFailedPasswordLine(line, blacklist); err != nil {
			return err
		}
	}
	return nil
}

// readStdIn shows that ideally we are reading a pipe
func readStdIn(blacklist map[string]*connectInfo) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error:", err)
		}
		parseSSHLogLine(text, blacklist)
	}
}
func main() {
	var blacklist map[string]*connectInfo
	readStdIn(blacklist)
	fmt.Println("parser")
}
