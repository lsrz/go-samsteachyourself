package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

//创建子命令
func flagUsage() {
	usageText := `this is an example cli tool .
Usage:
c173命令行flag command [arguments]
The commands are :
uppercase uppercase a string
lowercase lowercase a string
Use "example05 [command] --help " for more information about a command`
	fmt.Fprintf(os.Stderr, " %s \n", usageText)
}

func main() {

	flag.Usage = flagUsage
	uppercaseCmd := flag.NewFlagSet("uppercase", flag.ExitOnError)
	lowercaseCmd := flag.NewFlagSet("lowercase", flag.ExitOnError)

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "uppercase":
		a := uppercaseCmd.String("s", "", "A string of text to be uppercased")
		uppercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToUpper(*a))
	case "lowercase":
		a := lowercaseCmd.String("s", "", "A string of text to be lowercased")
		lowercaseCmd.Parse(os.Args[2:])
		fmt.Println(strings.ToLower(*a))
	default:
		flag.Usage()
	}
	/*
		go run c173命令行flag.go uppercase -s "i want to grow up"
		I WANT TO GROW UP

		go run c173命令行flag.go lowercase -s "I DO NOT WANT TO GROW UP"
		i do not want to grow up
	*/
}
