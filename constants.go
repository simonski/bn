package main

const DEFAULT_BUILDFILE = "Buildnumber"
const ENV_BUILDFILE = "BUILDNUMBER_FILE"
const COMMAND_GIT = "git"
const COMMAND_GET = "get"
const COMMAND_INIT = "init"
const COMMAND_UPGRADE_MAJOR = "major"
const COMMAND_UPGRADE_MINOR = "minor"
const COMMAND_UPGRADE_REVISION = "revision"
const COMMAND_HELP = "help"

const USAGE = `
Usage: bn <command> <options> (-f <buildfile>)

Commands
	get                             - returns the version
	init                            - initialises a Buildnumber file
	major                           - upgrades the major version
	minor                           - upgrades the minor version
	revision                        - upgrades the revision
	git                             - returns the git hash/branch

`
