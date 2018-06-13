package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	respond := make(chan string, 5)

	go checkDNS(respond, "pragmacoders.com", "ns1.nameserver.com")

	for {
		queryResp := <-respond
		fmt.Printf("Got Response:\t %s\n", queryResp)
	}
}

func checkDNS(respond chan<- string, query string, ns string) {
	for {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		respond <- fmt.Sprintf("%s responded to query: %s", ns, query)
	}
}
