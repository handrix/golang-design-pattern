package main

import "fmt"

func decorator(f func(string)) func(string) {
	return func(s string) {
		f(s)
	}
}

func run(name string) {
	fmt.Println(name + " run")
}

func main() {
	decorator(run)("zequn")
}
