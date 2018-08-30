package main

import (
	"io"
	"log"
	"net"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal("InterfaceAddrs: ", err.Error())
	}

	ipString := "This is "
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipString += ipnet.IP.String() + " "
			}

		}
	}

	ipString += "\r\n"
	io.WriteString(w, ipString)
}

func main() {
	http.HandleFunc("/", echoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

