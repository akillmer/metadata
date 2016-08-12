package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/andykillmer/go-dcraw-json"
	"os"
	"strings"
)

type Task struct {
	Id       int    `json:"id"`
	Filename string `json:"filename"`
}

func main() {
	cmdPath := strings.TrimSuffix(os.Args[0], "metadata") + "dcraw-json"
	if err := dcraw.Path(cmdPath); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//go runTasks()

	task := Task{}
	scanner := bufio.NewScanner(os.Stdin)
	args := []string{"-i", "-v", "-J"}
	buffer := new(bytes.Buffer)

	for scanner.Scan() {
		input := scanner.Bytes()

		if err := json.Unmarshal(input, &task); err != nil {
			fmt.Fprintf(os.Stderr, "Could not unmarshal task: %s\n", err)
			fmt.Println(string(input))
			continue
		}

		buffer.Reset()
		err := dcraw.Run(append(args, task.Filename), buffer)

		if buffer.Len() == 0 {
			fmt.Fprintf(os.Stderr, "{\"id\":%d, \"error\":%s, \"response\":%s}\n", task.Id, "\""+err.Error()+"\"", "null")
		} else {
			fmt.Fprintf(os.Stdout, "{\"id\":%d, \"error\":%s, \"response\":%s}\n", task.Id, "null", string(buffer.Bytes()))
		}
	}
}
