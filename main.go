package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var filename = "tasks.json"

type Task struct {
	Id          int
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
	loadTasks()
	defer saveTasks()

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

func lastTaskId() int {
	if len(tasks) == 0 {
		return 0
	}
	m := tasks[0].Id
	for _, task := range tasks {
		if task.Id > m {
			m = task.Id
		}
	}
	return m
}

func add(args []string) {

	id := lastTaskId() + 1
	t := Task{
		Id:          id,
		Description: args[0],
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, t)
}

func loadTasks() {
	if _, err := os.Stat(filename); err == nil {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if os.IsNotExist(err) {
		return
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

func saveTasks() {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := json.Marshal(tasks)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, err = file.Write(data)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
