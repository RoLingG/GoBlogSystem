package main

import (
	"log"
	"net"
)

func main() {
	interfaces, err := net.Interfaces() //从网卡获取本机ip
	if err != nil {
		log.Fatal(err)
	}
	for _, i2 := range interfaces {
		address, err := i2.Addrs()
		if err != nil {
			log.Fatal(err)
			continue
		}
		for _, addr := range address {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			//过滤掉ipv6
			ipv4 := ipNet.IP.To4()
			if ipv4 == nil {
				continue
			}
		}
	}
}
