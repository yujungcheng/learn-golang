package main

import (
  "flag"
  "fmt"
  "os"
  "strings"
)

func flagUsage() {
  usageText := `subcommand tool example

  Usage:
    subcommands command [arguments]
  The commands are:
    uppercase - uppercase a string
    lowercase - lowercase a string
  Example:
    subcommands uppercase -s "my lowercase string"
  Use "subcommands --help" for more info.`
  fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func main() {
  flag.Usage = flagUsage // set usage func
  uppercaseCmd := flag.NewFlagSet("uppercase", flag.ExitOnError)
  lowercaseCmd := flag.NewFlagSet("lowercase", flag.ExitOnError)

  if len(os.Args) == 1 {
    flag.Usage()
    return
  }

  switch os.Args[1] {
  case "uppercase":
    s := uppercaseCmd.String("s", "", "A string of text to be uppercased")
    uppercaseCmd.Parse(os.Args[2:])
    fmt.Println(strings.ToUpper(*s))
  case "lowercase":
    s := lowercaseCmd.String("s", "", "A string of text to be lowercased")
    lowercaseCmd.Parse(os.Args[2:])
    fmt.Println(strings.ToLower(*s))
  default:
    flag.Usage()
  }
}
