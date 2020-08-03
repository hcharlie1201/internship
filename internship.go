package main

import (
  "fmt"
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
        Name:    "complete",
        Aliases: []string{"c"},
        Usage:   "complete a task on the list",
        Action:  func(c *cli.Context) error {
          fmt.Println("completed task: ", c.Args().First())
          return nil
        },
      },
      {
        Name:        "add [company_name]",
        Aliases:     []string{"a"},
        Usage:       "options for task templates",
        Action: func(c *cli.Context) error {
            //check if file exists
            ifexists := checkFile("internship.txt")
            if ifexists == false {
                fmt.Println("File Already Exists")
            }
            arg2 := c.Args().Get(1)
            if arg2 == "" {
                fmt.Println("Must specify a company you want to add")
                return nil
            }
            f := returnFile()
            f.WriteString(arg2 + "\n")
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
