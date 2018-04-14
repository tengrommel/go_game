package logger

import (
	"sync"
	"io"
	"fmt"
)

type Logger struct {
	write chan string
	wg sync.WaitGroup
}

func New(w io.Writer, capacity int) *Logger {
	l := Logger{
		write:make(chan string, capacity),
	}
	l.wg.Add(1)
	go func() {
		for d:=range l.write{
			fmt.Fprintln(w, d)
		}
		l.wg.Done()
	}()
	return &l
}

func (l *Logger)Shutdown()  {
	close(l.write)
	l.wg.Wait()
}

func (l *Logger)Write(data string)  {
	select {
	case l.write <- data:
	default:
		fmt.Println("Dropping the write")
	}
}
