package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/aerissecure/pulid"
)

func main() {
	// Define the count flag
	count := flag.Int("n", 1, "Number of IDs to generate (default: 1)")

	// Customize the help message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s PREFIX [-count N]\n\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Arguments:")
		fmt.Fprintln(os.Stderr, "  PREFIX      Prefix to prepend to the ID (required, positional)")
		fmt.Fprintln(os.Stderr, "\nFlags:")
		flag.PrintDefaults()
	}

	// Parse the flags
	flag.Parse()

	// Validate positional argument for PREFIX
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Error: PREFIX argument is required.")
		flag.Usage()
		os.Exit(1)
	}
	prefix := strings.ToUpper(flag.Arg(0)) // Get the first positional argument and make it uppercase

	// Validate count (optional)
	if *count < 1 {
		fmt.Fprintln(os.Stderr, "Error: -count must be greater than 0.")
		os.Exit(1)
	}

	// Generate and print IDs
	for i := 0; i < *count; i++ {
		id := pulid.MustNew(prefix)
		fmt.Println(id)
	}
}
