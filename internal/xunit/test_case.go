package xunit

type TestCase struct {
	Name string  `xml:"name,attr"`
	Time float32 `xml:"time,attr"`
}
