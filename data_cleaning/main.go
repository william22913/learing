package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/william22913/learning/data_cleaning/pkg"
)

func main() {

	dirName := "clean_data"

	regex, err := pkg.ReadFile("./regex.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	files, err := ioutil.ReadDir("./" + dirName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	regexExp, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(files); i++ {
		var temp string

		if files[i].IsDir() {
			continue
		}

		data, err := pkg.ReadFile(dirName + "/" + files[i].Name())
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		splitEnter := strings.Split(data, "\n")

		for j := 0; j < len(splitEnter); j++ {
			if j == 0 {
				temp = temp + splitEnter[j]
			} else {
				if regexExp.Match([]byte(splitEnter[j])) {
					temp = temp + "\n" + splitEnter[j]
				}
			}
		}

		err = pkg.WriteFile(dirName, files[i].Name(), temp)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

	}

}
