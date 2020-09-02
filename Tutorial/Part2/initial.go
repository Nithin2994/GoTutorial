package main

import (
	"fmt"
	"strconv"
)

//ClassID ....
var ClassID = "4th class"

//StudentsType ..
type StudentsType []string

func main1() {

	var title string // ""

	title = "Nithin1"
	title1 := "Nithin"

	fmt.Println(title)
	fmt.Println(title1)
	fmt.Println(ClassID)

	var rollNum int = 10 // 0
	var percent float32  // 0

	fmt.Println(title + " " + strconv.Itoa(rollNum))

	_, err := strconv.ParseInt("123", rollNum, 32)
	if err != nil {
		fmt.Println("Error in conversion ", err)
	}

	var isFresher bool // false
	isFresher = true

	fmt.Println(rollNum, percent, isFresher)

	// [1,2,3,4,5] ,["a","b","c"]
	//  0,1,2,3,4
	var subjects [5]string
	subjects[0] = "Maths"
	subjects[1] = "Science"
	subjects[2] = "English"
	subjects[3] = "Social"
	subjects[4] = "GK"

	var students StudentsType
	students = make([]string, 0)
	students = append(students, "Nithin")
	students = append(students, "Nithin1")
	students = append(students, "Nithin2")
	students = append(students, "Nithin3")

	fmt.Println("length of students : ", len(students))

	var score map[string]int
	score = make(map[string]int, 0)
	score["Nithin"] = 60
	score["Nithin1"] = 60

	const totalMarks = "Nithin"

	for i, v := range students {
		fmt.Println(i, v)
	}
	var j int = 0
	for j = 10; j < 10; j++ {
		fmt.Println("j : ", j)
	}
	// for i, v := range subjects {
	// 	fmt.Println(i, v)
	// }

	// for key, value := range title1 {
	// 	fmt.Println(key, value)
	// }

	var i int = 0

	i += 2 // i = i + 2 , -= , *=
	for {
		i++
		fmt.Println("i :", i)

		if i%2 == 0 && i%4 == 0 {
			//break
			fmt.Println("Even")
		} else if i%3 == 0 {
			fmt.Println("xyz")
		} else {
			fmt.Println("Odd")
		}
		if i == 10 {
			break
		}
	}

	switch i {
	case 10:
		fmt.Println("> num is 10")
		break
	case 20:
		fmt.Println(">>num is 20")
		break
	default:
		fmt.Println("num is ", i)
	}

	fmt.Println("hello")
}
