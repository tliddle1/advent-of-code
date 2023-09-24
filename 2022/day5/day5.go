package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	day := 5
	fmt.Println(part1(day), part2(day))
}

func part1(day int) string {
	filename := fmt.Sprintf("day%d/day%d.txt", day, day)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	line1, err := reader.ReadString('\n')
	line2, err := reader.ReadString('\n')
	line3, err := reader.ReadString('\n')
	line4, err := reader.ReadString('\n')
	line5, err := reader.ReadString('\n')
	line6, err := reader.ReadString('\n')
	line7, err := reader.ReadString('\n')
	line8, err := reader.ReadString('\n')
	s1 := NewStringStack()
	s1.Push(string(rune(int(line8[1]))))
	s1.Push(string(rune(int(line7[1]))))
	s1.Push(string(rune(int(line6[1]))))
	s1.Push(string(rune(int(line5[1]))))
	s2 := NewStringStack()
	s2.Push(string(rune(int(line8[5]))))
	s2.Push(string(rune(int(line7[5]))))
	s2.Push(string(rune(int(line6[5]))))
	s2.Push(string(rune(int(line5[5]))))
	s2.Push(string(rune(int(line4[5]))))
	s2.Push(string(rune(int(line3[5]))))
	s2.Push(string(rune(int(line2[5]))))
	s3 := NewStringStack()
	s3.Push(string(rune(int(line8[9]))))
	s3.Push(string(rune(int(line7[9]))))
	s3.Push(string(rune(int(line6[9]))))
	s3.Push(string(rune(int(line5[9]))))
	s3.Push(string(rune(int(line4[9]))))
	s3.Push(string(rune(int(line3[9]))))
	s3.Push(string(rune(int(line2[9]))))
	s3.Push(string(rune(int(line1[9]))))
	s4 := NewStringStack()
	s4.Push(string(rune(int(line8[13]))))
	s4.Push(string(rune(int(line7[13]))))
	s4.Push(string(rune(int(line6[13]))))
	s4.Push(string(rune(int(line5[13]))))
	s4.Push(string(rune(int(line4[13]))))
	s4.Push(string(rune(int(line3[13]))))
	s5 := NewStringStack()
	s5.Push(string(rune(int(line8[17]))))
	s5.Push(string(rune(int(line7[17]))))
	s5.Push(string(rune(int(line6[17]))))
	s5.Push(string(rune(int(line5[17]))))
	s5.Push(string(rune(int(line4[17]))))
	s6 := NewStringStack()
	s6.Push(string(rune(int(line8[21]))))
	s6.Push(string(rune(int(line7[21]))))
	s6.Push(string(rune(int(line6[21]))))
	s6.Push(string(rune(int(line5[21]))))
	s6.Push(string(rune(int(line4[21]))))
	s6.Push(string(rune(int(line3[21]))))
	s6.Push(string(rune(int(line2[21]))))
	s6.Push(string(rune(int(line1[21]))))
	s7 := NewStringStack()
	s7.Push(string(rune(int(line8[25]))))
	s7.Push(string(rune(int(line7[25]))))
	s7.Push(string(rune(int(line6[25]))))
	s8 := NewStringStack()
	s8.Push(string(rune(int(line8[29]))))
	s8.Push(string(rune(int(line7[29]))))
	s8.Push(string(rune(int(line6[29]))))
	s8.Push(string(rune(int(line5[29]))))
	s8.Push(string(rune(int(line4[29]))))
	s8.Push(string(rune(int(line3[29]))))
	s8.Push(string(rune(int(line2[29]))))
	s9 := NewStringStack()
	s9.Push(string(rune(int(line8[33]))))
	s9.Push(string(rune(int(line7[33]))))
	s9.Push(string(rune(int(line6[33]))))
	s9.Push(string(rune(int(line5[33]))))
	s9.Push(string(rune(int(line4[33]))))
	s9.Push(string(rune(int(line3[33]))))
	s9.Push(string(rune(int(line2[33]))))
	s9.Push(string(rune(int(line1[33]))))
	ship := Ship{Cargo: []StringStack{*s1, *s2, *s3, *s4, *s5, *s6, *s7, *s8, *s9}}
	_, _ = reader.ReadString('\n')
	_, _ = reader.ReadString('\n')
	/////////////////////////////
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		l := strings.Split(line, " ")
		ship.Move(num(l[1]), num(l[3]), num(l[5]))
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	/////////////////////////////
	result := ""
	for _, stack := range ship.Cargo {
		result += stack.Pop()
	}
	return result
	/////////////////////////////
}

