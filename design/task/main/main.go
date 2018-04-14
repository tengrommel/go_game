package main

import (
	"log"
	"time"
	"github.com/tengrommel/go_game/design/task"
	"sync"
)

// names provides a set of names to display.
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter provides special support for printing names.
type namePrinter struct {
	name string
}

func (m namePrinter)Work()  {
	log.Println(m.name)
	time.Sleep(3*time.Second)
}

func main() {
	const routines = 10

	// Create a task pool.
	t := task.New(routines)

	var wg sync.WaitGroup
	wg.Add(routines * len(names))

	for i:=0;i<routines;i++ {
		// Iterate over the slice of names.
		for _, name := range names{
			// Create a namePrinter and provide the specific name.
			np := namePrinter{
				name:name,
			}
			go func() {
				// Submit the task to be worked on. When Do returns, we know it is being handled.
				t.Do(np)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	t.Shutdown()
}
