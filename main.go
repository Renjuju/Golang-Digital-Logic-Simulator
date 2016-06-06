package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var x int

func main() {
	c1, c2, c3, c4 := make(chan string), make(chan string), make(chan string), make(chan []string)

	wg.Add(4)
	go clock(c3)
	go reader(c1, c3)
	go printer(c1, c2, c4)
	go readGates(c4)
	wg.Wait()
}

func convertToBoolean(x int) bool {
	if x == 1 {
		return true
	}
	return false
}

func readGates(c4 chan []string) {
	circuits := <-c4
	println("\nReading gates...")

	var gateChannels [3]chan bool
	for i := range gateChannels {
		gateChannels[i] = make(chan bool)
	}

	statements := strings.Split(circuits[0], "")
	elements := []int{}
	var tempVal int

	// asks for intial values of inputs
	for i, element := range statements {
		println("[", i, "]", "What would you like the value of ", element, " to be ")
		fmt.Scanln(&tempVal)
		elements = append(elements, tempVal)
	}

	fmt.Println("The number of elements: ", len(elements))

	// uses gates to parse inputs
	for n := 1; x < len(circuits); n++ {
		line := strings.Split(circuits[n], " ")
		for i, element := range line {
			println(i, element)
		}
		// var retVal bool
		for i := range elements {
			go func(i int) {
				gateChannels[i] <- convertToBoolean(elements[i])
			}(i)
		}

		// todo: better parsing for varying letters
		switch line[0] {
		case "AND":
			println(and(gateChannels[0], gateChannels[1]))
		case "OR":
			println(or(gateChannels[0], gateChannels[1]))
		case "NOR":
			println(nor(gateChannels[0], gateChannels[1]))
		case "NOT":
			println(not(gateChannels[0]))
		case "XOR":
			println(xor(gateChannels[0], gateChannels[1]))
		case "NAND":
			println(nand(gateChannels[0], gateChannels[1]))
		}
	}

	wg.Done()
}

func clock(c3 chan string) {
	x = 0
	println("How much clock pulses per second?")
	var pulsesPerSecond int
	// fmt.Scanln(&pulsesPerSecond)
	pulsesPerSecond = 3
	c3 <- "done"
	for {
		x++
		// println(x)
		time.Sleep(time.Second / time.Duration(pulsesPerSecond))
		x--
	}
}

// prints and sends array of string to readGates
func printer(c1 chan string, c2 chan string, c4 chan []string) {
	fileName := <-c1
	circuit, err := readLines(fileName)
	check(err)
	c4 <- circuit
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

//logic gates
func not(x chan bool) bool {
	return !<-x
}

func and(x, y chan bool) bool {
	return <-x && <-y
}

func or(x, y chan bool) bool {
	if <-x || <-y {
		return true
	}
	return false
}

func nand(x, y chan bool) bool {
	if !and(x, y) {
		return true
	}
	return false
}

func nor(x, y chan bool) bool {
	if or(x, y) == false {
		return true
	}
	return false
}

func xor(x, y chan bool) bool {
	if <-x && !<-y || <-y && !<-x {
		return true
	}
	return false
}

func flipFlop(d, q chan int) {
	// swapping states when clock is either 0 or 1
	var store int
	if x == 1 {
		q <- store
	}
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
