package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	author := os.Getenv("AUTHOR")
	if author == "" {
		author = "--Unknown--"
	}
	hostname, err := os.Hostname()
	if err != nil {
		panic(fmt.Sprintf("failed to get hostname: %s", err.Error()))
	}
	ip, err := getIPAddress()
	if err != nil {
		panic(fmt.Sprintf("failed to get ip address: %s", err.Error()))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("public/index.html"))

		data := map[string]string{
			"Hostname": hostname,
			"IP":       ip.String(),
			"Author":   author,
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Printf("failed to execute template: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Printf("server has stopped: %s", err.Error())
	}
}

func getIPAddress() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if !ip.IsLoopback() && ip.To4() != nil {
				return ip, nil
			}
		}
	}
	return nil, errors.New("address not found")
}
