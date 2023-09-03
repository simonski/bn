# README

## Introduction

This is the repository for Buildnumber, `bn` program.

It manages your version number using semanitic versioning (major.minor.revision), via  `Buildnumber` file.  

Perhaps invoke it from your `Makefile` before you publish your code and you have a standard way of managing your versions.

## Usage

You can call it directly, for example

```bash
bn command (-f FILENAME)
```

### Initialise the current directory as your buildnumber dir

```bash
bn init
```

A `Buildnumber` file will be created.  This is the file that will store the version and be modified by the `bn` command from now on.

### Get the current version

```bash
bn get
```

### Upgrade the revision

```bash
bn revision
```

### Upgrade the minor (will reset revision to 0)

```bash
bn minor
```

### Upgrade the major version (will reset the minor.revision) to `0.0`

```bash
bn major
```

### Integration and automation

Now your program has a `Buildnumber` file, you can integrate the buildnumber into your CI/CD and client as you see fit.  

To see an example of integrating with go programs, try `bn help-go`


