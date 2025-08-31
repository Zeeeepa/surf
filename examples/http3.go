package main

import (
	"fmt"
	"log"

	"github.com/enetx/surf"
)

func main() {
	// Example 1: Chrome HTTP/3 fingerprint
	fmt.Println("=== Chrome HTTP/3 Example ===")
	chromeClient := surf.NewClient().Builder().
		Impersonate().Chrome().HTTP3().
		Build()

	r := chromeClient.Get("https://cloudflare-quic.com/").Do()

	switch {
	case r.IsOk():
		resp := r.Ok()
		fmt.Printf("Chrome H3 Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Chrome H3 Protocol: %s\n", resp.Proto)
		fmt.Printf("Chrome H3 Server: %s\n", resp.Headers.Get("server"))
	case r.IsErr():
		log.Printf("Chrome H3 request failed: %v", r.Err())
	}

	// g.String(r.Ok().RemoteAddress().String()).Println()

	r.Ok().Debug().Request(true).Response().Print()

	// Example 2: Firefox HTTP/3 fingerprint
	fmt.Println("\n=== Firefox HTTP/3 Example ===")
	firefoxClient := surf.NewClient().Builder().
		Impersonate().FireFox().HTTP3().
		Build()

	r = firefoxClient.Get("https://cloudflare-quic.com/").Do()

	switch {
	case r.IsOk():
		resp := r.Ok()
		fmt.Printf("Firefox H3 Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Firefox H3 Protocol: %s\n", resp.Proto)
		fmt.Printf("Firefox H3 Server: %s\n", resp.Headers.Get("server"))
	case r.IsErr():
		log.Printf("Firefox H3 request failed: %v", r.Err())
	}

	r.Ok().Debug().Request(true).Response().Print()

	// Example 3: Custom HTTP/3 configuration
	fmt.Println("\n=== Custom HTTP/3 Configuration ===")
	customClient := surf.NewClient().Builder().
		HTTP3Settings().
		Firefox(). // Use built-in Firefox fingerprint
		Set().
		Build()

	r = customClient.Get("https://cloudflare-quic.com/").Do()

	switch {
	case r.IsOk():
		resp := r.Ok()
		fmt.Printf("Custom H3 Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Custom H3 Protocol: %s\n", resp.Proto)
		fmt.Printf("Custom H3 Server: %s\n", resp.Headers.Get("server"))
	case r.IsErr():
		log.Printf("Custom H3 request failed: %v", r.Err())
	}

	r.Ok().Debug().Request(true).Response().Print()

	// Example 4: HTTP/3 with TLS fingerprinting (JA3)
	fmt.Println("\n=== HTTP/3 with JA3 Fingerprinting ===")
	ja3Client := surf.NewClient().Builder().
		JA().Chrome131().               // TLS fingerprint
		HTTP3Settings().Chrome().Set(). // HTTP/3 fingerprint
		Build()

	r = ja3Client.Get("https://cloudflare-quic.com/").Do()

	switch {
	case r.IsOk():
		resp := r.Ok()
		fmt.Printf("JA3 + H3 Status Code: %d\n", resp.StatusCode)
		fmt.Printf("JA3 + H3 Protocol: %s\n", resp.Proto)
		fmt.Printf("JA3 + H3 Server: %s\n", resp.Headers.Get("server"))
	case r.IsErr():
		log.Printf("JA3 + H3 request failed: %v", r.Err())
	}

	r.Ok().Debug().Request(true).Response().Print()
}
