package main

const DEFAULT_BUILDFILE = "Buildnumber"
const ENV_BUILDFILE = "BUILDNUMBER_FILE"
const COMMAND_GIT = "git"
const COMMAND_GET = "get"
const COMMAND_INIT = "init"
const COMMAND_UPGRADE_MAJOR = "major"
const COMMAND_VERSION = "version"
const COMMAND_UPGRADE_MINOR = "minor"
const COMMAND_UPGRADE_REVISION = "revision"
const COMMAND_HELP = "help"
const COMMAND_HELP_GO = "help-go"

const USAGE = `
Usage: bn <command> <options> (-f <buildfile>)

Commands
	init                            - initialises a Buildnumber file

	get                             - returns the version of YOUR application
	revision                        - upgrades the revision
	major                           - upgrades the major version
	minor                           - upgrades the minor version
	git                             - returns the git hash/branch

	help-go                         - how to integrate with a go program
	version                         - displays the INTERNAL bn version 

`

const USAGE_HELP_GO = `bn can be integrated with your go program with three steps:

1. In your Makefile "build" step, add the "bn revision" command
build:
	bn revision
	go build

2. Add to your main.go:
	//go:embed Buildnumber
	var Buildnumber embed.FS

	// Version returns the buildnumber version of the application
	// to be invoked by the user on the terminal using  "./app version":
	func Version() {
		data, _ := Buildnumber.ReadFile("Buildnumber")
		v := string(data)
		v = strings.ReplaceAll(v, "\n", "")
		fmt.Println(v)
	}

3. In the root dir of your application
	./bn init
	make build

Your program will now increment the revision on every "make build" invocation and make
the value available via the embedded "Buildnumber" file.

`
