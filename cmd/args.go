package cmd

import (
	"flag"
	"fmt"
	"os"
)

var (
	UserFlag     bool
	HostnameFlag bool
	KernelFlag   bool
	UptimeFlag   bool
	PackagesFlag bool
	ShellFlag    bool
	WMFlag       bool
	TerminalFlag bool
	CPUFlag      bool
	MemoryFlag   bool
)

func Args() {
	flag.BoolVar(&UserFlag, "u", false, "display the current user name")
	flag.BoolVar(&HostnameFlag, "h", false, "display the current hostname")
	flag.BoolVar(&KernelFlag, "k", false, "display the current kernel")
	flag.BoolVar(&UptimeFlag, "t", false, "display the uptime")
	flag.BoolVar(&PackagesFlag, "p", false, "display the amount of packages")
	flag.BoolVar(&ShellFlag, "s", false, "display the current terminal shell")
	flag.BoolVar(&WMFlag, "d", false, "display the current WM or DE")
	flag.BoolVar(&TerminalFlag, "c", false, "display the current terminal")
	flag.BoolVar(&CPUFlag, "i", false, "display the current CPU")
	flag.BoolVar(&MemoryFlag, "m", false, "display the information about RAM")
	flag.Parse()

	if UserFlag {
		getUser()
		fmt.Printf("%s\n", User)
		return
	}

	if HostnameFlag {
		getHostname()
		fmt.Printf("%s\n", Hostname)
		return
	}

	if KernelFlag {
		getKernel()
		fmt.Printf("%s\n", Kernel)
		return
	}

	if UptimeFlag {
		getUptime()
		fmt.Printf("%d mins\n", Uptime)
		return
	}

	if PackagesFlag {
		getPackages()
		fmt.Printf("%d\n", Packages)
		return
	}

	if ShellFlag {
		fmt.Printf("%s\n", os.Getenv("SHELL"))
		return
	}

	if WMFlag {
		fmt.Printf("%s\n", os.Getenv("XDG_SESSION_DESKTOP"))
		return
	}

	if TerminalFlag {
		fmt.Printf("%s\n", os.Getenv("TERM"))
		return
	}

	if CPUFlag {
		getCPU()
		fmt.Printf("%s\n", CPU)
		return
	}

	if MemoryFlag {
		getMemory()
		fmt.Printf("%d MiB / %d MiB\n", MemoryAvailable, MemoryTotal)
		return
	}

	nonFlagsArgs := flag.Args()
	if len(nonFlagsArgs) == 0 {
		displayFetch()
		return
	}
}
