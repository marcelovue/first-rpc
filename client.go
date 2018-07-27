package main

import (
	"fmt"
	"log"
	"net/rpc"
	// "encoding/json"
	// "io/ioutil"
	// "bytes"
	// "reflect"
)

type Task struct {
	Id          string
	Title       string
	Description string
}

func main() {
	var reply []Task

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Custom fatal:: ", err)
	}
	first := Task{Id: "1", Title: "Walk the dog", Description: "At 6AM"}
	second := Task{Id: "2", Title: "Feed the dog", Description: "At 6:30AM"}
	third := Task{Id: "3", Title: "Sedate the dog", Description: "At 12AM"}

	err = client.Call("Task.SetTask", first, &reply)
	err = client.Call("Task.SetTask", second, &reply)
	err = client.Call("Task.SetTask", third, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
	// task   := new(Task)
	// first  := Task{ id:"1",title:"Walk the dog",   description:"At 6AM"  }
	// second := Task{ id:"2",title:"Clean the dog",  description:"At 6PM"  }
	// third  := Task{ id:"3",title:"Sedate the dog", description:"At 11PM" }

	// task.SetTask(first)
	// task.SetTask(second)
	// task.SetTask(third)
	// task.SetTitle(3, "I can't do that")
	// fmt.Println(task.GetAllTasks())
}
