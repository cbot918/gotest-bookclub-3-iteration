package main

import (
	"fmt"
	"os"
	"testing"
)

func setup() {
	// Do something here.
	fmt.Printf("\033[1;33m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	// Do something here.
	fmt.Printf("\033[1;33m%s\033[0m", "> Teardown completed")
	fmt.Printf("\n")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestHello(t *testing.T) {

	name := "yale"

	got := Hello(name)

	want := "hello yale"

	if got != want {
		t.Errorf("want %v got %v ", want, got)
	}

}
