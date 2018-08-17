package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var privateIPBlocks []*net.IPNet

func init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
	} {
		_, block, _ := net.ParseCIDR(cidr)
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

func isPrivateIP(ip net.IP) bool {
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

var commandOptions = struct {
	ip string
}{
	"127.0.0.1",
}

func init() {
	flag.StringVar(&commandOptions.ip, "ip", commandOptions.ip, "Geo Location search Ip address")
	flag.Parse()
}

func main() {
	db, err := geoip2.Open("./GeoLite2-City.mmdb")
	db2, err := geoip2.Open("./GeoLite2-ASN.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ip := net.ParseIP(commandOptions.ip)

	fmt.Printf("PrivateNetwork: %t | GlobalUnicast: %t", isPrivateIP(ip), ip.IsGlobalUnicast())
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	record2, _ := db2.ASN(ip)

	PrettyPrint(record)

	fmt.Println("-----------ASN------------------")
	PrettyPrint(record2)

}
