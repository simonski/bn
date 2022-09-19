# README

## Introduction

This is the repository for Buildnumber, `bn` program.

It manages your version number via a number of different mechanisms.

You can call it directly, for example

```bash
bn command (-f FILENAME)
```

## Usage

### Get the current

```bash
bn get
```

### Upgrade the revision.

```bash
bn revision
```

### Upgrade the minor version (will reset revision to 0)

```bash
bn minor
```

### Upgrade the major version (will reset the minor.revision) to `0.0`

```bash
bn major
```
