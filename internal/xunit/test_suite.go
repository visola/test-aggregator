package xunit

type TestSuite struct {
	ErrorCount   int        `xml:"errors,attr"`
	FailureCount int        `xml:"failures,attr"`
	Hostname     string     `xml:"hostname,attr"`
	Name         string     `xml:"name,attr"`
	SkipCount    int        `xml:"skipped,attr"`
	TestCount    int        `xml:"tests,attr"`
	TestCases    []TestCase `xml:"testcase"`
	Time         float32    `xml:"time,attr"`
}
