package main

import (
	"bufio"
	"fmt"
	"os"
)

type Change struct {
	name string
	id   int
	time int
}

type System struct {
	history []Change
}

func NewSystem() *System {
	return &System{
		history: []Change{},
	}
}

func (s *System) Change(name string, time, id int) {
	s.history = append(s.history, Change{name, id, time})
}

func (s *System) Get(id, time int) string {
	var nameAtTime string
	for _, change := range s.history {
		if change.id == id && change.time <= time {
			nameAtTime = change.name
		}
		fmt.Println(change.id, change.time, "------", nameAtTime)
	}
	return nameAtTime
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	system := NewSystem()
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		curTime := 1
		for j := 0; j < n; j++ {
			var str string
			fmt.Fscan(in, &str)
			if str == "CHANGE" {
				var name string
				var id int
				fmt.Fscan(in, &name, &id)
				system.Change(name, curTime, id)
			} else if str == "GET" {
				var id, time int
				fmt.Fscan(in, &id, &time)
				res := system.Get(id, time)
				fmt.Fprintln(out, res)
			}
			curTime++
		}
	}
}
