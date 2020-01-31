package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

var hold = make(chan int)

func main() {
	var pid string

	pidPtr := flag.String("p", "", "PID of the process to hold")
	namePtr := flag.String("n", "", "name of the process to hold")
	flag.Parse()

	pid = *pidPtr

	if *namePtr != "" && pid == "" {
		out, err := exec.Command("pgrep", *namePtr).Output()
		pid = strings.Replace(fmt.Sprintf("%s", out), "\n", "", -1)
		if err != nil || pid == "" {
			log.Fatalf("Invalid process name: \"%s\"", *namePtr)
		}
	}
	if pid != "" {
		go releaseOnLeave(pid)
		err := exec.Command("kill", "-TSTP", pid).Run()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("PID suspended: [%s]\n", pid)
		<-hold
	}

}

func releaseOnLeave(pid string) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		err := exec.Command("kill", "-CONT", pid).Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("PID released: [%s]", pid)
		os.Exit(0)
	}()
}
