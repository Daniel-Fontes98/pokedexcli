package main

import "os"

func callbackExit(c *Config) error {
	os.Exit(0)
	return nil
}