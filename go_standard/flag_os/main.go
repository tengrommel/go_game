package main

import (
	"flag"
	"fmt"
	"os"

)

func main() {
	//name := flag.String("name", "Gopher", "name of the gopher")
	//age := flag.Int("age", 2, "age of the gopher")
	//shy := flag.Bool("shy", false, "is the gopher shy?")
	//flag.Parse()
	//fmt.Printf("Gopher Stats\nName: %s\nAge: %d\nShy: %t\n", *name, *age, *shy)

	var (
		name string
		age int
		shy bool
	)

	//flag.StringVar(&name, "name", "Gopher", "name of the gopher")
	//flag.StringVar(&name, "n", "Gopher", "name of the gopher (shorthand)")
	//flag.IntVar(&age, "age", 2, "age of the gopher")
	//flag.IntVar(&age, "a", 2, "age of the gopher(shorthand)")
	//flag.BoolVar(&shy, "shy", true, "is the gopher shy?")
	//flag.BoolVar(&shy, "s", true, "is the gopher shy(shorthand)?")
	//flag.Parse()
	//fmt.Printf("Gopher Stats\nName: %s\nAge: %dShy: %t\n", name, age, shy)
	flag.StringVar(&name, "name", defaultName(), "name of the gopher")
	flag.StringVar(&name, "n", "Gopher", "name of the gopher (shorthand)")
	flag.IntVar(&age, "age", 2, "age of the gopher")
	flag.IntVar(&age, "a", 2, "age of the gopher(shorthand)")
	flag.BoolVar(&shy, "shy", true, "is the gopher shy?")
	flag.BoolVar(&shy, "s", true, "is the gopher shy(shorthand)?")
	flag.Parse()
	fmt.Printf("Gopher Stats\nName: %s\nAge: %dShy: %t\n", name, age, shy)

}

func defaultName() string {
	if os.Getenv("GOPHER_DEFAULT_NAME") != ""{
		return os.Getenv("GOPHER_DEFAULT_NAME")
	}
	return "Gopher"
}
