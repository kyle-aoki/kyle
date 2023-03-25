package main

import "fmt"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func must[T any](t T, err error) T {
	check(err)
	return t
}

func mainRecover() {
	if *flags.debug {
		return
	}
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
