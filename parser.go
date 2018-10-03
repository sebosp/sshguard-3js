package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type sshConnectInfo struct {
	sourceIP   string
	targetUser string
	count      int
	lastSeen   Time
}

func parseFailedPasswordLine(line string) (sshConnectInfo, err) {
	lineContent := split(line)
	if lineContent.len() != 9 {
		return nil, errors.New("number of fields in Failed Password not 9")
	}
	return sshConnectInfo{
		sourceIP:   lineContent[6],
		targetUser: lineContent[4],
		count:      1,
		lastSeen:   time.Now(),
	}, nil
}
func parseSSHLogLine(line string, blacklist map[string]int) err {
	if strings.startsWith("Failed password for ") {
		if connectInfo, err = parseFailedPasswordLine(line); err != nil {
			return err
		}
	}
}
func readStdIn() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("error:", err)
		}
		parseSSHLogLine(text)
	}

}
func main() {
	var blacklist map[string]int
	readStdIn()
	fmt.Println("parser")
}
