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

type sshConnectInfo struct {
	targetUser map[string]int
	count      int
	lastSeen   time.Time
}

func parseFailedPasswordLine(line string, blacklist map[string]sshConnectInfo) error {
	lineContent := strings.Split(line, " ")
	if len(lineContent) != 9 {
		return errors.New("number of fields in Failed Password not 9")
	}
	sourceIP := lineContent[6]
	_, ok := blacklist[sourceIP]
	if ok {
		blacklist[sourceIP].count++
	} else {
		blacklist[lineContent[6]] = sshConnectInfo{
			targetUser: {lineContent[4]: 1},
			count:      1,
			lastSeen:   time.Now(),
		}
	}
	return nil
}
func parseSSHLogLine(line string, blacklist map[string]sshConnectInfo) error {
	if strings.HasPrefix(line, "Failed password for ") {
		if err := parseFailedPasswordLine(line, blacklist); err != nil {
			return err
		}
	}
	return nil
}
func readStdIn(blacklist map[string]sshConnectInfo) {
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
	var blacklist map[string]sshConnectInfo
	readStdIn(blacklist)
	fmt.Println("parser")
}
