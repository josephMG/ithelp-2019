package translation

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ProcTranslateFiles() {
	var (
		root            string = "./testdata/translate"
		validationCount int    = 0
		testCount       int    = 0
		trainCount      int    = 0
	)
	files, err := ioutil.ReadDir(root + "/en")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if _, err := os.Stat(root + "/zh/" + file.Name()); err == nil {
			// 先把有英文也有中文的檔案，分別讀到`enLines`、`zhLines`
			var enLines, enErr = readLines(root + "/en/" + file.Name())
			var zhLines, zhErr = readLines(root + "/zh/" + file.Name())
			if enErr != nil || zhErr != nil {
				continue
			}
			// 刪掉行數不一樣的部分
			enLines, zhLines = normalizeLines(enLines, zhLines)

			// 塞入tsv data
			if trainCount < 11000 {
				trainCount += len(enLines)
				writeCSV(root+"/train.tsv", enLines, zhLines)
			} else if testCount < 1000 {
				testCount += len(enLines)
				writeCSV(root+"/test.tsv", enLines, zhLines)
			} else if validationCount < 1000 {
				validationCount += len(enLines)
				writeCSV(root+"/validation.tsv", enLines, zhLines)
			} else {
				fmt.Println("Done!")
				break
			}
		}
	}
}

func writeCSV(filename string, enLines []string, zhLines []string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	defer writer.Flush()
	var data = [][]string{}
	for i := 0; i < len(enLines); i++ {
		data = append(data, []string{enLines[i], zhLines[i]})
	}
	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func normalizeLines(enLines []string, zhLines []string) ([]string, []string) {
	var (
		enArrP  []int
		zhArrP  []int
		enIndex int = 0
		zhIndex int = 0
	)
	for i := 0; i < len(enLines); i++ {
		if i == len(enLines)-1 {
			enArrP = append(enArrP, i)
		} else if enLines[i] == "<P>" {
			enArrP = append(enArrP, i)
		}
	}
	for i := 0; i < len(zhLines); i++ {
		if i == len(zhLines)-1 {
			zhArrP = append(zhArrP, i)
		} else if zhLines[i] == "<P>" {
			zhArrP = append(zhArrP, i)
		}
	}
	if len(zhArrP) != len(enArrP) {
		enLines = nil
		zhLines = nil
		return nil, nil
	}
	if len(zhArrP) == 0 || len(enArrP) == 0 {
		return nil, nil
	}
	if enIndex == 0 && enArrP[0] != 0 {
		assignRange(enLines, enIndex, enArrP[0])
	}
	if zhIndex == 0 && zhArrP[0] != 0 {
		assignRange(zhLines, zhIndex, zhArrP[0])
	}
	for i := 0; i < len(zhArrP); i++ {
		if i == len(zhArrP)-1 {
			zhIndex = len(zhLines)
			enIndex = len(enLines)
		} else {
			zhIndex = zhArrP[i+1]
			enIndex = enArrP[i+1]
		}
		if zhIndex-zhArrP[i] != enIndex-enArrP[i] {
			assignRange(zhLines, zhArrP[i], zhIndex)
			assignRange(enLines, enArrP[i], enIndex)
		}
	}
	return deleteP(enLines), deleteP(zhLines)
}

func assignRange(lines []string, start int, end int) {
	for i := start; i < end; i++ {
		lines[i] = "<P>"
	}
}

func deleteP(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "<P>" {
			r = append(r, str)
		}
	}
	return r
}
