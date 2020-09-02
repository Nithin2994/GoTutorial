package funcs

import (
	"fmt"
	"strconv"
)

//Print ...
func Print(s string, num int, tags []string) {
	fmt.Println("Name : ", s)
	fmt.Println("Id   : ", num)
	fmt.Println("Tags :", tags)
	s = "Nithin111"
}

//FormatedPrintStatement ...
func FormatedPrintStatement(s string, num int, tags []string) (string, int) {
	var result string
	result = s + " with Id : " + strconv.Itoa(num)
	result += "["
	for _, tag := range tags {
		result += tag + ","
	}
	result += "]"
	return result, 10
}

//AddTag ...
func AddTag(tags *[]string, newTag string) {
	*tags = append(*tags, newTag)
}
