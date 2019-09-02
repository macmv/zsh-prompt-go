package main

import (
  . "github.com/macmv/zsh-prompt/lib"
  "fmt"
  "os/exec"
  "strings"
  "os"
)

func main() {
  colors := []string {
    "06a",
    "f43",
    "09d",
    "96f",
    "f72",
    "7c0",
    "8f2",
    "000",
  }
  args := []string {"-lah"}
  if (len(os.Args) > 1) {
    args = append(args, os.Args[1:]...)
  }
  out_bytes, _ := exec.Command("ls", args...).Output()
  out := string(out_bytes)
  lines := strings.Split(out, "\n")[1:]
  for i := 0; i < len(lines) - 1; i++ {
    var sections []Section
    cols := strings.Split(lines[i], " ")
    string_sections := []string {"", "", "", "", "", "", "", "", ""}
    for j := 0; j < len(cols); j++ {
      if j < 8 {
        if cols[j] == "" {
          cols = append(cols[:j], cols[j+1:]...)
          string_sections[j] = " " + string_sections[j]
          j--
        } else {
          string_sections[j] += cols[j]
        }
      } else {
        string_sections[8] += strings.Join(cols[j:], " ")
        break
      }
    }
    color_index := 0
    for j := 0; j < 9; j++ {
      if (j == 5) {
        sections = append(sections, Section {Text: string_sections[j] + " " + string_sections[j + 1], Fg: "fff", Bg: colors[color_index]})
        j++
      } else {
        sections = append(sections, Section {Text: string_sections[j], Fg: "fff", Bg: colors[color_index]})
      }
      color_index++
    }
    fmt.Println(GenerateSections("\uE0B0", sections, false))
  }
}
