package main

import ("fmt"
        "bufio"
        "os"
        "sync"
      )

func printer(c1 chan string, c2 chan string) {
  fileName:= <-c1
  circuit, err:= readLines(fileName)
  check(err)
  for i, circuit:= range circuit {
    fmt.Println(i, circuit)
  }
}

func reader(c1 chan string) {
  fmt.Println("What's the name of your file?")
  var fileName string
  fmt.Scanln(&fileName)
  go func() {
    c1 <- fileName
  }()
}

func main() {
  c1, c2, done := make(chan string), make(chan string), make(chan bool)
  go reader(c1)
  go printer(c1, c2)
}

//logic gates
func not(x bool) bool{
  return !x
}

func and(x, y bool) bool {
  return x && y
}

func or(x, y bool) bool {
  if(x || y) {
    return true
  } else {
    return false
  }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// reads line by line
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
