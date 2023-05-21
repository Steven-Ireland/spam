package main

import (
	"flag"
	"fmt"
	"os/exec"
	"time"
)

type Status string

const (
	SUCCESS Status = "success"
	ERROR   Status = "error"
)

type ProgramReport struct {
	duration time.Duration
	status   Status
}

func main() {
	rate := flag.Int("r", 1, "Number of times to execute per second")

	flag.Parse()

	command := flag.Args()[0]
	args := flag.Args()[1:]

	reported := make(chan ProgramReport)

	go report(reported)

	for {
		milliseconds := 1000 * float64(1.0) / float64(*rate)
		time.Sleep(time.Duration(milliseconds) * time.Millisecond)

		go run(command, args, reported)
	}
}

func run(command string, args []string, channel chan<- ProgramReport) {
	start := time.Now()

	_, err := exec.Command(command, args...).Output()

	end := time.Now()

	status := SUCCESS
	if err != nil {
		status = ERROR
	}

	channel <- ProgramReport{duration: (end.Sub(start)), status: status}
}

// Keeps a running average of execution times.
// Future work: Output running average at program close
func report(channel <-chan ProgramReport) {
	numReports := 0
	average := 0.0

	for {
		report := <-channel

		numReports = numReports + 1

		if numReports > 1 {
			average = average * ((float64(numReports) - 1) / float64(numReports))
			average = average + float64(report.duration.Milliseconds())/float64(numReports)
		} else {
			average = float64(report.duration.Milliseconds())
		}

		fmt.Printf("%d\n", report.duration.Milliseconds())
	}
}
