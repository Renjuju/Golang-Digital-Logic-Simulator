package main

import ("fmt"
        "bufio"
        "os"
      )

func main() {
  circuit, err:= readLines("input.txt")
  check(err)
  for i, circuit:= range circuit {
    fmt.Println(i, circuit)
  }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//http://stackoverflow.com/questions/5884154/golang-read-text-file-into-string-array-and-write
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  check(err)
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}
