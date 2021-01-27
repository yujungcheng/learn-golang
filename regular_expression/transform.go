package main

import (
  "fmt"
  "regexp"
)

func main() {
  usernames := [4]string {
    "slimshady99",
    "!asdf£33£3",
    "roger",
    "Iamthebestuserofthisappevaaaar",
  }

  re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")
  an := regexp.MustCompile(":^alnum:")

  for _, username := range usernames {
    fmt.Printf("checking username %v\n", username)
    // truncated to 12 charactersr
    if len(username) > 12 {
      oldUsername := username
      username = username[:12]
      fmt.Printf("- trimmed username %v to %v\n", oldUsername, username)
    }
    if !re.MatchString(username) {
       // invalid characters are replaced with 'x'
       oldUsername := username
      username = an.ReplaceAllString(username, "x")
      fmt.Printf("- rewrote username %v to %v\n", oldUsername, username)
    }
  }
}
