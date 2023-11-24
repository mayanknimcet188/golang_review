package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type TimeZone struct {
	tz           string
	serverConfig string
}

func main() {
	arguments := os.Args[1:]
	var timeZones []TimeZone
	for _, argument := range arguments {
		splitted := strings.Split(argument, "=")
		timeZones = append(timeZones, TimeZone{splitted[0], splitted[1]})
	}
	for _, entry := range timeZones {
		fmt.Println("Timezone: ", entry.tz)
		connectToServer(entry.serverConfig)
	}
}

func connectToServer(serverConfig string) {
	conn, err := net.Dial("tcp", serverConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// TODO FIX: Currently the for loop is displaying only the time for the first entry in the array
// maybe a blocking operation
