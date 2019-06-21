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

var vertPtr, envPtr, newPtr, oldPtr, winnerPtr, loserPtr, percPtr *string

func init() {
	vertPtr = flag.String("vert", "life", "Vertical name (auto|home|life)")
	envPtr = flag.String("env", "staging", "Vertical environment (staging|prod)")
	newPtr = flag.String("new", "REQUIRED", "New version (will be scaled up)")
	oldPtr = flag.String("old", "REQUIRED", "Old version (will be scaled down)")
	winnerPtr = flag.String("winner", "b", "Winner leg (100% A/B)")
	loserPtr = flag.String("loser", "a", "Loser leg (100% A/B)")

	flag.Parse()
}

func main() {
	fmt.Print("Checking flags...\n\n")

	// Validating flags
	if *newPtr == "REQUIRED" {
		panic("'new' (New Version) flag is required!")
	}
	if *oldPtr == "REQUIRED" {
		panic("'old' (Old Version) flag is required!")
	}

	templateFile := "event_template_staging.json"
	if *envPtr == "prod" {
		templateFile = "event_template_prod.json"
	}
	template, err := ioutil.ReadFile("./internal/templates/" + templateFile)
	check(err)
	fmt.Print("Template loaded...\n\n")

	fmt.Printf("Flags ok:\n"+
		"\tVertical: %s\n"+
		"\tEnvironment: %s\n"+
		"\tOld Version: %s\n"+
		"\tNew Version: %s\n"+
		"\tLoser Leg: %s\n"+
		"\tWinner Leg: %s\n\n",
		*vertPtr, *envPtr, *oldPtr, *newPtr, *loserPtr, *winnerPtr)

	var vertical string
	if *vertPtr != "auto" {
		vertical = fmt.Sprintf("%s-", *vertPtr)
	}

	replacer := strings.NewReplacer(
		"{verticalName}", *vertPtr,
		"{vertical}", vertical,
		"{environment}", *envPtr,
		"{newVersion}", *newPtr,
		"{oldVersion}", *oldPtr,
		"{winnerLeg}", *winnerPtr,
		"{loserLeg}", *loserPtr,
	)

	resultFile := fmt.Sprintf("./%s_%s_%s(%s).json", *vertPtr, *envPtr, *newPtr, *winnerPtr)
	fmt.Println(fmt.Sprintf("Generating '%s'...\n", resultFile))

	result := replacer.Replace(string(template))
	err2 := ioutil.WriteFile(resultFile, []byte(result), 0644)
	check(err2)

	fmt.Println("Done! \\o/")
}
