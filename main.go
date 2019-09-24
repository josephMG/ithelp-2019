// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	language "./modules/natural_language"
	translation "./modules/translation"
	video "./modules/video"
	vision "./modules/vision"
)

// Usage: `docker run -it --env PROJECT_ID=YOUR_PROJECT_ID golang ./app [DayXX]`
//				 PROJECT_ID is used in Day16
func main() {
	arg := os.Args[1]

	if arg == "Day3" {
		vision.DetectLabel(os.Stdout, "./testdata/furniture.jpg")
	} else if arg == "Day4" {
		vision.DetectText(os.Stdout, "./testdata/las-vegas.jpg")
		vision.DetectFaces(os.Stdout, "./testdata/face.jpg")
	} else if arg == "Day7" {
		video.DemoCode(os.Stdout, "gs://cloud-samples-data/video/cat.mp4")
	} else if arg == "Day8" {
		video.ShotChangeURI(os.Stdout, "gs://cloud-samples-data/video/gbikes_dinosaur.mp4")
		video.TextDetectionGCS(os.Stdout, "gs://python-docs-samples-tests/video/googlework_short.mp4")
	} else if arg == "Day11" {
		language.DemoCode(os.Stdout, "Hello World")
	} else if arg == "Day12" {
		language.AnalyzeSentiment("gs://cloud-samples-tests/natural-language/gettysburg.txt")
		// language.AnalyzeSentiment("Four score and seven years ago our fathers brought forth on this continent, a new nation, conceived in Liberty, and dedicated to the proposition that all men are created equal. Now we are engaged in a great civil war, testing whether that nation, or any nation so conceived and so dedicated, can long endure. We are met on a great battle-field of that war. We have come to dedicate a portion of that field, as a final resting place for those who here gave their lives that that nation might live. It is altogether fitting and proper that we should do this. But, in a larger sense, we can not dedicate—we can not consecrate—we can not hallow—this ground. The brave men, living and dead, who struggled here, have consecrated it, far above our poor power to add or detract. The world will little note, nor long remember what we say here, but it can never forget what they did here. It is for us the living, rather, to be dedicated here to the unfinished work which they who fought here have thus far so nobly advanced. It is rather for us to be here dedicated to the great task remaining before us—that from these honored dead we take increased devotion to that cause for which they gave the last full measure of devotion—that we here highly resolve that these dead shall not have died in vain—that this nation, under God, shall have a new birth of freedom—and that government of the people, by the people, for the people, shall not perish from the earth.")
	} else if arg == "Day15" {
		translation.TranslateText([]string{"Dr. Watson, come here!", "Bring me some coffee!"})
	} else if arg == "Day16" {
		translation.TranslateTextV3Beta1([]string{"Dr. Watson, come here!", "Bring me some coffee!"})
	} else if arg == "Day17" {
		// Please download http://data.statmt.org/news-commentary/v14/ en and zh files to testdata/translate
		// Usage: docker run -v ${PWD}/testdata:/app/testdata -it golang ./app Day17
		procTranslateFiles()
	}
}

func procTranslateFiles() {
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
