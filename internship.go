package main

import (
  "fmt"
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
        Aliases: []string{"approved"},
        Usage:   "complete a task on the list",
        Action:  func(c *cli.Context) error {
          fmt.Println("completed task: ", c.Args().First())
          return nil
        },
      },
      {
        Name:        "add",
        Aliases:     []string{"a"},
        Usage:       "adding the internship with the date listed",
        Action: func(c *cli.Context) error {
            //check if file exists
            ifexists := checkFile("internship.txt")
            if ifexists == false {
                fmt.Println("File Already Exists")
            }
            arg2 := c.Args().First()
            if arg2 == "" {
                fmt.Println("Must specify a company you want to add")
                return nil
            }
            f := returnFile()
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
            ifexists := checkFile("internship.txt")
            if ifexists == false {
                fmt.Println("File Already Exists")
                return nil
            }
            f := returnFile()
            scanner := bufio.NewScanner(f)
            if err := scanner.Err(); err != nil {
                fmt.Println(err)
                return nil
            }
            fmt.Println("Current applied internships")
            for scanner.Scan() {
                fmt.Println(scanner.Text())
            }
            f.Close()
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
    if _, err := os.Stat("internship.txt"); os.IsNotExist(err){
       return false
    } else {
        return true
    }
}

func returnFile()*os.File {
    filePath, _ := filepath.Abs("internship.txt")
    f, err := os.Create(filePath)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return f
}
