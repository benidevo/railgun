package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.1.0-dev"

func main() {
	var (
		showVersion = flag.Bool("version", false, "Show version information")
		showHelp    = flag.Bool("help", false, "Show help information")
	)

	flag.Parse()

	if *showVersion {
		fmt.Printf("railgun version %s\n", version)
		os.Exit(0)
	}

	if *showHelp {
		printHelp()
		os.Exit(0)
	}

	fmt.Println("Railgun - High-performance streaming data pipeline for PostgreSQL")
	fmt.Println("Under development. Use --help for more information.")
}

func printHelp() {
	fmt.Println("Railgun - High-performance streaming data pipeline for PostgreSQL")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  railgun [options]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  --version    Show version information")
	fmt.Println("  --help       Show this help message")
	fmt.Println()
	fmt.Println("This tool is under active development.")
	fmt.Println("Visit https://github.com/benidevo/railgun for more information.")
}
