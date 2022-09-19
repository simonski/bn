package main

import (
	"fmt"
	"os"

	cli "github.com/simonski/cli"
)

var VALID_COMMANDS = [...]string{
	COMMAND_INIT,
	COMMAND_GET,
	COMMAND_UPGRADE_MAJOR,
	COMMAND_UPGRADE_MINOR,
	COMMAND_UPGRADE_REVISION,
}

func IsValidCommand(command string) bool {
	for _, e := range VALID_COMMANDS {
		if e == command {
			return true
		}
	}
	return false
}

func main() {
	cli := cli.New(os.Args)
	cli.Shift()
	command := cli.GetCommand()

	if command == COMMAND_GET {
		Get(cli)
	} else if command == COMMAND_INIT {
		Init(cli)
	} else if command == COMMAND_UPGRADE_MAJOR {
		UpgradeMajor(cli)
	} else if command == COMMAND_UPGRADE_MINOR {
		UpgradeMinor(cli)
	} else if command == COMMAND_UPGRADE_REVISION {
		UpgradeRevision(cli)
	} else if command == COMMAND_HELP || command == "" {
		Help(cli)
	} else {
		fmt.Printf("Error, '%v' not found.\n", command)
		os.Exit(1)
	}

}

func Get(c *cli.CLI) {
	v := Load(c)
	fmt.Println(v.String())
}

func UpgradeMajor(c *cli.CLI) {
	v := Load(c)
	v.IncrementMajor()
	v.Save(v.Filename)
	fmt.Println(v.String())
}

func UpgradeMinor(c *cli.CLI) {
	v := Load(c)
	v.IncrementMinor()
	v.Save(v.Filename)
	fmt.Println(v.String())
}

func UpgradeRevision(c *cli.CLI) {
	v := Load(c)
	v.IncrementRevision()
	v.Save(v.Filename)
	fmt.Println(v.String())
}

func Init(c *cli.CLI) {
	version := Load(c)
	version.Save(version.Filename)
}

func Help(c *cli.CLI) {
	fmt.Print(USAGE)
}
