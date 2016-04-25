package main

import (
	"fmt"
	"github.com/adammck/instance_vars/providers/gcp"
	"os"
)

type Provider interface {
	Get() (map[string]string, error)
}

func main() {
	var prv Provider
	prv = gcp.New()
	vars, err := prv.Get()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}

	for k, v := range vars {
		fmt.Printf("%s=%s\n", k, v)
	}
}
