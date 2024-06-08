package cmd

const (
	Cyan  = "\033[36m"
	Bold  = "\033[1m"
	Reset = "\033[0m"
)

const tmpl = `
%s%s%s%s@%s%s%s%s
%s
%s%sOS:%s %s %s
%s%sKernel:%s %s
%s%sUptime:%s %d mins
%s%sPackages:%s %d
%s%sShell:%s %s
%s%sWM:%s %s
%s%sTerminal:%s %s
%s%sCPU:%s %s
%s%sMemory:%s %d MiB / %d MiB

`
