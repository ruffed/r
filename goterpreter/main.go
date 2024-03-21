package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	ExOk      = 0  // exited correct
	ExUsage   = 64 // improper cmd usage
	ExDataErr = 65
	ExNoInput = 66 // input file was blank
	hadError  = false
)

// error handling.
func reportError(line int64, msg string) {
	report(line, "", msg)
}

func report(line int64, where string, msg string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error %s: %s\n", line, where, msg)

	hadError = true
}

func error(t LoxToken, message string) {
	if t.Type == Eof {
		report(t.Line, "at end", message)
	} else {
		report(t.Line, "at "+t.Lexeme, message)
	}
}

func run(source string) {
	l := NewLoxScanner(source)

	tokens := l.ScanTokens()

	for _, v := range tokens {
		fmt.Println(v)
	}
}

func runFile(b []byte) {
	run(string(b))

	if hadError {
		os.Exit(ExDataErr)
	}
}

func runPrompt() {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")

		s, err := r.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read line: %v\n", err)
			os.Exit(1)
		}

		if s == "\n" {
			os.Exit(ExOk)
		}

		s = strings.TrimRight(s, "\n")

		run(s)

		hadError = false
	}
}

func main() {
	if n := len(os.Args); n > 2 {
		fmt.Fprintln(os.Stderr, "Usage: goterpreter [file]")
		os.Exit(ExUsage)
	} else if n == 2 {
		bytes, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to open %s. (%v)\n", os.Args[1], err)
			os.Exit(ExNoInput)
		}

		runFile(bytes)
	} else {
		runPrompt()
	}
}
