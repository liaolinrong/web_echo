package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal("InterfaceAddrs: ", err.Error())
	}

	ipString := "This is: \r\n"
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipString += "\t" + ipnet.IP.String() + "\r\n"
			}

		}
	}

	//ipString += "\r\n"
	io.WriteString(w, ipString)
}

func echoHostnameHandler(w http.ResponseWriter, r *http.Request) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		log.Fatal("InterfaceAddrs: ", err.Error())
	}

	ipString := "This is: \r\n"
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipString += "\t" + ipnet.IP.String() + "\r\n"
			}

		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		ipString += "failed to get hostname \r\n"
	} else {
		ipString += "hostname is: \r\n\t" + hostname + " \r\n"
	}

	io.WriteString(w, ipString)
}

func main() {
	const ROOT_ROUTER string = "/"
	const HOSTNAME_ROUTER string = "/hostname"

	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		fmt.Println("env LISTEN_PORT is not set, listen default port: 8080")
		listenPort = "8080"
	} else {
		fmt.Printf("env LISTEN_PORT is %s, listen %s\n", listenPort, listenPort)
	}

	fmt.Printf("\nsupport http route:\n")
	fmt.Println(ROOT_ROUTER)
	fmt.Println(HOSTNAME_ROUTER)

	http.HandleFunc(ROOT_ROUTER, echoHandler)
	http.HandleFunc(HOSTNAME_ROUTER, echoHostnameHandler)
	err := http.ListenAndServe(":"+listenPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
