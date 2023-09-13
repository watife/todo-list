package main

import (
	"fmt"
	"github.com/watife/todo/todo"
	"os"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	var l todo.List

	err := l.Get(todoFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1:
		for _, item := range l {
			// List current items
			fmt.Println(item.Task)
		}
	default:
		// Add new item
		item := strings.Join(os.Args[1:], "")
		l = l.Add(item)
	}

	err = l.Save(todoFileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
