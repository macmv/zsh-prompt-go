package lib

import (
  "strconv"
  "fmt"
)

type Section struct {
  Text string
  Fg string
  Bg string
}

func getRGB(code string) string {
  if len(code) != 3 {
    panic("String must be 3 chars!")
  }
  color, err := strconv.ParseInt(code, 16, 32)
  if err != nil {
    panic(err)
  }
  r := (color / 256) * 17
  g := (color / 16 % 16) * 17
  b := (color % 16) * 17
  return fmt.Sprintf("%d;%d;%d", r, g, b)
}

func Paint(text string, fg string, bg string) string {
  var str string
  str += "\033[38;2;"
  str += getRGB(fg)
  str += ";48;2;"
  str += getRGB(bg)
  str += "m"
  str += text
  str += "\033[38;2;255;255;255;48;2;0;0;0m"
  return str
}

/*
separator: the triangle to seperate sections
sections: array of strings and colors to join
length: length of sections
direction:
  true: is the prompt on the right
  false: is the prompt on the left
  this changes how the colors should change with the separators
*/
func GenerateSections(separator string, sections []Section, direction bool) string {
  var str string
  for i := 0; i < len(sections); i++ {
    if sections[i].Text == "" {
      sections = append(sections[:i], sections[i + 1:]...) // doesn't check [i + 1].Text == ""
    }
  }
  for i := 0; i < len(sections); i++ {
    var sec1, sec2 Section;
    if direction { // right side (left facing arrows)
      if i <= 0 {
        sec1.Text = separator
        sec1.Fg = sections[i].Bg
        sec1.Bg = "000"
      } else {
        sec1.Text = separator
        sec1.Fg = sections[i].Bg
        sec1.Bg = sections[i - 1].Bg
      }
      sec2 = sections[i]
      str += Paint(" ", sec1.Fg, sec1.Bg)
      str += Paint(sec1.Text, sec1.Fg, sec1.Bg)
      str += Paint(" ", sec2.Fg, sec2.Bg)
      str += Paint(sec2.Text, sec2.Fg, sec2.Bg)
    } else { // left side (right facing arrows)
      sec1 = sections[i]
      if i >= len(sections) - 1 {
        sec2.Text = separator
        sec2.Fg = sections[i].Bg
        sec2.Bg = "000"
      } else {
        sec2.Text = separator
        sec2.Fg = sections[i].Bg
        sec2.Bg = sections[i + 1].Bg
      }
      str += Paint(sec1.Text, sec1.Fg, sec1.Bg)
      str += Paint(" \033[0m", sec1.Fg, sec1.Bg)
      str += Paint(sec2.Text, sec2.Fg, sec2.Bg)
      str += Paint(" \033[0m", sec2.Fg, sec2.Bg)
    }
  }
  return str
}
