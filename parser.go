package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"strings"
	"time"

	geoip2 "github.com/oschwald/geoip2-golang"
)

// connectInfo could potentially be used for other connection attempts, not just SSH
type connectInfo struct {
	targetUser map[string]int
	count      int
	lastSeen   time.Time
	latitude   float64
	longitude  float64
}

// parseSSHFailedPasswordLine checks SSH failed password and feeds the blacklist struct
func parseSSHFailedPasswordLine(line string, blacklist map[string]*connectInfo, db *geoip2.Reader) error {
	// Example of error line:
	// "Failed password for root from WW.XX.YY.ZZ port 6721 ssh2"
	lineContent := strings.Split(line, " ")
	if len(lineContent) != 9 {
		return errors.New("number of fields in Failed Password not 9")
	}
	sourceIP := lineContent[5]
	ip := net.ParseIP(sourceIP)
	if ip == nil {
		return errors.New("Invalid sourceIP")
	}
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
		latitude := -10000.0
		longitude := -10000.0
		record, err := db.City(ip)
		if err == nil {
			latitude = record.Location.Latitude
			longitude = record.Location.Longitude
		}
		blacklist[sourceIP] = &connectInfo{
			targetUser: map[string]int{targetUser: 1},
			count:      1,
			lastSeen:   time.Now(),
			latitude:   latitude,
			longitude:  longitude,
		}
	}
	return nil
}

// parseSSHLogLine should cover main ssh log lines related to errors.
func parseSSHLogLine(line string, blacklist map[string]*connectInfo, db *geoip2.Reader) error {
	if strings.HasPrefix(line, "Failed password for ") {
		if err := parseSSHFailedPasswordLine(line, blacklist, db); err != nil {
			return err
		}
	}
	return nil
}

// exportJSON writes a JSON file version of the blacklist table.
func exportJSON(blacklist map[string]*connectInfo) {
	min := float64(math.MaxInt32)
	max := float64(0)
	// Find the min and max
	for _, ip := range blacklist {
		ipCount := float64(ip.count)
		if max < ipCount {
			max = ipCount
		}
		if min > ipCount {
			min = ipCount
		}
	}
	max = max + 1
	scale := (max - min) / 10
	buckets := [10]string{}
	fmt.Printf("Max is %f, Min is %f, scale is %f\n", max, min, scale)
	// Create dataset "buckets" for the data:
	for _, ip := range blacklist {
		bucketNumber := int(math.Floor((float64(ip.count) - min) / scale))
		if buckets[bucketNumber] != "" {
			buckets[bucketNumber] += ","
		}
		buckets[bucketNumber] += fmt.Sprintf("%f,%f,%d", ip.latitude, ip.longitude, ip.count)
	}
	fmt.Println("var data = [")
	for i, bucketData := range buckets {
		if i > 0 {
			fmt.Println(",")
		}
		fmt.Printf("['series%d',[%s]]", i, bucketData)
	}
	fmt.Println("];")
}

// readStdIn shows that ideally we are reading a pipe
func readStdIn(blacklist map[string]*connectInfo, db *geoip2.Reader) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("error:", err)
			}
			break
		}
		parseSSHLogLine(text, blacklist, db)
	}
}
func main() {
	blacklist := map[string]*connectInfo{}
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	readStdIn(blacklist, db)
	exportJSON(blacklist)
}
