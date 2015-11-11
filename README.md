# kproc
[![GoDoc](https://godoc.org/github.com/codeskyblue/kexec?status.svg)](https://godoc.org/github.com/codeskyblue/kexec)

This is a golang lib, add a `Terminate` command to exec.

Tested on _windows, linux, darwin._

This lib has been used in [fswatch](https://github.com/codeskyblue/fswatch).

## Usage

	go get -v github.com/codeskyblue/kexec

example: see more [examples](examples)

	package main
	
	import "github.com/codeskyblue/kexec"

	func main() {
		p := kexec.CommandString("python flask_main.py")
		p.Start()
		p.Terminate(syscall.SIGKILL)
	}
