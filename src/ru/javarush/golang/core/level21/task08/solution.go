package main

import "fmt"

type ServerAddr struct {
	Host string
	Port int
	Zone string
}

func main() {
	var n int
	fmt.Scan(&n)

	addrs := make([]ServerAddr, 0, n)
	for i := 0; i < n; i++ {
		var host string
		var port int
		fmt.Scan(&host, &port)

		// TODO: создайте ServerAddr через именованный литерал, заполнив только Host и Port; Zone не задавайте
		addrs = append(addrs, ServerAddr{
			Host: host,
			Port: port,
		})
	}

	for i := 0; i < n; i++ {
		a := addrs[i]

		// TODO: выведите адрес в формате <host>:<port> zone="<zone>" (Zone всегда в кавычках), в исходном порядке
		fmt.Printf("%s:%d zone=%q\n", a.Host, a.Port, a.Zone)
	}
}
