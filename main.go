package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	usage := `Ping.

Usage:
  ping [--count COUNT] DESTINATION

Options:
  -h --help                Show this message.
  -c, --count COUNT   Send this many pings [default: 20].`

	arguments, _ := docopt.Parse(usage, nil, true, "Ping", false)
	host := arguments["DESTINATION"].(string)
	numPings := 20
	if arguments["--count"] != nil {
		numPings, _ = strconv.Atoi(arguments["--count"].(string))
	}

	rAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		panic(err)
	}

	fmt.Printf("PING %s (%s)\n", host, rAddr)

	conn, err := net.DialIP("ip4:icmp", nil, rAddr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	pid := os.Getpid()
	var id1 = byte(pid & 0xff00 >> 8)
	var id2 = byte(pid & 0xff)

	var timeout = 1 * time.Second
	var interval = 1 * time.Second
	var messageLength = 8

	for i := 0; i < numPings; i++ {

		msg := MakeEchoRequest(i, messageLength, id1, id2)

		startTime := time.Now()
		deadline := startTime.Add(timeout)
		conn.SetDeadline(deadline)

		if _, err = conn.Write(msg[0:messageLength]); err != nil {
			continue
		}

		response := make([]byte, 64)
		for {
			numRead, _, err := conn.ReadFrom(response)
			if err != nil {
				panic(err)
			}

			if response[0] == 0 {
				fmt.Printf("%d bytes from %s (%s): time=%v\n",
					numRead,
					host, rAddr,
					time.Since(startTime))
			}

			break
		}

		time.Sleep(interval)
	}
}
