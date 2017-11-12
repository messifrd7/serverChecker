package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)

const SleepTime = 10

func main() {
	savelog()
	log.Println("serverChecker started")
	log.Println("reading servers to check")

	servers := readConfig("servers.json")
	log.Println(servers)

	ticker := time.NewTicker(time.Second * SleepTime)
	for _ = range ticker.C {
		for _, server := range servers {
			checkURL(server.URL)
			time.Sleep(time.Millisecond * 3000)
		}
	}
}
func checkURL(url string) {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		color.Red("[url]:" + url + " - ERROR: " + err.Error())
		log.Println(err.Error())
	} else {
		//log.Println(string(resp.StatusCode) + resp.Status)
		color.Green("[url]: " + url + " - " + string(resp.StatusCode) + resp.Status)
	}
}
