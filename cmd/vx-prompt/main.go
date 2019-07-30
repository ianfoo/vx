package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ianfoo/vx"
)

func main() {
	const msg = "VX: Only a little bit of nerve damage"
	fmt.Printf("%s\n%s\n", msg, strings.Repeat("=", len(msg)))

	s := bufio.NewScanner(os.Stdin)
	for {
		if !prompt("Do you wish to continue the VX? ", s) {
			if prompt("Are you sure you wish to exit the VX? ", s) {
				return
			}
		}
	}
}

func prompt(p string, s *bufio.Scanner) bool {
	for {
		fmt.Print(p)
		if !s.Scan() {
			return false
		}
		proceed, valid := vx.Proceed(s.Text())
		if !valid {
			fmt.Println("Please enter a valid response.")
			continue
		}
		return proceed
	}
}
