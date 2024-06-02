package cmd

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

var hey string

func displayFetch() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	s := strings.TrimSpace(user.Username) + "@" + hostname

	for i := 0; i < len(s); i++ {
		hey += "-"
	}

	version, err := os.ReadFile("/proc/version")
	if err != nil {
		panic(err)
	}
	kernel := strings.Fields(strings.TrimSpace(string(version)))

	uptime_readall, err := os.ReadFile("/proc/uptime")
	if err != nil {
		panic(err)
	}
	uptime := strings.Fields(strings.TrimSpace(string(uptime_readall)))
	parts := strings.Split(uptime[0], ".")
	uptime2int, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n%s\n", s, hey)
	fmt.Printf("OS:\t%s %s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("KERNEL: %s\n", kernel[2])
	fmt.Printf("WM: \t%s\n", os.Getenv("XDG_SESSION_DESKTOP"))
	fmt.Printf("UPTIME: %d mins\n", uptime2int/60)
	fmt.Printf("")
}
