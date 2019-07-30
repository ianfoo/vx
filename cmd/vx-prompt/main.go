package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ianfoo/vx"
)

func main() {
	const msg = "VX: Only a little bit of nerve damage"
	fmt.Printf("%s\n%s\n", msg, strings.Repeat("=", len(msg)))

	s := bufio.NewScanner(os.Stdin)
	for {
		proceed, err := prompt("Do you wish to continue the VX? ", s)
		if !proceed || err != nil {
			exit, err := prompt("Are you sure you wish to exit the VX? ", s)
			if exit || err != nil {
				break
			}
		}
	}
}

func prompt(p string, s *bufio.Scanner) (bool, error) {
	for {
		fmt.Print(p)
		if !s.Scan() {
			if err := s.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				return false, err
			}
			return false, io.EOF
		}
		proceed, valid := vx.Proceed(s.Text())
		if !valid {
			fmt.Println("Please enter a valid response.")
			continue
		}
		return proceed, nil
	}
}
