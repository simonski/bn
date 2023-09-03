package main

import (
	"embed"
	"fmt"
	"os"
	"strings"

	git "github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	cli "github.com/simonski/cli"
	"github.com/simonski/goutils"
)

//go:embed Buildnumber
var Buildnumber embed.FS

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
	} else if command == COMMAND_VERSION {
		MyVersion()
	} else if command == COMMAND_UPGRADE_REVISION {
		UpgradeRevision(cli)
	} else if command == COMMAND_HELP || command == "" {
		Help(cli)
	} else if command == COMMAND_HELP_GO {
		HelpGo(cli)
	} else if command == COMMAND_GIT {
		fmt.Println(GitInfo(cli))

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
	filename, err := GetFilename(c)
	if goutils.FileExists(filename) {
		fmt.Printf("File `%v` already exists, cannot init.\n", filename)
		os.Exit(1)
	}
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	version := NewFromFile(filename)
	// version := Load(c)
	version.Save(version.Filename)
	fmt.Printf("Initialised `%v` successfully.\n", filename)
}

func Help(c *cli.CLI) {
	fmt.Print(USAGE)
}

func HelpGo(c *cli.CLI) {
	fmt.Print(USAGE_HELP_GO)
}

func GitInfo(c *cli.CLI) string {
	dir := c.GetStringOrDefault("-dir", ".")
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return ""
	}
	h, err := repo.Head()
	if err != nil {
		panic(err)
	}
	return h.String()
}

func MyVersion() {
	data, _ := Buildnumber.ReadFile("Buildnumber")
	v := string(data)
	v = strings.ReplaceAll(v, "\n", "")
	fmt.Println(v)
}
