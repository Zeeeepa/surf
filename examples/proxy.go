package main

import (
	"fmt"
	"log"

	"github.com/enetx/surf"
)

func main() {
	type Proxy struct {
		ISTor bool   `json:"IsTor"`
		IP    string `json:"IP"`
	}

	const url = "https://check.torproject.org/api/ip"

	// for random select proxy from slice
	r := surf.NewClient().
		Builder().
		Proxy([]string{
			"socks5://127.0.0.1:9050",
			"socks5://127.0.0.1:9050",
		}).
		Build().
		Get(url).
		Do()

	// r := surf.NewClient().
	// 	Builder().
	// 	Proxy("http://127.0.0.1:8080").
	// 	Build().
	// 	Get(url).
	// 	Do()

	if r.IsErr() {
		log.Fatal(r.Err())
	}

	var proxy Proxy

	r.Ok().Body.JSON(&proxy)

	fmt.Printf("is tor: %v, ip: %s", proxy.ISTor, proxy.IP)
}
