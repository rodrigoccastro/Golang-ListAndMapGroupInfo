package main

import (
	"fmt"
	"strings"
)

func main() {
	paths := []string{
        "root/a 1.txt(abcd) 2.txt(efgh)",
        "root/c 3.txt(abcd)",
        "root/c/d 4.txt(efgh)",
        "root 4.txt(efgh)",
    }
    result := findDuplicate(paths)
    fmt.Println(result)
}

func findDuplicate(paths []string) [][]string {
    // Input: paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"]
    // Output: [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]

    var output [][]string
    contentMap := make(map[string][]string)

	for _, path := range paths {
        parts := strings.Split(path, " ")
        directory := parts[0]
        for i := 1; i < len(parts); i++ {
			fileName, fileContent := getNameAndContent(directory, parts[i])
            contentMap[fileContent] = append(contentMap[fileContent], fileName)
        }
    }

	for _, files := range contentMap {
		output = append(output, files)
    }

    return output
}

func getNameAndContent(directory string, str string) (string, string) {
	fileContent := strings.Split(str, "(")[1]
	fileContent = strings.TrimRight(fileContent, ")")
	fileName := directory + "/" + strings.Split(str, "(")[0]
    return fileName, fileContent
}