package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strconv"
)

type TaskID string

type Task struct {
	Id          string
	Title       string
	Description string
}

var tasks []Task

func (t *Task) GetAllTasks(task *Task, reply *[]Task) error {
	*reply = tasks
	return nil
}
func (t *Task) SetTask(task Task, reply *[]Task) error {
	fmt.Println("I was asked for set this task", task)
	tasks = append(tasks, task)
	*reply = tasks
	return nil
}
func (t *Task) SetTitle(task *Task, reply *[]Task) error {
	id, _ := strconv.Atoi(task.Id)
	tasks[id].Title = task.Title
	*reply = tasks
	return nil
}
func main() {
	task := new(Task)
	err := rpc.Register(task)
	if err != nil {
		log.Fatal("Puts meu deu erro no register", err)
	}
	rpc.HandleHTTP()
	listener, err2 := net.Listen("tcp", ":1234")
	if err2 != nil {
		log.Fatal("Deu erro, só que numero 2", err2)
	}

	err3 := http.Serve(listener, nil)
	if err3 != nil {
		log.Fatal("Nao deu pra abrir o servidor", err3)
	}
}

// func (t Task) GetTask(id TaskID) (Task, error) {
// 	var tasks
// 	return
// }

// func DeleteTask(id TaskID) (bool, error) {

// }
// func EditTask(id TaskID, data Task) (bool, error) {

// }
