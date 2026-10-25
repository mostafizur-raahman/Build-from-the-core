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
	count := 0
	notPid := 0
	for _, e := range entries {
		// fmt.Println(e)
		pid, err := strconv.Atoi(e.Name())
		if err != nil {
			continue
		}
		data, err := os.ReadFile(fmt.Sprintf("/proc/%d/comm", pid))
		if err != nil {
			notPid++
			continue
		}

		name := strings.TrimSpace(string(data))
		count++
		fmt.Println(count, pid, name)
	}
	fmt.Println("Total pid : ", count)
	fmt.Println("Not Pid ", notPid)
}
