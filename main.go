package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

const baseConfigPath = "/home/aym/xconfigs"

var configs = []string{
	"xmonad",
	"qtile",
}

func getConfigPaths(cNames []string) map[string]string {
	result := map[string]string{}
	// Convert all the config names to its configuration file path.
	for _, c := range cNames {
		result[c] = filepath.Join(
			baseConfigPath,
			fmt.Sprintf("%s.xinitrc", c),
		)
	}
	return result
}

func printConfigOptions(cNames []string) {
	fmt.Println("Chose your wm:\n")
	for i, c := range cNames {
		fmt.Printf("%d) -> %s\n", i, c)
	}
}

func printTryAgainPrompt() {
	fmt.Println("Wrong option, try again...")
}

func main() {
	paths := getConfigPaths(configs)
	printConfigOptions(configs)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n -> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read")
			printTryAgainPrompt()
			continue
		}

		// Delete the newline.
		text = text[:len(text)-1]

		confIdx, err := strconv.Atoi(text)
		if err != nil || confIdx < 0 || confIdx >= len(configs) {
			// There was an error or the index falls out of the valid options.
			printTryAgainPrompt()
			continue
		}
		confName := configs[confIdx]
		confPath := paths[confName]
		cmd := exec.Command("startx", confPath)

		err = cmd.Run()

		if err != nil {
			fmt.Println("Done! try another one")
		} else {
			fmt.Println("Wm exited incorrectly! try another one")
		}
	}
}
