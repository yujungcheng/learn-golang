package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// + operator
	var string_one string
	var string_two string

	string_one = "String One"
	string_two = "String Two"

	fmt.Println("string_one+string_two: ", string_one+string_two)
	fmt.Println("string_one+space+string_two: ", string_one+" "+string_two)

	// += operator
	var string_three string
	string_three = "String Three"
	string_three += string_one
	fmt.Println("String append: ", string_three)

	// bytes buffer
	var bytes_buffer bytes.Buffer

	bytes_buffer.WriteString("one")
	bytes_buffer.WriteString("two")
	bytes_buffer.WriteString("three")
	fmt.Println("Bytes Buffer: ", bytes_buffer.String())

	bytes_buffer.WriteString(" Four")
	bytes_buffer.WriteString(" Five")
	fmt.Println("Bytes Buffer: ", bytes_buffer.String())

	// sprintf
	sprintf_string := fmt.Sprintf("%s%s", string_one, string_two)
	fmt.Println("Sprintf: ", sprintf_string)

	// join
	string_slice := []string{"string1", "string2", "string3"}
	string_join := strings.Join(string_slice, ", ")
	fmt.Println("String join: ", string_join)

	// Concatenate with int
	number := 10
	text := "emails"
	concatenate_int := fmt.Sprintf("%d %s", number, text)
	fmt.Println(concatenate_int)

	concatenate_int = fmt.Sprint(number, " ", text)
	fmt.Println(concatenate_int)

	concatenate_int = strconv.Itoa(number) + " " + text
	fmt.Println(concatenate_int)

	concatenate_int = strings.Join([]string{strconv.Itoa(number), " ", text}, "")
	fmt.Println(concatenate_int)

	// strings.Builder
	// https://www.calhoun.io/concatenating-and-building-strings-in-go/
	// https://www.calhoun.io/6-tips-for-using-strings-in-go/
	var sb strings.Builder
	strs := []string{"aaa", "bbb", "ccc"}
	for _, r := range strs {
		sb.WriteString(r)
	}
	fmt.Println("strings Builder:", sb.String())

	sb.Reset() // reset builder

	runes := []rune{'1', '2', '3', 'a', 'b', 'c'}
	for _, r := range runes {
		sb.WriteRune(r)
	}
	fmt.Println("strings Builder:", sb.String())

	sb.Reset()

	for i := 5; i >= 1; i-- {
		fmt.Fprintf(&sb, "%d...", i)
	}
	sb.WriteString("ignition")
	fmt.Println(sb.String())
}
