package main

import (
	"PM/input"
	"PM/lock"
	"PM/process"
	"PM/storage"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	stg := storage.Storage{}
	stg.Init(10)

	inp := input.Input{}.Init()

	lock.Init()

	for true {
		fmt.Print("> ")
		cmd := strings.Split(inp.Get(), " ")

		switch {
		case cmd[0] == "new":
			ID, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}

			delay, err := strconv.Atoi(cmd[3])
			if err != nil {
				panic(err)
			}

			stg.Add(process.Process{
				PID:       int32(ID),
				Delay:     int32(delay),
				Task:      cmd[1],
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Terminate: false,
			})
		case cmd[0] == "kill":
			ID, err := strconv.Atoi(cmd[1])
			if err != nil {
				panic(err)
			}

			stg.Kill(int32(ID))
		case cmd[0] == "monitor":
			stg.View()
		case cmd[0] == "terminate":
			break
		}
	}
}