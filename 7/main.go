package main

import (
	"fmt"
	"github.com/simpsonw/aoc-2022/utils"
	"log"
	"regexp"
)

type File struct {
	Name     string
	Parent   *File
	Size     int
	IsDir    bool
	Children []*File
}

func (f File) String() string {
	if f.IsDir {
		return fmt.Sprintf("%s (dir)", f.Name)
	} else {
		return fmt.Sprintf("%s (file, %d)", f.Name, f.Size)
	}
}

func (f *File) hasChildWithName(name string, IsDir bool) bool {
	for _, v := range f.Children {
		if v.Name == name && v.IsDir == IsDir {
			return true
		}
	}
	return false
}

type Command struct {
	Name     string
	Argument string
}

var root File
var pwd *File
var sum int
var targetDirectorySize int
var freeDiskSpace int

const TOTAL_DISK_SPACE = 70000000
const REQUIRED_DISK_SPACE = 30000000

func main() {
	root = File{
		Name:  "/",
		IsDir: true,
	}
	pwd = &root
	lines := utils.GetLines()
	for _, l := range lines {
		if l == "" {
			break
		}
		cmd := parseCommand(l)
		if cmd != nil {
			runCmd(cmd)
		} else {
			// line is output
			processOutput(l)
		}
	}
	freeDiskSpace = TOTAL_DISK_SPACE - traverseHelper(&root)
	fmt.Printf("Current free disk space: %d\n", freeDiskSpace)
	traverse(&root, 0)
	fmt.Printf("Sum %d\n", sum)
}

func traverse(f *File, level int) {
	for _, v := range f.Children {
		if v.IsDir {
			traverse(v, level+1)
		}
	}
	totalSize := traverseHelper(f)

	if (freeDiskSpace + totalSize) >= REQUIRED_DISK_SPACE {
		if targetDirectorySize == 0 {
			targetDirectorySize = totalSize
		} else if targetDirectorySize > totalSize {
			targetDirectorySize = totalSize
			fmt.Printf("Deleting %s (size: %d) will result in %d free disk space, which is more than %d\n", f, totalSize, freeDiskSpace+totalSize, REQUIRED_DISK_SPACE)
		}
	}
	// Part 1
	//if totalSize < 100000 {
	//	sum += totalSize
	//	fmt.Printf("%s has a size less than 100000 (%d)\n", f, totalSize)
	//}
}

func traverseHelper(f *File) (totalSize int) {
	for _, v := range f.Children {
		if v.IsDir == false {
			totalSize += v.Size
		}
	}
	for _, v := range f.Children {
		if v.IsDir {
			totalSize += traverseHelper(v)
		}
	}
	return totalSize
}

func processOutput(l string) {
	name, err := parseDirectory(l)
	if err == nil && !pwd.hasChildWithName(name, true) {
		pwd.Children = append(pwd.Children, &File{Name: name, IsDir: true, Parent: pwd})
		return
	}
	name, size, err := parseFile(l)
	if err == nil && !pwd.hasChildWithName(name, false) {
		pwd.Children = append(pwd.Children, &File{Name: name, Size: size})
		return
	}
}

func parseDirectory(l string) (string, error) {
	var directoryName string
	_, err := fmt.Sscanf(l, "dir %s", &directoryName)
	return directoryName, err
}

func parseFile(l string) (string, int, error) {
	var name string
	var size int
	_, err := fmt.Sscanf(l, "%d %s", &size, &name)
	return name, size, err
}

func runCmd(cmd *Command) {
	switch cmd.Name {
	case "cd":
		cd(cmd)
	case "ls":
		return
	default:
		log.Fatalf("Unrecognized command: %s\n", cmd.Name)
	}
}

func cd(cmd *Command) {
	switch cmd.Argument {
	case "..":
		if pwd.Parent != nil {
			pwd = pwd.Parent
		}
	case "/":
		pwd = &root
	default:
		for _, v := range pwd.Children {
			if v.IsDir && v.Name == cmd.Argument {
				pwd = v
			}
		}
	}
	//fmt.Printf("%s %s (pwd: %s)\n", cmd.Name, cmd.Argument, pwd)
}

func parseCommand(line string) (cmd *Command) {
	cmd = &Command{}
	r := regexp.MustCompile(`^\$\s+(?P<cmd>\S+)\s*(?P<arg>\S+)?`)
	matches := r.FindStringSubmatch(line)
	if len(matches) == 0 {
		return nil
	}
	cmdIndex := r.SubexpIndex("cmd")
	cmd.Name = matches[cmdIndex]
	argIndex := r.SubexpIndex("arg")
	if argIndex > -1 {
		cmd.Argument = matches[argIndex]
	}
	return
}
