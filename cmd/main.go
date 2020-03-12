package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/VinnieApps/test-aggregator/internal/xunit"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		base := filepath.Base(path)
		if strings.HasPrefix(base, "TEST-") {
			data, _ := ioutil.ReadFile(path)
			testSuite := &xunit.TestSuite{}

			if err := xml.Unmarshal([]byte(data), testSuite); err != nil {
				fmt.Printf("Error reading file: %s\n%v\n", path, err)
			} else {
				fmt.Printf("-- Test suite found: %s, Total Test Cases: %d\n", testSuite.Name, testSuite.TestCount)
				fmt.Printf("  Took: %f ms\n", testSuite.Time)
				fmt.Println("  Test cases:")
				for _, testCase := range testSuite.TestCases {
					fmt.Printf("    %s (%f ms)\n", testCase.Name, testCase.Time)
				}
			}
		}
		return nil
	})
}
