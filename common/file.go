package common

import (
	"bufio"
	"os"
)

func ReadAllLines(path string) ([]string, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    sc := bufio.NewScanner(f)
	var lines []string
    for sc.Scan() {
        lines = append(lines, sc.Text())
    }
    if err := sc.Err(); err != nil {
        return nil, err 
	}
	return lines, nil
}

func ReverseString(s string) string{
    n := len(s)
    runes := make([]rune, n)
    for _, rune := range s {
        n--
        runes[n] = rune
    }
    return string(runes[n:])
}
