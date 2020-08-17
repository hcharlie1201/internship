package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:                   "internship",
		Usage:                  "tracks internship",
		UseShortOptionHandling: true,
		Action: func(c *cli.Context) error {
			fmt.Println("Hello users of the cli!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "initalize the file",
				Action: func(c *cli.Context) error {
					ifexists := checkFile("internship.txt")
					if ifexists == true {
						fmt.Println("File Already Exists, use the delete function if you want to start a new file")
						return nil
					}
					filePath, _ := filepath.Abs("internship.txt")
					f, err := os.Create(filePath)
					if err != nil {
						fmt.Println(err)
						return nil
					}
					f.Close()
					fmt.Println("Successfully Initialized")
					return nil
				},
			},
			{
				Name:    "approved",
				Aliases: []string{"app"},
				Usage:   "when an internship gets back to you, and approved you",
				Action: func(c *cli.Context) error {
					nameCompany := c.Args().First()
					if !checkFile("added.txt") {
						os.Create("added.txt")
					}
					if !checkFile("internship.txt") {
						fmt.Println("Don't have the internship.txt file yet!")
					}
					e := setData(nameCompany, "added.txt")
					if e != nil {
						fmt.Println(e)
						return nil
					}
					return nil
				},
			},
			{
				Name:    "rejected",
				Aliases: []string{"rej"},
				Usage:   "when an internship gets back to you, and rejected you",
				Action: func(c *cli.Context) error {
					nameCompany := c.Args().First()
					e := setData(nameCompany, "rejected.txt")
					if e != nil {
						fmt.Println(e)
						return nil
					}
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "adding the internship with the date listed",
				Action: func(c *cli.Context) error {
					//check if file exists
					arg2 := c.Args().First()
					currentTime := time.Now()
					if arg2 == "" {
						fmt.Println("Must specify a company you want to add")
						return nil
					}
					if checkFile("internship.txt") {
						f, _ := os.OpenFile("internship.txt", os.O_RDWR|os.O_APPEND, 0660)
						defer f.Close()
						arg2 = arg2 + " " + currentTime.String() + "\n"
						if _, errorz := f.WriteString(arg2); errorz != nil {
							panic(errorz)
						}
						fmt.Println("Successfully added " + arg2 + "\n")
						f.Close()
						return nil
					}
					f := returnFile("internship.txt")
					arg2 = arg2 + " " + currentTime.String()
					f.WriteString(arg2 + "\n")
					fmt.Println("Successfully added " + arg2 + "!")
					f.Close()
					return nil
				},
			},
			{
				Name:    "gui",
				Usage:   "opens up a gui to view your files added",
				Aliases: []string{"g"},
				Action: func(c *cli.Context) error {
					GUI()
					fmt.Println("GUI opened up")
					return nil
				},
			},
			{
				Name:    "list",
				Usage:   "list current internship applied, approved, and rejected",
				Aliases: []string{"l"},
				Action: func(c *cli.Context) error {
					f, _ := os.Open("internship.txt")
					scanner := bufio.NewScanner(f)
					if err := scanner.Err(); err != nil {
						fmt.Println(err)
						return nil
					}
					fmt.Println("Current applied internships:")
					for scanner.Scan() {
						fmt.Println(scanner.Text())
					}
					if checkFile("added.txt") == true {
						f2, _ := os.Open("added.txt")
						fmt.Println("Current approved internships:")
						scanner2 := bufio.NewScanner(f2)
						for scanner2.Scan() {
							fmt.Println(scanner2.Text())
						}
						f2.Close()
					}
					if checkFile("rejected.txt") == true {
						f2, _ := os.Open("rejected.txt")
						fmt.Println("Current rejected internships:")
						scanner2 := bufio.NewScanner(f2)
						for scanner2.Scan() {
							fmt.Println(scanner2.Text())
						}
						f2.Close()
					}
					f.Close()
					return nil
				},
			},
			{
				Name:    "remove",
				Usage:   "remove a specified subcommand ",
				Aliases: []string{"rm"},
				Subcommands: []*cli.Command{
					{
						Name:    "all",
						Usage:   "Delete all files",
						Aliases: []string{"a"},
						Action: func(c *cli.Context) error {
							a := []string{"internship.txt", "added.txt", "rejected.txt"}
							for _, val := range a {
								if checkFile(val) == true {
									e := os.Remove(val)
									if e != nil {
										log.Fatal(e)
										return nil
									}
								}
							}
							fmt.Println("Successfully deleted everything if it exists")
							return nil
						},
					},
					{
						Name:    "company",
						Usage:   "delete a speficied company that you added",
						Aliases: []string{"c"},
						Action: func(c *cli.Context) error {
							comp := c.Args().First()
							if comp == "" {
								fmt.Println("Need to specify a company you want to delete.")
								return nil
							}
							if !checkFile("internship.txt") {
								fmt.Println("Need to add an internship.txt file first!")
							}
							e := deleteData(comp, "internship.txt")
							if e != nil {
								fmt.Println(e)
								return nil
							}
							fmt.Println("Deleted " + comp + "sucessfully if it exists.")
							return nil
						},
					},
					{
						Name:    "file",
						Usage:   "delete a file",
						Aliases: []string{"p"},
						Action: func(c *cli.Context) error {
							fname := c.Args().First()
							if fname == "" {
								fmt.Println("Need to specify a company you want to delete.")
								return nil
							}
							os.Remove(fname)
							fmt.Println("File deleted if it exists.")
							return nil
						},
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func checkFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func returnFile(path string) *os.File {
	filePath, _ := filepath.Abs(path)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return f
}

func setData(nameCompany string, path string) error {
	f, _ := os.Open("internship.txt")
	scanner := bufio.NewScanner(f)
	os.Create("temp.txt")
	f2, _ := os.OpenFile("temp.txt", os.O_RDWR|os.O_APPEND, 0660)
	addedFile, _ := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	defer f.Close()
	defer f2.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, nameCompany) {
			if _, errorz := f2.WriteString(line + "\n"); errorz != nil {
				panic(errorz)
			}
		} else {
			fmt.Println("Added internship to the " + path + " filepath")
			if _, errorz := addedFile.WriteString(line + "\n"); errorz != nil {
				panic(errorz)
			}
		}
	}
	f.Close()
	e := os.Remove("internship.txt")
	if e != nil {
		fmt.Println(e)
		return e
	}
	os.Rename("temp.txt", "internship.txt")
	addedFile.Close()
	f2.Close()
	return nil
}

func deleteData(nameCompany string, path string) error {
	f, _ := os.Open("internship.txt")
	scanner := bufio.NewScanner(f)
	os.Create("temp.txt")
	f2, _ := os.OpenFile("temp.txt", os.O_RDWR|os.O_APPEND, 0660)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, nameCompany) {
			f2.WriteString(line + "\n")
		} else {
			fmt.Println("Removed " + nameCompany + "successfully")
		}
	}
	f.Close()
	e := os.Remove("internship.txt")
	if e != nil {
		fmt.Println(e)
		return e
	}
	os.Rename("temp.txt", "internship.txt")
	f2.Close()
	return nil
}
