package main

import (
	"fmt"

	"github.com/enetx/surf"
)

func main() {
	// to get remote server ip address
	cli := surf.NewClient().Builder().GetRemoteAddress().Build()

	r := cli.Get("ya.ru").Do()
	if r.IsErr() {
		fmt.Println(r.Err())
		return
	}

	fmt.Println(r.Ok().RemoteAddress())
}
