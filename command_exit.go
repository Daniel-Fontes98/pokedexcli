package main

import "os"

func callbackExit(*Config, ...string) error {
	os.Exit(0)
	return nil
}