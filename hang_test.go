package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	if os.Getenv("HANG_SUBPROCESS") == "true" {
		time.Sleep(5 * time.Second)
		fmt.Println("HANG")
		return
	}
	m.Run()
}

func TestHang(t *testing.T) {
	cmd := exec.Command(os.Args[0], "-test.run=TestHang", "-test.v")
	cmd.Env = append(os.Environ(), "HANG_SUBPROCESS=true")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
}
