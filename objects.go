package main

import (
	"errors"
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

// Findfile looks in the current directory then "walks" upwards
// until it either finds a file matching the name or stops at $HOME
// If a file is not found, filename is returned as-is
func Findfile(filename string, VERBOSE bool) string {
	// } else {
	home := os.Getenv("HOME")
	if VERBOSE {
		fmt.Printf("Home=%v\n", home)
	}
	path, _ := os.Getwd()
	if VERBOSE {
		fmt.Printf("cur_dir   : %v\n", path)
	}

	new_path := path
	for {
		candidate := new_path + "/" + filename
		if goutils.FileExists(candidate) {
			if VERBOSE {
				fmt.Printf("candidate : %v [EXIST!]\n", candidate)
			}
			return candidate
		}
		if VERBOSE {
			fmt.Printf("candidate : %v [NOT EXIST]\n ", candidate)
		}

		if new_path == home {
			if VERBOSE {
				fmt.Println("new_path == home, returning original")
			}
			return filename
		} else {
			if VERBOSE {
				fmt.Println("new_path != home")
			}
			// take a directory off the path and put the fileame on
			splits := strings.Split(new_path, "/")
			new_path = ""
			for index := 0; index < len(splits)-1; index++ {
				if splits[index] == "" {
					continue
				}
				new_path += "/"
				new_path += splits[index]
			}
			if VERBOSE {
				fmt.Printf("new_path  : %v\n", new_path)
			}
			candidate = new_path + "/" + filename
			if VERBOSE {
				fmt.Printf("candidate : %v\n", candidate)
			}
			if goutils.FileExists(candidate) {
				return candidate
			}

		}
	}
}

func GetFilename(c *cli.CLI) (string, error) {
	VERBOSE := c.Contains("-v")
	value := c.GetStringOrDefault("-f", "")
	if value == "" {
		value = c.GetStringOrDefault("-file", "")
	}
	if value == "" {
		value = os.Getenv(ENV_BUILDFILE)
		if c.IS_VERBOSE {
			fmt.Printf("ENV value %v is %v\n", ENV_BUILDFILE, value)
		}
	}
	if value == "" {
		value = DEFAULT_BUILDFILE
	}
	filename := Findfile(value, VERBOSE)
	info, _ := os.Stat(filename)
	if info.IsDir() {
		return "", errors.New(value + " is a directory.")
	} else {
		return filename, nil
	}

}

func Load(c *cli.CLI) *Version {
	filename, err := GetFilename(c)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
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
