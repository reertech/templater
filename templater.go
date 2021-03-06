package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"text/template"
)

var templateFileName string
var outputFileName string

func init() {
	flag.StringVar(&templateFileName, "template", "", "template file or - for stdin")
	flag.StringVar(&templateFileName, "t", "", "template file (shorthand)")

	flag.StringVar(&outputFileName, "output", "-", "output file")
	flag.StringVar(&outputFileName, "o", "-", "output file (shorthand)")
}

func parseInt(str string) int {
	v, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		log.Fatal("Failed to parse the template: ", err)
		return 0
	}
	return int(v)
}

func seq(until int) []int {
	arr := make([]int, until)
	for i := 0; i < until; i++ {
		arr[i] = i
	}
	return arr
}

func include(fileName string) string {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Failed to include the file: ", err)
		return ""
	}
	return string(contents)
}

func getTemplateContent() string {
	if templateFileName == "-" {
		templateContent, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Failed to read the template from stdin: ", err)
			return ""
		}
		return string(templateContent)
	} else {
		templateContent, err := ioutil.ReadFile(templateFileName)
		if err != nil {
			log.Fatal("Failed to read the template from file: ", err)
			return ""
		}
		return string(templateContent)
	}
}

func getOutputFile() *os.File {
	if outputFileName == "-" {
		return os.Stdout
	} else {
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			log.Fatal("Failed to create the output file: ", err)
		}
		return outputFile
	}
}

func main() {
	flag.Parse()

	if templateFileName == "" {
		log.Fatal("Please define template name with the -template option. Use -h to see others.")
		return
	}

	funcMap := template.FuncMap{
		"env":      os.Getenv,
		"parseInt": parseInt,
		"seq":      seq,
		"include":  include,
	}

	t, err := template.New("Template").Funcs(funcMap).Parse(getTemplateContent())

	if err != nil {
		log.Fatal("Failed to parse the template: ", err)
		return
	}

	outputFile := getOutputFile()
	err = t.Execute(outputFile, "")
	if err != nil {
		log.Fatal("Failed to render the template: ", err)
		return
	}

	outputFile.Close()
}
