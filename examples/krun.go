/*
Tornado ignore signal CTRL-C, so I write this program
*/
package main

import (
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/codeskyblue/kexec"
)

func main() {
	cmd := kexec.CommandString(strings.Join(os.Args[1:], " ")) //(os.Args[1], os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)

	go func() {
		for _ = range sigC {
			log.Println("Signal interrupt catched. terminate")
			cmd.Terminate(syscall.SIGKILL)
		}
	}()

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
