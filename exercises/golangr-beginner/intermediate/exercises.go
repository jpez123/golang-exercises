package intermediate

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

// FilesExercise - Helper func
//Read line by line
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// FilesExercise - exercise 1
func FilesExercise() {
	//Checks if file exist
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exist")
	}

	//Read a file
	b, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	str := string(b)
	fmt.Println(str)

	//Read line by line
	lines, err := readLines("test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for i, line := range lines {
		fmt.Println(i, line)
	}

	//Write to new/old file
	file, err := os.OpenFile("test.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Writing in file\n")

	// Renames file
	// os.Rename("test.txt", "testRename.txt")
}

// StructExercise - declarations
type House struct {
	price int
	city  string
}

// StructExercise - exercise 2
func StructExercise() {
	//Creating a struct
	var noRooms House

	noRooms.price = 15
	noRooms.city = "Toronto"

	fmt.Println(noRooms)
}

// MapsExercise - exercise 3
func MapsExercise() {
	//Creating nested maps
	website := map[string]map[string]string{
		"Google": map[string]string{
			"name": "Google",
			"type": "Search",
		},
	}

	fmt.Println(website["Google"]["name"])
}

// RandExercise - exercise 4
func RandExercise() {
	//Creates random number
	rand.Seed(time.Now().UnixNano())
	diceRoll := rand.Intn(6-1) + 1
	fmt.Printf("Dice: %d\n", diceRoll)
}

// PointerExercise - exercise 5
func PointerExercise() {
	//Memory addresses
	x := 2
	fmt.Printf("%d %v\n", x, &x)

	//Declare pointer
	var pointer *int = &x
	fmt.Printf("%v %d \n", pointer, *pointer)
}

// SlicesExercise - exercise 6
func SlicesExercise() {
	//Slicing strings
	str := "Golang"
	slice := str[0:2]

	fmt.Println(slice)

	//Slicing arrays
	slice2 := []int{1, 2, 3}
	fmt.Printf("Slice: %v %d %d\n", slice2[1:], len(slice2[1:]), cap(slice2))
}

// Math - methods
type Math struct{}

func (m Math) add(x, y int) int {
	sum := x + y
	return sum
}

// MethodExercise - exercise 7
func MethodExercise() {
	// Call struct method
	var x Math
	fmt.Println(x.add(1, 2))
}

// DeferExercise - exercise 8
func DeferExercise() {
	// Defer call
	defer fmt.Println("Hello")
	fmt.Println("World")
}

// swapValue - method
func swapValue(x, y int) (int, int) {
	return y, x
}

// ReturnsExercise - exercise 8
func ReturnsExercise() {
	//Returns 2 numbers
	fmt.Println(swapValue(1, 2))
}

// printStudents - exercise 9
func printStudents(student ...string) {
	for _, v := range student {
		fmt.Printf("Student: %s\n", v)
	}
}

// VariadicExercise - exercise 9
func VariadicExercise() {
	//Insert x number of students
	printStudents("John", "Mary", "Rachel")
}

// RecursionFunc - exercise 10
func RecursionFunc(x int) int {
	if x > 5 {
		return x
	}
	return RecursionFunc(x + 1)
}

// RecursionExercise - exercise 10
func RecursionExercise() {
	//Calls itself
	fmt.Println(RecursionFunc(1))
}

// GorountinesExercise - exercise 11
func GorountinesExercise() {
	//Concurrent threads
	var x Math
	go fmt.Println(x.add(3, 5))
	fmt.Println(x.add(5, 4))
	fmt.Scanln()
}

// ChannelString - exercise 12
func f(c chan string) {
	c <- "f() was here"
}

func f2(c chan string) {
	msg := <-c
	fmt.Println(msg)
}

// ChannelsExercise - exercise 12
func ChannelsExercise() {
	//Sending message between channels
	var c chan string = make(chan string)
	go f(c)
	go f2(c)

	fmt.Scanln()
}
