package __init__

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
)

var _ = BeforeEach(
	func() {
		currentTest := CurrentSpecReport().LeafNodeText
		fmt.Println("Before the first test", currentTest)
		if currentTest == "should correctly add two numbers" {
			Skip("Skipping this test")
		}

	})

var _ = BeforeEach(
	func() {
		fmt.Println("Before each test second time")

	})
