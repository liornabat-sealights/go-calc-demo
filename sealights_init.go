package main

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/ginkgo/v2/types"
	"log"
	"time"
)

type TestResult struct {
	Name           string
	State          string
	Duration       time.Duration
	FailureMessage string
	Location       string
}

var _ = BeforeEach(
	func() {
		currentTest := CurrentSpecReport().LeafNodeText
		log.Printf("Sealights hook BeforeEach checking if test %s need to be skiped\n", currentTest)
	})

var _ = AfterEach(
	func() {
		log.Println("Sealights hook AfterEach checking the test result")
		testReport := CurrentSpecReport()
		if testReport.State.Is(types.SpecStatePassed) {
			log.Printf("Test %s passed\n", testReport.LeafNodeText)
		}
		if testReport.State.Is(types.SpecStateFailed) {
			log.Printf("Test %s failed\n", testReport.LeafNodeText)

		}
		if testReport.State.Is(types.SpecStateSkipped) {
			log.Printf("Test %s skipped\n", testReport.LeafNodeText)
		}

	})

func init() {
	log.Println("This is the start of Sealights hook init, we are adding a place holder for calling the Sealights backend to retrieve the test data")
	time.Sleep(2 * time.Second)
}
