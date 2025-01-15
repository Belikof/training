package main

import ("fmt"
		"os"
)

// func getPassedArgs(minArgs int) []string {
// 	if len(os.Args) < minArgs {
// 		fmt.Printf("At least %v arguments are needed\n", minArgs)
// 		os.Exit(1)
// 	}
// 	var args []string
// 	for i := 1; i < len(os.Args); i++ {
// 		args = append(args, os.Args[i])
// 	}
// 	return args
// }

// func findLongest(args []string) string {
// 	var longest string
// 	for i := 0; i < len(args); i++ {
// 		if len(args[i]) > len(longest) {
// 			longest = args[i]
// 		}
// 	}
// 	return longest
// }

// func main() {
// 	if longest := findLongest(getPassedArgs(3)); len(longest) > 0 {
// 		fmt.Println("The longest word passed was:", longest)
// 	} else {
// 		fmt.Println("There was an error")
// 		os.Exit(1)
// 	}
// }

func getPassedArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func getLocals(extraLocals []string) []string {
	var locales []string
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocals...)
	return locales
}

func main() {
	locales := getLocals(getPassedArgs())
	fmt.Println("Locales to use:", locales)
}