package beginner

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// HelloWorld - exercise 1
func HelloWorld() {
	fmt.Println("Hello World")
}

// StringExercise - exercise 2
func StringExercise() {
	//Characters
	str1 := "Golang Example"
	fmt.Printf("The first character is %c == %d \n", str1[0], str1[0])
	//Appending
	str1 += ": Strings!"
	fmt.Printf("%v is %v long\n", str1, len(str1))
	//Joining
	str2 := strings.Join([]string{str1, "Joined!"}, ":")
	fmt.Println(str2)
	//Multiple name variables
	str3, str4 := "John Zhang", "Kiwi Bunny"
	println(str3 + " " + str4)
}

// KeyboardInput - exercise 3
func KeyboardInput() {
	//Reads from keyboard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your city:")
	city, _ := reader.ReadString('\n')
	fmt.Println("You live in " + city)
	//Checks if number is between 0 and 10
	newReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number")
	number, _ := newReader.ReadString('\n')
	number = strings.Replace(number, "\n", "", -1)
	num, _ := strconv.Atoi(number)
	if num > 0 && num < 10 {
		fmt.Println("Within 1-10")
	}
}

// Exercise 4 variable declaration
var (
	variable1            int     = 2
	variable2            float32 = 2.14543
	variable3, variable4         = "string", false
)

// VariablesExercise - exercise 4
func VariablesExercise() {
	//Printing multiple strings (%c)
	fmt.Printf("%s %s \n", variable3, variable3)

	//Pninting multiple numbers (%d, %f)
	fmt.Printf("%d %.2f \n", variable1, variable2)

	//Exchanging numbers
	variable1, variable2 = int(variable2), float32(variable1)
	fmt.Printf("%d %f \n", variable1, variable2)

	//Calculate year given age & date of birth
	var birthyear int

	today := time.Now()
	todayMonth := today.Month()
	todayDay := today.Day()
	todayYear := today.Year()

	age := 31
	birthday := "10/10"

	birthdayArr := strings.Split(birthday, "/")
	birthdayMonth, _ := strconv.Atoi(birthdayArr[0])
	birthdayDay, _ := strconv.Atoi(birthdayArr[1])
	if int(todayMonth) <= birthdayMonth && int(todayDay) <= birthdayDay {
		birthyear = todayYear - age - 1
	} else {
		birthyear = todayYear - age
	}

	fmt.Println(birthyear)
}

// ArrayExercise - exercise 5
func ArrayExercise() {
	//Declaration of new array
	arr := new([1]int)
	arr1 := [...]int{1, 2, 3}
	fmt.Println(*arr)
	fmt.Println(arr1)

	//Array with numbers & strings
	numArray := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nameArray := [3]string{"John", "Rachel", "Jeff"}
	fmt.Printf("%d %s \n", numArray[3], nameArray[2])
}

// LoopExercise - exercise 6
func LoopExercise() {
	//Looping through arrays
	arr := [5]int{1, 2, 3, 4, 5}
	for k, v := range arr {
		fmt.Printf("Arr loop: %d %d\n", k, v)
	}

	//Nested loops
	for i := 0; i < 5; i++ {
		for x := 1; x < 3; x++ {
			fmt.Printf("Nested loop: %d %d \n", i, x)
		}
	}

	//While loop
	i := 0
	for i < len(arr) {
		i++
		fmt.Println(i)
	}
}

// IfExercise - exercise 7
func IfExercise() {
	//Divide by zero if greater than 0
	randomNum := rand.Intn(10-1) * 1
	if randomNum > 0 {
		dividedNum := randomNum / 2
		fmt.Printf("Random num: %d %d\n", randomNum, dividedNum)
	}
}
