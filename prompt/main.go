package main

import (
  . "github.com/macmv/zsh-prompt/lib"
  "strconv"
  "os"
  "fmt"
  "strings"
  "os/user"
  "os/exec"
  "time"
  "github.com/lestrrat/go-strftime"
)

/*
TODO:
- make get_cwd have powerline spacers instead of /
- fix bugs with tab completion
*/

func get_cwd(separator string) string {
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
  out_arr := strings.Split(cwd, "/")
  out_arr[len(out_arr) - 1] = "\033[1m" + out_arr[len(out_arr) - 1]
  return strings.Join(out_arr, separator)
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
func get_git() string {
  out_bytes, _ := exec.Command("git", "branch").Output()
  out := string(out_bytes)
  var branch string
  if len(out) == 0 {
    return ""
  } else {
    branch = strings.Split(out, "\n")[0][2:]
  }
  out_bytes, _ = exec.Command("git", "diff").Output()
  out = string(out_bytes)
  var status string
  if out == "" {
    status = ""
  } else {
    status = "* "
  }
  return status + branch
}
func get_time() string {
  result, err := strftime.Format("%I:%M:%S %p", time.Now())
  if (err != nil) {
    panic(err)
  }
  return result
}
func get_date_letters() string {
  result, err := strftime.Format("%a, %b %d, %Y", time.Now())
  if (err != nil) {
    panic(err)
  }
  return result
}
func get_date_numbers() string {
  result, err := strftime.Format("%Y-%m-%d", time.Now())
  if (err != nil) {
    panic(err)
  }
  return result
}

func main() {
  left_sections := []Section {
    Section {Text: " " + get_host(), Fg: "fff", Bg: "06a"},
    Section {Text: get_user(), Fg: "fff", Bg: "09d"},
    Section {Text: get_cwd(" \uE0B1 "), Fg: "fff", Bg: "555"},
    Section {Text: get_git(), Fg: "000", Bg: "ff0"},
  }
  right_sections := []Section {
    Section {Text: get_time(), Fg: "000", Bg: "ccc"}, // time
    Section {Text: get_date_letters(), Fg: "fff", Bg: "888"}, // date but letters
    Section {Text: get_date_numbers(), Fg: "fff", Bg: "555"}, // date but numbers
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
  left_length := 2 + UTF8Length(left_string)
  right_length := UTF8Length(right_string)
  spacer_size := width - left_length - right_length
  if (spacer_size > 0) {
    fmt.Print(strings.Repeat(" ", spacer_size))
    fmt.Print(right_string)
  }
  fmt.Print("\n")
  fmt.Print(Paint("┗━ ", "06a", "000"))
}
