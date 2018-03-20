package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var vertPtr, envPtr, newPtr, oldPtr, notifyPtr, percPtr *string

func init() {
	vertPtr = flag.String("vert", "home", "Vertical name (home|life)")
	envPtr = flag.String("env", "staging", "Vertical environment (staging|prod)")
	newPtr = flag.String("new", "REQUIRED", "New version (will be scaled up)")
	oldPtr = flag.String("old", "REQUIRED", "Old version (will be scaled down)")
	notifyPtr = flag.String("notify", "@jacob.murphy", "List of people to notify")
	percPtr = flag.String("percentage", "100", "Percentage of instances with new version (10|33|50|100)")

	flag.Parse()
}

func main() {

	template, err := ioutil.ReadFile("./internal/templates/event_template.json")
	check(err)

	fmt.Print("Template loaded, checking flags...\n\n")

	// Validating flags
	if *newPtr == "REQUIRED" {
		panic("'new' (New Version) flag is required!")
	}
	if *oldPtr == "REQUIRED" {
		panic("'old' (Old Version) flag is required!")
	}

	fmt.Printf("Flags ok:\n"+
		"\tVertical: %s\n"+
		"\tEnvironment: %s\n"+
		"\tNew Version: %s\n"+
		"\tOld Version: %s\n"+
		"\tNotify: %s\n\n"+
		"\tPercentage: %s\n\n",
		*vertPtr, *envPtr, *newPtr, *oldPtr, *notifyPtr, *percPtr)

	replacer := strings.NewReplacer(
		"{vertical}", *vertPtr,
		"{environment}", *envPtr,
		"{newVersion}", *newPtr,
		"{oldVersion}", *oldPtr,
		"{notify}", *notifyPtr,
		"{percentage}", *percPtr,
	)

	resultFile := fmt.Sprintf("./%s_%s_%s_%s.json", *vertPtr, *envPtr, *newPtr, *percPtr)
	fmt.Println(fmt.Sprintf("Generating '%s'...\n", resultFile))

	result := replacer.Replace(string(template))
	err2 := ioutil.WriteFile(resultFile, []byte(result), 0644)
	check(err2)

	fmt.Println("Done! \\o/")
}
