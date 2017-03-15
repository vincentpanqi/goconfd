package net2

import (
	"net"
	"regexp"
	"strings"
)

func CheckIp(ip string) bool {
	if m, _ := regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}(\\:[0-9]{1,5})?$", ip); !m {
		return false
	}
	return true
}

func GetLocalIPv4() (string, error) {
	ips, err := GetLocalIPs()
	if err != nil || len(ips) == 0 {
		return "", err
	}
	ip := ips[0]
	return ip, nil
}

func GetLocalIPs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ips := []string{}
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip := ipnet.IP.String()
					ips = append(ips, ip)
					//fmt.Printf("%v : %s (%s)\n", i.Name, ipnet.IP.String(), ipnet.IP.DefaultMask())
				}
			}
		}
	}
	return ips, nil
}

func ParseHosts(hosts string) []string {
	results := []string{}
	idx := strings.Index(hosts, ";")
	if idx == -1 {
		results = append(results, hosts)
	} else {
		hs := strings.Split(hosts, ";")
		for _, h := range hs {
			h = strings.TrimSpace(h)
			if CheckIp(h) {
				results = append(results, h)
			}
		}
	}
	return results
}
