package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ParseFlags parses the command line args, allowing flags to be specified after positional args.
func ParseFlags() error {
	return ParseFlagSet(flag.CommandLine, os.Args[1:])
}

// ParseFlagSet works like flagset.Parse(), except positional arguments are not required to come after flag arguments.
func ParseFlagSet(flagset *flag.FlagSet, args []string) error {
	var positionalArgs []string
	for {
		if err := flagset.Parse(args); err != nil {
			return err
		}
		args = args[len(args)-flagset.NArg():]
		if len(args) == 0 {
			break
		}
		positionalArgs = append(positionalArgs, args[0])
		args = args[1:]
	}
	return flagset.Parse(positionalArgs)
}

// Check if file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func main() {

	unixTimestampFlag := flag.String("u", "", "Unix timestamp, seconds since Jan 01 1970")
	noCreateFlag := flag.Bool("c", false, "Do no create any files")
	accessTimeOnlyFlag := flag.Bool("a", false, "Change only the access time")
	modifyTimeOnlyFlag := flag.Bool("m", false, "Change only the modification time")
	timestampFlag := flag.String("t", "", "Use time format YYYYMMDDHHMMSS")

	flag.Usage = func() {
		fmt.Println("Usage: touch FILE [OPTION]... ")
		fmt.Println("Update the access and modification times of each FILE to the current time.")
		fmt.Println("")
		fmt.Println("A FILE argument that does not exist is created empty, unless -c or -h is supplied.")
		fmt.Println("If you supply both -u and -t only -u will be used")
		fmt.Println("")
		fmt.Println("  -u STAMP        Unix timestamp, seconds since Jan 01 1970")
		fmt.Println("  -t TIME         Use time format YYYYMMDDHHMMSS")
		fmt.Println("  -c              Do no create any files")
		fmt.Println("  -a              Change only the access time")
		fmt.Println("  -m              Change only the modification time")
		fmt.Println("")
	}

	ParseFlags()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	// Set newTimeValue to current time, unless overridden
	newTimeValue := time.Now().Local()

	if timestampFlag != nil && *timestampFlag != "" {
		ts, err := time.Parse("20060102150405", *timestampFlag)
		if err != nil {
			fmt.Println("Cannot parse time, use YYYYMMDDHHMMSS e.g. 20060102150405")
			os.Exit(2)
		}
		newTimeValue = ts
	}

	if unixTimestampFlag != nil && *unixTimestampFlag != "" {

		i, err := strconv.ParseInt(*unixTimestampFlag, 10, 64)
		if err != nil {
			fmt.Println("Cannot parse unix time, use integer value e.g. 1672524000 for Jan 01 2023 00:00:00")
			os.Exit(2)
		}
		newTimeValue = time.Unix(i, 0)
	}

	accessTime := time.Time{}
	modifyTime := time.Time{}

	if *accessTimeOnlyFlag {
		accessTime = newTimeValue
	}

	if *modifyTimeOnlyFlag {
		modifyTime = newTimeValue
	}

	if !*accessTimeOnlyFlag && !*modifyTimeOnlyFlag {
		accessTime = newTimeValue
		modifyTime = newTimeValue
	}

	// Start processing FILEs
	for n, filename := range flag.Args() {

		_ = n

		// Create file if not exists and -c is not supplied
		if !fileExists(filename) && !*noCreateFlag {
			f, err := os.Create(filename)
			if err != nil {
				fmt.Println("Cannot create file\n", err)
				os.Exit(2)
			}
			f.Close()
		}

		//Set both access time and modified time of the file to the current time
		if fileExists(filename) {
			err := os.Chtimes(filename, accessTime, modifyTime)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
