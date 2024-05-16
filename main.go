package main

import (
	"fmt"
	"os"
	"strconv"

	"ROMANS/romans"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The program requires an input at os.Args[1]")
		return
	}
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error during number conversion to int type")
		return
	}
	if input > 3999 {
		fmt.Println("The program requires an input not higher than 3999")
		return
	}
	fmt.Printf("%s\n", romans.Romans(input))
	fmt.Printf("%s\n", romans.IntToRoman(input))
}
