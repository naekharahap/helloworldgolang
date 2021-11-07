package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type helloMessage struct {
	Message string   `json:"message"`
	Host    string   `json:"host"`
	Ip      []string `json:"ip"`
}

func HelloWorld() string {
	var h helloMessage

	h.Message = "Hello world golang"

	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err)
	}
	h.Host = hostname

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return fmt.Sprintf(`{"error": "%s"}`, err)
	}
	for _, addr := range addrs {
		h.Ip = append(h.Ip, addr.String())
	}

	m, err := json.MarshalIndent(&h, "", "  ")
	if err != nil {
		fmt.Println(err)
		return fmt.Sprintf(`{"error": "%s"}`, err)
	}

	return string(m)
}
