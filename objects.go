package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	cli "github.com/simonski/cli"
	"github.com/simonski/goutils"
)

type Version struct {
	Major    int
	Minor    int
	Revision int
	Filename string
}

func (v *Version) String() string {
	return fmt.Sprintf("%v.%v.%v", v.Major, v.Minor, v.Revision)
}

func (v *Version) IncrementRevision() *Version {
	v.Revision += 1
	return v
}

func (v *Version) IncrementMinor() *Version {
	v.Minor += 1
	v.Revision = 0
	return v
}

func (v *Version) IncrementMajor() *Version {
	v.Major += 1
	v.Minor = 0
	v.Revision = 0
	return v
}

func New(value string) *Version {
	splits := strings.Split(value, ".")
	major_s := splits[0]
	minor_s := splits[1]
	revision_s := splits[2]
	major, _ := strconv.Atoi(major_s)
	minor, _ := strconv.Atoi(minor_s)
	revision, _ := strconv.Atoi(revision_s)
	return &Version{Major: major, Minor: minor, Revision: revision}
}

func NewFromFile(filename string) *Version {
	if goutils.FileExists(filename) {
		s := goutils.Load_file_to_strings(filename)
		line := s[0]
		v := New(line)
		v.Filename = filename
		return v
	} else {
		v := New("0.0.0")
		v.Filename = filename
		return v
	}
}

func Load(c *cli.CLI) *Version {
	filename := c.GetFileExistsOrDefault("-f", DEFAULT_BUILDFILE)
	return NewFromFile(filename)
}

func (v *Version) Save(filename string) {
	line := v.String()
	f, err := os.Create(filename)
	if err != nil {
		print(err)
	}
	defer f.Close()
	f.WriteString(line)
	f.WriteString("\n")
}
