package main

import ("fmt"
        "bufio"
        "os"
        "time"
        "sync"
        "strings"
      )

var wg sync.WaitGroup

func readGates(c4 chan []string) {
  circuits := <-c4
  println(circuits[0])
  statements := strings.Split(circuits[0], "")
  elements := []int{}
  var tempVal int

  // asks for intial values of inputs
  for i, element:= range statements {
    println(i, element)
    println("What would you like the value of ", element, " to be ")
    fmt.Scanln(&tempVal)
    elements = append(elements, tempVal)
  }

  // uses gates to parse inputs
  line := strings.Split(circuits[1], "+")
  fmt.Println(line)
  for i, element:=range line {
    println(i, element)
  }

  wg.Done()
}

func clock(c3 chan string) {
  var x int
  x = 0
  println("How much clock pulses per second?")
  var pulsesPerSecond int
  // fmt.Scanln(&pulsesPerSecond)
  pulsesPerSecond = 3
  c3<-"done"
  for {
    x++
    // println(x)
    time.Sleep(time.Second / time.Duration(pulsesPerSecond))
    x--
  }
  wg.Done()
}

// prints and sends array of string to readGates
func printer(c1 chan string, c2 chan string, c4 chan []string) {
  fileName:= <-c1
  circuit, err:= readLines(fileName)
  check(err)
  // for i, circuit:= range circuit {
  //   fmt.Println(i, circuit)
  // }
  c4<-circuit
  wg.Done()
}

func reader(c1 chan string, c3 chan string) {
  <-c3
  fmt.Println("What's the name of your file?")
  var fileName string
  // fmt.Scanln(&fileName)
  fileName = "input.txt"
  go func() {
    c1 <- fileName
  }()
  wg.Done()
}

func main() {
  c1, c2, c3, c4 := make(chan string), make(chan string), make(chan string), make(chan []string)
  wg.Add(4)
  go clock(c3)
  go reader(c1, c3)
  go printer(c1, c2, c4)
  go readGates(c4)
  wg.Wait()
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

func nand(x,y bool) bool {
    if (!and(x, y)) {
      return true
    }
    return false
}

func nor(x,y bool) bool {
  return true
}

func xor() bool {
  //return or(and(x, not(y), and(y, not(x))))
  return true
}

func check(e error) {
    if e != nil {
        // panic(e)
    }
}

// reads line by line
// boiler plate code from
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