func part2(day int) string {
	filename := fmt.Sprintf("day%d/day%d.txt", day, day)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	line1, err := reader.ReadString('\n')
	line2, err := reader.ReadString('\n')
	line3, err := reader.ReadString('\n')
	line4, err := reader.ReadString('\n')
	line5, err := reader.ReadString('\n')
	line6, err := reader.ReadString('\n')
	line7, err := reader.ReadString('\n')
	line8, err := reader.ReadString('\n')
	s1 := NewStringStack()
	s1.Push(string(rune(int(line8[1]))))
	s1.Push(string(rune(int(line7[1]))))
	s1.Push(string(rune(int(line6[1]))))
	s1.Push(string(rune(int(line5[1]))))
	s2 := NewStringStack()
	s2.Push(string(rune(int(line8[5]))))
	s2.Push(string(rune(int(line7[5]))))
	s2.Push(string(rune(int(line6[5]))))
	s2.Push(string(rune(int(line5[5]))))
	s2.Push(string(rune(int(line4[5]))))
	s2.Push(string(rune(int(line3[5]))))
	s2.Push(string(rune(int(line2[5]))))
	s3 := NewStringStack()
	s3.Push(string(rune(int(line8[9]))))
	s3.Push(string(rune(int(line7[9]))))
	s3.Push(string(rune(int(line6[9]))))
	s3.Push(string(rune(int(line5[9]))))
	s3.Push(string(rune(int(line4[9]))))
	s3.Push(string(rune(int(line3[9]))))
	s3.Push(string(rune(int(line2[9]))))
	s3.Push(string(rune(int(line1[9]))))
	s4 := NewStringStack()
	s4.Push(string(rune(int(line8[13]))))
	s4.Push(string(rune(int(line7[13]))))
	s4.Push(string(rune(int(line6[13]))))
	s4.Push(string(rune(int(line5[13]))))
	s4.Push(string(rune(int(line4[13]))))
	s4.Push(string(rune(int(line3[13]))))
	s5 := NewStringStack()
	s5.Push(string(rune(int(line8[17]))))
	s5.Push(string(rune(int(line7[17]))))
	s5.Push(string(rune(int(line6[17]))))
	s5.Push(string(rune(int(line5[17]))))
	s5.Push(string(rune(int(line4[17]))))
	s6 := NewStringStack()
	s6.Push(string(rune(int(line8[21]))))
	s6.Push(string(rune(int(line7[21]))))
	s6.Push(string(rune(int(line6[21]))))
	s6.Push(string(rune(int(line5[21]))))
	s6.Push(string(rune(int(line4[21]))))
	s6.Push(string(rune(int(line3[21]))))
	s6.Push(string(rune(int(line2[21]))))
	s6.Push(string(rune(int(line1[21]))))
	s7 := NewStringStack()
	s7.Push(string(rune(int(line8[25]))))
	s7.Push(string(rune(int(line7[25]))))
	s7.Push(string(rune(int(line6[25]))))
	s8 := NewStringStack()
	s8.Push(string(rune(int(line8[29]))))
	s8.Push(string(rune(int(line7[29]))))
	s8.Push(string(rune(int(line6[29]))))
	s8.Push(string(rune(int(line5[29]))))
	s8.Push(string(rune(int(line4[29]))))
	s8.Push(string(rune(int(line3[29]))))
	s8.Push(string(rune(int(line2[29]))))
	s9 := NewStringStack()
	s9.Push(string(rune(int(line8[33]))))
	s9.Push(string(rune(int(line7[33]))))
	s9.Push(string(rune(int(line6[33]))))
	s9.Push(string(rune(int(line5[33]))))
	s9.Push(string(rune(int(line4[33]))))
	s9.Push(string(rune(int(line3[33]))))
	s9.Push(string(rune(int(line2[33]))))
	s9.Push(string(rune(int(line1[33]))))
	ship := Ship{Cargo: []StringStack{*s1, *s2, *s3, *s4, *s5, *s6, *s7, *s8, *s9}}
	_, _ = reader.ReadString('\n')
	_, _ = reader.ReadString('\n')
	/////////////////////////////
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		l := strings.Split(line, " ")
		ship.Move2(num(l[1]), num(l[3]), num(l[5]))
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	/////////////////////////////
	result := ""
	for _, stack := range ship.Cargo {
		result += stack.Pop()
	}
	return result
	/////////////////////////////
}

func trimSpace(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}

func num(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type Ship struct {
	Cargo []StringStack
}

func (s *Ship) Move(count, s1, s2 int) {

	for i := 0; i < count; i++ {
		s.Cargo[s2-1].Push(s.Cargo[s1-1].Pop())
	}
}

func (s *Ship) Move2(count, s1, s2 int) {
	stack := *NewStringStack()
	for i := 0; i < count; i++ {
		stack.Push(s.Cargo[s1-1].Pop())
	}
	for j := 0; j < count; j++ {
		s.Cargo[s2-1].Push(stack.Pop())
	}
}

// StringStack represents a stack that only accepts strings.
type StringStack struct {
	items []string
}

// Push adds a string element to the top of the stack.
func (s *StringStack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the string element from the top of the stack.
func (s *StringStack) Pop() string {
	if len(s.items) == 0 {
		return ""
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

// NewStringStack creates and returns a new string stack.
func NewStringStack() *StringStack {
	return &StringStack{items: make([]string, 0)}
}
