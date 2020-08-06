package main

import (
  "fmt"
  "strings"
  "bufio"
  "time"
  "log"
  "os"
  "github.com/urfave/cli/v2"
  "path/filepath"
)

func main() {
  app := &cli.App{
    Name: "internship",
    Usage: "tracks internship",
    Action: func(c *cli.Context) error {
      fmt.Println("Hello users of the cli!")
      return nil
    },
    Commands: []*cli.Command{
      {
        Name:    "init",
        Aliases: []string{"i"},
        Usage:   "initalize the file",
        Action:  func(c *cli.Context) error {
            ifexists := checkFile("internship.txt")
            if ifexists == true {
                fmt.Println("File Already Exists")
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
        Action:  func(c *cli.Context) error {
            nameCompany := c.Args().First()
            e := SetData(nameCompany, "added.txt")
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
        Action:  func(c *cli.Context) error {
            nameCompany := c.Args().First()
            e := SetData(nameCompany, "rejected.txt")
            if e != nil {
                fmt.Println(e)
                return nil
            }
            return nil
        },
      },
      {
        Name:        "add",
        Aliases:     []string{"a"},
        Usage:       "adding the internship with the date listed",
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
            fmt.Println("Successfully added" + arg2)
            f.Close()
            return nil
          },
      },
      {
          Name: "list",
          Usage: "List current internship applied, approved, and rejected",
          Action: func(c *cli.Context) error {
            f := returnFile("internship.txt")
            scanner := bufio.NewScanner(f)
            if err := scanner.Err(); err != nil {
                fmt.Println(err)
                return nil
            }
            f.Close()
            fmt.Println("Current applied internships")
            for scanner.Scan() {
                fmt.Println(scanner.Text())
            }
            if checkFile("added.txt") == true {
                f2 := returnFile("added.txt")
                fmt.Println("Current approved internships")
                scanner2 := bufio.NewScanner(f2)
                for scanner2.Scan() {
                    fmt.Println(scanner2.Text())
                }
                f2.Close()
            }
            if checkFile("rejected.txt") == true {
                f2 := returnFile("rejected.txt")
                fmt.Println("Current rejected internships")
                scanner2 := bufio.NewScanner(f2)
                for scanner2.Scan() {
                    fmt.Println(scanner2.Text())
                }
                f2.Close()
            }
            return nil
          },
      },
      {
         Name:  "remove",
         Usage: "remove an existing template",
         Action: func(c *cli.Context) error {
              fmt.Println("removed task template: ", c.Args().First())
              return nil
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
    if _, err := os.Stat(path); os.IsNotExist(err){
       return false
    } else {
        return true
    }
}

func returnFile(path string)*os.File {
    filePath, _ := filepath.Abs("internship.txt")
    f, err := os.Create(filePath)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return f
}

func SetData(nameCompany string, path string) error {
    f := returnFile("internship.txt")
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
            f2.WriteString(line)
        } else {
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
