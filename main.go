package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Goflix")

	srv := newServer()
	srv.store = &dbStore{}
	err := srv.store.Open()
	if  err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	defer srv.store.Close()
}
