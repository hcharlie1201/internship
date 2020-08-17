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
						fmt.Println("File Already Exists")
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
					if arg2 == "" {
						fmt.Println("Must specify a company you want to add")
						return nil
					}
					f := returnFile("internship.txt")
					currentTime := time.Now()
					arg2 = arg2 + " " + currentTime.String()
					f.WriteString(arg2 + "\n")
					fmt.Println("Successfully added " + arg2)
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
					f, _ := os.Open("./internship.txt")
					scanner := bufio.NewScanner(f)
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
						f2, _:= os.Open("rejected.txt")
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
    if checkFile(path) {
        f, _ := os.Open(path)
        return f
    } else {
        f, _ := os.Create(path)
        return f
    }
}

func setData(nameCompany string, path string) error {
	f := returnFile(path)
	scanner := bufio.NewScanner(f)
	f2 := returnFile("temp.txt")
	addedFile := returnFile(path)
	for scanner.Scan() {
		line := scanner.Text()
        fmt.Println(line)
        if strings.Contains(line, nameCompany) {
			f2.WriteString(line)
            fmt.Println("Added internship to the " + path + " filepath")
            addedFile.WriteString(line + "\n")
		}
	}
	e := os.Remove("internship.txt")
	if e != nil {
		fmt.Println(e)
		return e
	}
	os.Rename("temp.txt", "internship.txt")
	addedFile.Close()
	f2.Close()
	f.Close()
	return nil
}

func deleteData(nameCompany string, path string) error {
	f := returnFile(path)
	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	f2 := returnFile("temp.txt")
	addedFile := returnFile(path)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, nameCompany) {
			addedFile.WriteString(line + "\n")
		} else {
			fmt.Println("Removed " + nameCompany + "successfully")
		}
	}
	e := os.Remove(path)
	if e != nil {
		fmt.Println(e)
		return e
	}
	os.Rename("temp.txt", path)
	addedFile.Close()
	f2.Close()
	f.Close()
	return nil
}
