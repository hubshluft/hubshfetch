package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

// #include <unistd.h>
import "C"

var (
	err                          error
	User                         string
	Hostname, Lines              string
	Kernel                       string
	Uptime                       int
	MemoryTotal, MemoryAvailable int
	Packages                     int
	CPU                          string
)

func getUser() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	User = u.Username
}

func getHostname() {
	u, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	Hostname = u
}

func getKernel() {
	u, err := os.ReadFile("/proc/version")
	if err != nil {
		panic(err)
	}

	Kernel = strings.Fields(strings.TrimSpace(string(u)))[2]
}

func getUptime() {
	u, err := os.ReadFile("/proc/uptime")
	if err != nil {
		panic(err)
	}

	uptime_fields := strings.Fields(strings.TrimSpace(string(u)))[0]
	parts := strings.Split(uptime_fields, ".")
	uptimeSeconds, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	hey := uptimeSeconds / 60

	Uptime = hey
}

func getMemory() {
	memory, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	mem := strings.Fields(strings.TrimSpace(string(memory)))
	mem2int, err := strconv.Atoi(mem[7])
	if err != nil {
		fmt.Println(err)
	}
	pages := C.sysconf(C._SC_PHYS_PAGES)
	pageSize := C.sysconf(C._SC_PAGE_SIZE)
	MemoryTotal = int(pages) * int(pageSize) / 1048576
	MemoryAvailable = MemoryTotal - mem2int/1024
}

func getPackages() {
	files, err := ioutil.ReadDir("/bin/")
	if err != nil {
		panic(err)
	}
	var packages int
	for num, _ := range files {
		packages = num
	}
	Packages = packages
}

func getCPU() {
	cpu, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}

	line := strings.Split(string(cpu), "\n")

	for _, line_str := range line {
		if strings.HasPrefix(line_str, "model name") {
			cpu_fields := strings.Fields(line_str)
			if len(cpu_fields) > 2 {
				CPU = strings.Join(cpu_fields[3:], " ")
				return
			}
		}
	}

}

func displayFetch() {
	getUser()
	getHostname()
	getKernel()
	getUptime()
	getMemory()
	getPackages()
	getCPU()

	s := strings.TrimSpace(User) + "@" + Hostname

	for i := 0; i < len(s); i++ {
		Lines += "-"
	}

	fmt.Printf(tmpl, Cyan, Bold, User, Reset, Cyan, Bold, Hostname, Reset,
		Lines,
		Cyan, Bold, Reset, runtime.GOOS, runtime.GOARCH,
		Cyan, Bold, Reset, Kernel,
		Cyan, Bold, Reset, Uptime,
		Cyan, Bold, Reset, Packages,
		Cyan, Bold, Reset, os.Getenv("SHELL"),
		Cyan, Bold, Reset, os.Getenv("XDG_SESSION_DESKTOP"),
		Cyan, Bold, Reset, os.Getenv("TERM"),
		Cyan, Bold, Reset, CPU,
		Cyan, Bold, Reset, MemoryAvailable, MemoryTotal)
}
