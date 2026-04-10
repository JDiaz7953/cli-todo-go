package main

import "fmt"

func printHelp(){
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("     add [task]                  Add task to your list                       ")
	fmt.Println("     update [id] [update task]   update task based on the id         ")
	fmt.Println("     delete [id]                 delete task based on the id")
	fmt.Println("     mark-in-progress [id]       mark task as in-progress based on the id")
	fmt.Println("     mark-done [id]              mark task as done based on id")
	fmt.Println("     list                        list all task")
	fmt.Println("     list todo                   list unfinished task")
	fmt.Println("     list done                   list completed task")
	fmt.Println("     list in-progress            list task in progress")
}