package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)

const SleepTime = 60
const QuarantineTicker = 10

func main() {
	savelog()
	log.Println("serverChecker started")
	log.Println("reading servers to check")

	servers := readConfig("servers.json")
	log.Println(servers)

	ticker := time.NewTicker(time.Second * SleepTime)
	for _ = range ticker.C {
		for _, server := range servers {
			color.Green(server.URL)
			err := checkURL(server.URL)
			if err != nil {
				go quarantine(server)
			}
		}
	}
}
func quarantine(server Server) {
	var errCount int
	var iterations int
	ticker := time.NewTicker(time.Second * QuarantineTicker)
	for _ = range ticker.C {
		color.Yellow("quarentene")
		fmt.Println(iterations)
		err := checkURL(server.URL)
		if err != nil {
			errCount++
		}
		iterations++

		if errCount > 3 {
			color.Red("server error")
			log.Println("[url]: " + server.URL + " - ERROR: " + err.Error())
			return
		}
		if iterations > 10 {
			return
		}
	}
}
func checkURL(url string) error {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(url)
	return err
}
