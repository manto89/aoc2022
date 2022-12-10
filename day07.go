package main

import (
	"fmt"
	"github.com/manto89/aoc2022/common"
)

type Day07 struct {
	inputPath string
}

func (d *Day07) executePart1() string {
	lines, err := common.ReadAllLines(d.inputPath)
	if err != nil {
        return fmt.Sprintf("open file error: %v", err)
	}
  workingDir, err := getDirectoriesFromLines(lines)
  if err != nil{
        return fmt.Sprintf("parse file error: %v", err)
  }
  //Navigate to root
  for workingDir.parent != nil{
    workingDir = workingDir.parent
  }
  //Calculate size of all directories
  for _, fd := range(workingDir.children){
    if fd.fileType == Directory{
      fd.size = calculateDirSize(fd.children)
    }
  }
  dirUnder100k := extractDirectoriesUnder100k(*workingDir)
  sizeSum := 0
  for _, d := range(dirUnder100k){
    sizeSum += d.size
  }
  return fmt.Sprintf("Total size of directories under 100k: %d", sizeSum)

}
func getDirectoriesFromLines(lines []string) (*FD, error){
  var workingDir *FD
	for i, line := range(lines){
		if string(line[0]) == "$"{
      command, destination, err := decodeCommand(line)
      if err != nil {
        return workingDir, fmt.Errorf("error decoding line %d: %v", i, err)
      }
      if command == "cd" {
        var newWorkingDir *FD
        if workingDir == nil {
          newWorkingDir = &FD{
          name: destination,
          }
        } else {
          if destination == ".." {
            newWorkingDir = workingDir.parent
          } else {
            for _, fd := range(workingDir.children) {
              if fd.name == destination {
                newWorkingDir = fd
                break
              }
            }
            if newWorkingDir.name == "" {
              return workingDir, fmt.Errorf("error moving to directory %s, working directory name: %s , children: %v", destination, workingDir.name, workingDir.children)
            }
          }
        }
        workingDir = newWorkingDir

        } 
		} else {
      fd, err := decodeLsOutput(line, workingDir)
      if err != nil{
        return workingDir, fmt.Errorf("error decoding line %d: %v", i, err)
      }
      workingDir.children = append(workingDir.children, fd)
    }
  }
  return workingDir, nil
}

func extractDirectoriesUnder100k(fd FD) ([]FD) {
  var ret []FD
  for _, child := range(fd.children){
    if child.fileType == Directory{
      if child.size < 100001 {
        ret = append(ret, *child)
      }
      newDirs := extractDirectoriesUnder100k(*child)
      ret = append(ret, newDirs...)
    }
  }
  return ret
}

func calculateDirSize(children []*FD) (int){
  var size int
  for _, child := range(children){
    if child.fileType == Directory{
      child.size = calculateDirSize(child.children)
    }
    size += child.size
  }
  return size
}


type FileType int

const ( 
  File FileType = iota
  Directory
)

type FD struct {
  parent *FD
  name string
  fileType FileType 
  size int
  children  []*FD
}

func decodeCommand (line string) (string, string, error){
  // cd / or cd abc
  var command string
  var dest string
	fmt.Sscanf(line, "$ %s %s", &command, &dest)
  if command != "cd" && command != "ls"{
    return command, dest, fmt.Errorf("cannot decode %s", line[2:4])
  }
  return command, dest, nil
}

func decodeLsOutput(line string, parent *FD) (*FD, error){
  var ret FD
  var name string
  var size int
  readParams, _ := fmt.Sscanf(line, "%d %s", &size, &name)
  if readParams < 2 {
    _, err := fmt.Sscanf(line, "dir %s", &name)
    if err != nil {
      return nil, err
    }
    ret = FD{
      fileType: Directory,
      name: name,
      parent: parent,
    }
  } else {
    ret = FD{
      fileType: File,
      name: name, 
      parent: parent,
      size: size,
    }
  }
  return &ret, nil
}


func (d *Day07) executePart2() string {
	return "Not implemented"
}
