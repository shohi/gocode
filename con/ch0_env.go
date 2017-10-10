package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func printEnv() {
	env := os.Environ()
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()
	fmt.Fprintf(tw, "Key\tValue\n")
	fmt.Println(os.Hostname())
	for _, raw := range env {
		pair := strings.Split(raw, "=")
		key := pair[0]
		value := pair[1]
		fmt.Fprintf(tw, "%s\t%s\n", key, value)
	}
}

func main() {
	printEnv()
}
