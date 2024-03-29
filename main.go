// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"os"

	automl_table "./modules/automl_table"
	dialogflow "./modules/dialogflow"
	language "./modules/natural_language"
	speech_to_text "./modules/speech_to_text"
	text_to_speech "./modules/text_to_speech"
	translation "./modules/translation"
	video "./modules/video"
	vision "./modules/vision"
)

// Usage: `docker run -it --env PROJECT_ID=YOUR_PROJECT_ID golang ./app [DayXX]`
//				 PROJECT_ID is used in Day16, Day25
// Usage: `docker run -it --env PROJECT_NUMBER=XXXXXXX --env MODEL_NAME=XXXXXXXXXXXXXX golang ./app [DayXX]`
//				 used in Day27
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
		translation.ProcTranslateFiles()
	} else if arg == "Day19" {
		// Usage: docker run -v ${PWD}/testdata:/app/testdata -it golang ./app Day19
		text_to_speech.ConvertToSpeech("Hello world")
	} else if arg == "Day20" {
		text_to_speech.ListVoices(os.Stdout)
		text_to_speech.SSMLToSpeech("<speak>The <say-as interpret-as=\"characters\">SSML</say-as>" +
			"standard <break time=\"1s\"/>is defined by the" +
			"<sub alias=\"World Wide Web Consortium\">W3C</sub>.</speak>")
	} else if arg == "Day22" {
		speech_to_text.DemoCode("./testdata/speech_to_text/audio.raw")
	} else if arg == "Day23" {
		speech_to_text.ChineseSpeech("./testdata/speech_to_text/try.mp3")
	} else if arg == "Day25" {
		dialogflow.DetectIntentText("123456789", "嗨", "zh-TW")
	} else if arg == "Day27" {
		automl_table.OnlinePredict()
	}
}
