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
    "000",
  }
  args := []string {"-lah", "--color"}
  if (len(os.Args) > 1) {
    args = append(args, os.Args[1:]...)
  }
  out_bytes, err := exec.Command("ls", args...).Output()
  if (err != nil) {
    panic(err)
  }
  out := string(out_bytes)
  lines := strings.Split(out, "\n")[1:]
  for i := 0; i < len(lines) - 1; i++ {
    // cols: ["rwrwrwr", "", "", "", "2", "macmv", "macmv", "4.0K", "Sep", "1", "2019", "file.txt"]
    cols := strings.Split(lines[i], " ")
    // string_sections: ["rwrwrwr", "   2", "macmv", "macmv", "4.0K", "Sep", "1", "2019", "file.txt"]
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
    // sections: the Section objects used to generate the output
    var sections []Section
    // color_index: index of the colors array defined above
    color_index := 0
    for j := 0; j < 8; j++ {
      if (j == 5) { // this adds the date string together, so that ["Sep", "1", "2019"] -> ["Sep 1 2019"]
        sections = append(sections, Section {Text: string_sections[j] + " " + string_sections[j + 1] + " "  + string_sections[j + 2], Fg: "fff", Bg: colors[color_index]})
        j += 2
      } else {
        sections = append(sections, Section {Text: string_sections[j], Fg: "fff", Bg: colors[color_index]})
      }
      color_index++
    }
    // \uE0B0 is the powerline font seperator
    fmt.Print(GenerateSections("\uE0B0", sections, false))
    fmt.Println(string_sections[len(string_sections) - 1])
  }
}
