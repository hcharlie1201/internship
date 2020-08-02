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
        Name:        "add",
        Aliases:     []string{"a"},
        Usage:       "options for task templates",
        Action: func(c *cli.Context) error {
              fmt.Println("new task template: ", c.Args().First())
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
