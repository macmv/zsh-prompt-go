package lib

func UTF8Length(str string) int {
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
