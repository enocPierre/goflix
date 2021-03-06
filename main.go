package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Goflix")

	if err := run(); err != nil {
		if  err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
}

func run() error {
	srv := newServer()
	srv.store = &dbStore{}
	err := srv.store.Open()
	if  err != nil {
		return err
	}

	defer srv.store.Close()
	return nil
}