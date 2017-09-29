package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	strs := strings.Fields(text)
	n, _ := strconv.Atoi(strs[0])
	array_input, _ := reader.ReadString('\n')
	a := strings.Fields(array_input)

	b := make([]int, n)
	left_step, _ := strconv.Atoi(strs[1])
	for i := 0; i < n; i++ {
		j := i - left_step
		if j < 0 {
			j = j + n
		}
		b[j], _ = strconv.Atoi(a[i])
	}
	fmt.Println("After rotations")
	fmt.Printf("%v", b)
}
