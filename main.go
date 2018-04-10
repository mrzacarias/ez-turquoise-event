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
	vertPtr = flag.String("vert", "home", "Vertical name (auto|home|life)")
	envPtr = flag.String("env", "staging", "Vertical environment (staging|prod)")
	newPtr = flag.String("new", "REQUIRED", "New version (will be scaled up)")
	oldPtr = flag.String("old", "REQUIRED", "Old version (will be scaled down)")
	winnerPtr = flag.String("winner", "NO WINNER SET", "Winner leg (100% A/B)")
	loserPtr = flag.String("loser", "NO LOSER SET", "Loser leg (100% A/B)")
	percPtr = flag.String("percentage", "100", "Percentage of instances with new version (10|33|50|100)")

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

	template, err := ioutil.ReadFile("./internal/templates/event_template.json")
	check(err)
	fmt.Print("Template loaded...\n\n")

	fmt.Printf("Flags ok:\n"+
		"\tVertical: %s\n"+
		"\tEnvironment: %s\n"+
		"\tNew Version: %s\n"+
		"\tOld Version: %s\n"+
		"\tWinner Leg: %s\n"+
		"\tLoser Leg: %s\n"+
		"\tPercentage: %s\n\n",
		*vertPtr, *envPtr, *newPtr, *oldPtr, *winnerPtr, *loserPtr, *percPtr)

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
		"{percentage}", *percPtr,
	)

	resultFile := fmt.Sprintf("./%s_%s_%s(%s)_%s.json", *vertPtr, *envPtr, *newPtr, *winnerPtr, *percPtr)
	fmt.Println(fmt.Sprintf("Generating '%s'...\n", resultFile))

	result := replacer.Replace(string(template))
	err2 := ioutil.WriteFile(resultFile, []byte(result), 0644)
	check(err2)

	fmt.Println("Done! \\o/")
}
