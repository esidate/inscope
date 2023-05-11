package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	var scopeFile, domainsFile string
	var domainsReader io.Reader

	switch len(args) {
	case 2:
		scopeFile = args[0]
		domainsFile = args[1]
		domainsReader, _ = os.Open(domainsFile)
	case 1:
		scopeFile = args[0]
		domainsReader = os.Stdin
	default:
		fmt.Println("Usage: inscope <scope-file> [<domains-file>]")
		fmt.Println("If no domains file is provided, domains are read from stdin.")
		fmt.Println("Examples:")
		fmt.Println("inscope scope.txt domains.txt")
		fmt.Println("cat domains.txt | inscope scope.txt")
		os.Exit(1)
	}

	scopeDomains, err := readLines(scopeFile)
	if err != nil {
		fmt.Println("Error reading scope file:", err)
		os.Exit(1)
	}

	var domains []string
	scanner := bufio.NewScanner(domainsReader)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if isInScope(scopeDomains, domain) {
			domains = append(domains, domain)
		} else {
			fmt.Println("Deleting", domain)
		}
	}

	if domainsFile != "" {
		writeLines(domains, domainsFile)
	} else {
		for _, domain := range domains {
			fmt.Println(domain)
		}
	}
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	return lines, scanner.Err()
}

func isInScope(scopeDomains []string, domain string) bool {
	for _, scopeDomain := range scopeDomains {
		if strings.HasSuffix(domain, scopeDomain) {
			return true
		}
	}
	return false
}

func writeLines(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}
