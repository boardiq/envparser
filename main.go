package main

import (
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "envparser"
	app.Usage = "Take a config file and writes it to destination, interpolating any ENV strings."
	app.Version = "0.0.1"

	app.Action = func(c *cli.Context) {
		if exists := checkFile(c.Args().First()); !exists {
			println("No matching file, can't continue.")
			os.Exit(1)
		}
		println("File located, proceding.")
		candidate := c.Args()[0]
		content, _ := ioutil.ReadFile(candidate)
		newstring := os.ExpandEnv(string(content))

		ioutil.WriteFile(c.Args()[1], []byte(newstring), 0644)
		println("All done.")
	}

	app.Run(os.Args)
}

func checkFile(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
