package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Who has Ball eating problems?: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	myVarOne := name + " Sucks Balls..."
	fmt.Println("Why does "+name+"'s Mouth smell?", myVarOne)
}
