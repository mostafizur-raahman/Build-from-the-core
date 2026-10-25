package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	entries, err := os.ReadDir("/proc")
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		pid, err := strconv.Atoi(e.Name())
		if err != nil {
			continue
		}

		data, err := os.ReadFile(fmt.Sprintf("/proc/%d/comm", pid))
		if err != nil {
			continue
		}

		name := strings.TrimSpace(string(data))

		code, err := readState(pid)
		if err != nil {
			panic(err)
		}
		status, err := readStatus(pid)
		if err != nil {
			continue
		}

		ppid := status["PPid"]
		vmSize := status["VmSize"]
		vmRSS := status["VmRSS"]

		fds, err := countFDs(pid)
		if err != nil {
			fds = -1
		}

		execN, writeN, readN, err := countRegions(pid)
		if err != nil {
			execN, writeN, readN = -1, -1, -1 // no permission
		}

		fmt.Println(pid, name, "-", describeState(code), "| ppid:", ppid, "| vmsize:", vmSize, "| rss:", vmRSS, "| fds:", fds, "| r-x:", execN, "| rw-:", writeN, "| r--:", readN)

	}
}

func readStatus(pid int) (map[string]string, error) {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/status", pid))
	if err != nil {
		return nil, err
	}

	fields := map[string]string{}
	for _, line := range strings.Split(string(data), "\n") {
		key, value, found := strings.Cut(line, ":")
		if found {
			fields[key] = strings.TrimSpace(value)
		}
	}
	return fields, nil
}

func countFDs(pid int) (int, error) {
	entries, err := os.ReadDir(fmt.Sprintf("/proc/%d/fd", pid))
	if err != nil {
		return 0, err
	}
	return len(entries), nil
}

func readState(pid int) (string, error) {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/status", pid))

	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "State:") {
			value := strings.TrimSpace(strings.TrimPrefix(line, "State:"))
			return value[:1], nil
		}
	}

	return "", fmt.Errorf("State line not found")
}

func describeState(code string) string {
	switch code {
	case "R":
		return "Running or ready to run"
	case "S":
		return "Sleeping, waiting for an event"
	case "D":
		return "Waiting on disk or I/O (can't be interrupted)"
	case "T":
		return "Stopped"
	case "t":
		return "Stopped by a debugger"
	case "Z":
		return "Zombie (finished, parent hasn't collected it)"
	case "I":
		return "Idle (kernel thread)"
	default:
		return "Unknown"
	}
}

func countRegions(pid int) (execN, writeN, readN int, err error) {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/maps", pid))
	if err != nil {
		return 0, 0, 0, err
	}

	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		switch fields[1][:3] {
		case "r-x":
			execN++
		case "rw-":
			writeN++
		case "r--":
			readN++
		}
	}
	return execN, writeN, readN, nil
}
