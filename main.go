package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct {
	id          int
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

var commands map[string]func([]string) = map[string]func([]string){
	"add":              add,
	"update":           nil,
	"delete":           nil,
	"mark-in-progress": nil,
	"mark-done":        nil,
	"list":             nil,
}

var tasks []Task = []Task{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
this is task tracker cli and list and commands are :
	add
	delete
	update
	mark-in-progress
	mark-done
	list		
		`)
		os.Exit(0)
	}

	if _, ok := commands[os.Args[1]]; !ok {
		fmt.Println("command does not exists")
		os.Exit(1)
	}
	commands[os.Args[1]](os.Args[2:])
}

func add(args []string) {

}
