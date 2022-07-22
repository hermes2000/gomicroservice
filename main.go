package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ips", getIPJson)
	fmt.Printf("API endpoint -> http://localhost:8080/ips")
	http.ListenAndServe(":8080", nil)
}

type IP struct {
	Id        int    `json:"id"`
	IP        string `json:"ip"`
	Malicious bool   `json:"malicious"`
	Notes     string `json:"notes"`
}

type multipleIPs struct {
	IPs map[string]IP `json:"ips"`
}

var ipMap = map[string]IP{
	"127.0.0.1": IP{
		Id:        1,
		IP:        "127.0.0.1",
		Malicious: false,
		Notes:     "Home sweet home",
	},
	"8.8.8.8": IP{
		Id:        2,
		IP:        "8.8.8.8",
		Malicious: false,
		Notes:     "Google IP",
	},
	"123.123.123.123": IP{
		Id:        3,
		IP:        "123.123.123.123",
		Malicious: true,
		Notes:     "Compromised one of our computers",
	},
}

func getIP() multipleIPs {
	var m multipleIPs
	m.IPs = make(map[string]IP)
	for k, ip := range ipMap {
		m.IPs[k] = ip
	}

	return m
}

func getIPJson(w http.ResponseWriter, r *http.Request) {
	ips := getIP()
	data, err := json.Marshal(ips)
	if err != nil {
		log.Fatalln(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}
