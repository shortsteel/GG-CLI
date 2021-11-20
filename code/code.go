package code

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Code() *cli.Command {
	command := cli.Command{
		Name:    "code",
		Usage:   "The GG CLI",
		Aliases: []string{"co", "c"},
		Subcommands: []*cli.Command{
			cmdList(),
		},
	}

	return &command
}

func cmdList() *cli.Command {
	command := cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "list code content",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "suffixes",
				Usage:   "All path and file matching the suffixes will be skipped",
				Aliases: []string{"suf"},
			},
			&cli.StringSliceFlag{
				Name:    "prefixes",
				Usage:   "All path and file matching the prefixes will be skipped",
				Aliases: []string{"pre"},
			},
		},
		Action: func(context *cli.Context) error {
			excludedSuffixes := context.StringSlice("suffixes")
			excludedPrefixes := context.StringSlice("prefixes")

			root := context.Args().First()
			if root == "" {
				root = "."
			}

			isExcluded := func(name string) bool {
				for _, value := range excludedSuffixes {
					if strings.HasSuffix(name, value) {
						return true
					}
				}
				for _, value := range excludedPrefixes {
					if strings.HasSuffix(name, value) {
						return true
					}
				}
				return false
			}

			var walkDirFunc func(dir string) error
			walkDirFunc = func(dir string) error {
				files, _ := ioutil.ReadDir(dir)
				for _, file := range files {
					currentFilePath := filepath.Join(dir, file.Name())
					if !isExcluded(file.Name()) {
						if file.IsDir() {
							walkDirFunc(currentFilePath)
						} else {
							fmt.Println()
							fmt.Println()
							fmt.Println("**********************")
							fmt.Println("** " + currentFilePath)
							fmt.Println("**********************")
							currentFile, _ := os.Open(currentFilePath)
							fmt.Print(ioutil.ReadAll(currentFile))
						}
					}
				}
				return nil
			}

			err := walkDirFunc(root)
			if err != nil {
				return err
			}

			return nil
		},
	}
	return &command
}
