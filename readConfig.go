package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Server struct {
	URL string `json:"url"`
}

func readConfig(path string) []Server {
	var servers []Server

	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &servers)

	return servers
}
