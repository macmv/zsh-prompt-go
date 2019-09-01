package main

import (
  . "github.com/macmv/zsh-prompt/lib"
  "strconv"
  "os"
  "fmt"
  "strings"
  "os/user"
)

func utf8_length(str string) int {
  length := 0
  for i := 0; i < len(str); i++ {
    let := str[i]
    if let >> 6 == 2 {
      continue
    }
    // skip until hit 'm'
    if let == '\033' {
      var j int
      for j = 0; j < 1000; j++ {
        let = str[i + j]
        if j + i > len(str) {
          j--
          break
        }
        if let == 'm' {
          break
        }
      }
      i += j
    } else {
      length += 1
    }
  }
  return length
}

func get_cwd() string {
  cwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  usr, err := user.Current()
  if err != nil {
    panic(err)
  }
  if strings.HasPrefix(cwd, usr.HomeDir) {
    cwd = cwd[len(usr.HomeDir):]
    cwd = "~" + cwd
  }
  return cwd
}

func get_user() string {
  usr, err := user.Current()
  if err != nil {
    panic(err)
  }
  return usr.Username
}

func get_host() string {
  name, err := os.Hostname()
  if err != nil {
    panic(err)
  }
  return name
}

func git_branch() string {
  cwd, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  git.Branch
}

func main() {
  left_sections := []Section {
    Section {Text: " " + get_host(), Fg: "fff", Bg: "06a"},
    Section {Text: get_user(), Fg: "fff", Bg: "09d"},
    Section {Text: get_cwd(), Fg: "fff", Bg: "555"},
    Section {Text: git_branch(), Fg: "fff", Bg: "2a2"},
  }
  right_sections := []Section {
    Section {Text: "time", Fg: "000", Bg: "ccc"}, // time
    Section {Text: "date with letters", Fg: "fff", Bg: "888"}, // date but letters
    Section {Text: "date with numbers", Fg: "fff", Bg: "555"}, // date but numbers
  }
  if (len(os.Args) != 2) {
    panic("Must enter width as first arg")
  }
  width, err := strconv.Atoi(os.Args[1])
  if (err != nil) {
    panic(err)
  }
  left_string := GenerateSections("\uE0B0", left_sections, false)
  right_string := GenerateSections("\uE0B2", right_sections, true)
  fmt.Print(Paint("┏━", "06a", "000"))
  fmt.Print(left_string)
  left_length := 2 + utf8_length(left_string)
  right_length := utf8_length(right_string)
  fmt.Print(strings.Repeat(" ", width - left_length - right_length))
  fmt.Print(right_string)
  fmt.Print("\n")
  fmt.Print(Paint("┗━ ", "06a", "000"))
}
