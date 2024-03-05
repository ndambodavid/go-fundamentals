package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cli() {
	fmt.Println("what do you want to show?")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	fmt.Println(s + "!")
}
