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
	day := 7
	fmt.Println(part1(day))
}

func part1(day int) int {
	filename := fmt.Sprintf("day%d/day%d.txt", day, day)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	/////////////////////////////
	root := Directory{
		Name:            "/",
		Files:           nil,
		ParentDirectory: nil,
		SubDirectories:  make(map[string]Directory),
		Size:            0,
	}
	/////////////////////////////
	curr := root
	addMode := false
	for {
		line, err := reader.ReadString('\n')
		line = trimSpace(line)
		/////////////////////////////
		if line == "$ cd /" {
			addMode = false
			curr = root
		} else if line == "$ ls" {
			//add mode
			addMode = true
		} else if line == "$ cd .." {
			addMode = false
			curr = *curr.ParentDirectory
		} else if line[:4] == "$ cd" {
			addMode = false
			sub := line[5:]
			curr = curr.SubDirectories[sub]
		} else if addMode {
			add := strings.Split(line, " ")
			if add[0] == "dir" {
				curr.AddDirectory(add[1])
			} else {
				size := num(add[0])
				curr.AddFile(add[1], size)
			}
		} else {
			fmt.Println("What do I do with " + line)
		}
		/////////////////////////////
		if err != nil {
			break
		} //EOF
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
	/////////////////////////////
	fmt.Printf("%#v", root)
	return 0
	/////////////////////////////
}

func trimSpace(s string) string {
	return string(bytes.TrimSpace([]byte(s)))
}

type Directory struct {
	Name            string
	Files           []File
	ParentDirectory *Directory
	SubDirectories  map[string]Directory
	Size            int
}

func (this *Directory) AddDirectory(name string) {
	subDirectory := Directory{
		Name:            name,
		Files:           nil,
		ParentDirectory: this,
		SubDirectories:  make(map[string]Directory),
		Size:            0,
	}
	this.SubDirectories[name] = subDirectory
}

func (this *Directory) AddFile(name string, size int) {
	f := File{Name: name, Size: size}
	this.increaseSize(size)
	this.Files = append(this.Files, f)
}

func (this *Directory) increaseSize(inc int) {
	this.Size += inc
	this.ParentDirectory.increaseSize(inc)
}

type File struct {
	Name string
	Size int
}

func num(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
