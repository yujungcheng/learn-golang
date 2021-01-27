package main

import (
  "fmt"
  "log"
  "regexp"
)

func main() {
  needle := "chocolate"
  haystack := "Chocolate is my favorite!"
  match, err := regexp.MatchString(needle, haystack)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(match)

  needle = "(?i)chocolate"  // indicate case insensitive search
  match, err = regexp.MatchString(needle, haystack)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(match)
}

/*
. Matches any character except a line break
* Matches the previous character zero or more
times
^ Signifies the start of a line
$ Signifies the end of a line
+ Matches one or more times
? Matches zero or one times[] Matches any character with the brackets
{n} Matches n times
{n,} Matches n or more times
{m,n} Matches at least m times and at most n times
*/
