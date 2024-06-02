package main

import (
	"runtime"

	"github.com/hubshluft/hubshfetch/cmd"
)

func main() {
	os_slice := []string{"linux", "netbsd", "openbsd", "dragonfly", "freebsd"}
	os_type := false

	for _, str := range os_slice {
		if str == runtime.GOOS {
			os_type = true
			break
		}
	}
	if os_type == true {
		cmd.Args()
	} else {
		return
	}
}
