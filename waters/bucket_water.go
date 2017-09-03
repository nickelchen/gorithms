package main

/*
There are three buckets, 1, 2, 3. their capacities are 8, 5, 3.
now the first bucket is full of water, which is 8. the others are empty.
There are no lines on these buckets,
you can only pour waters from source bucket to target bucket, which will cause:
	target bucket full,
	or source bucket empty.

Find a pour procedure, to result in waters in buckets are 4, 4, 0

*/

import (
	"fmt"
)

type DumpAction struct {
	from  int
	to    int
	water int
}

var buckets [3]int = [3]int{8, 5, 3}

type State struct {
	waters [3]int
	action DumpAction
}

func tryDump(path []State, state State, action DumpAction) (next State, ok bool) {
	waters := state.waters
	dumped := action.water

	// be careful, if state.waters is slice, you will change the previous slice
	// in our case, it's array, so we are modifing a copy of previous array
	waters[action.from] -= dumped
	waters[action.to] += dumped

	// check every bucket for waters.
	for i := 0; i < 3; i++ {
		if waters[i] > buckets[i] || waters[i] < 0 {
			return next, false
		}
	}

	next.waters = waters
	next.action = action

	// check loop
	for _, s := range path {
		if s.waters[0] == waters[0] &&
			s.waters[1] == waters[1] &&
			s.waters[2] == waters[2] {
			return next, false
		}
	}

	return next, true
}

func finishDump(state State) bool {
	waters := state.waters
	return waters[0] == 4 && waters[1] == 4
}

func run(path []State) {
	state := path[len(path)-1]
	if finishDump(state) {
		printPath(path)
		return
	}

	waters := state.waters
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			//
			canDumpWaters := []int{buckets[j] - waters[j], waters[i]}
			for _, water := range canDumpWaters {
				action := DumpAction{
					from:  i,
					to:    j,
					water: water,
				}
				if next, ok := tryDump(path, state, action); ok {
					path = append(path, next)
					run(path)
					path = path[0 : len(path)-1]
				}
			}

		}
	}
}

func printPath(path []State) {
	fmt.Println("Get result:")

	for _, state := range path {
		waters := state.waters
		action := state.action

		fmt.Printf("Pour from %d to %d with %d water, now waters are %v\n",
			action.from+1, action.to+1, action.water, waters)
	}
}
func main() {
	action := DumpAction{
		from:  -1,
		to:    1,
		water: 8,
	}
	state := State{
		waters: [3]int{8, 0, 0},
		action: action,
	}

	path := []State{state}
	run(path)
}
