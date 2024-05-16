package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

func TestGetBlockType(t *testing.T) {
	tests := []struct {
		input        string
		expectedOpen string
	}{
		{"# this is a heading", "h1"},
		{"## this is a heading", "h2"},
		{"### this is a heading", "h3"},
		{"#### this is a heading", "h4"},
		{"##### this is a heading", "h5"},
		{"###### this is a heading", "h6"},
		{"1. this is an ordered list", "ol"},
		{"- this is an unordered list", "ul"},
		{"--- this is a line break", "br"},
	}
	fmt.Println("\nTest TestGetBlockType")
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.input)
		fmt.Println(testname)
	}
}

func TestSplitLine(t *testing.T) {
	tests := []struct {
		inputString   string
		expectedBlock string
		expectedBody  string
	}{
		{"# This has a header", "#", "This has a header"},
		{"- This has a bullet", "-", "This has a bullet"},
		{"This has a paragraph", "", "This has a paragraph"},
	}
	fmt.Println("\nTest TestSplitLine")
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			actualBlock, actualBody := SplitLine(tt.inputString)
			if actualBlock != tt.expectedBlock || actualBody != tt.expectedBody {
				t.Errorf("got .%s. .%s., wanted .%s. .%s.", actualBlock, actualBody, tt.expectedBlock, tt.expectedBody)
			}
		})
	}
}

func TestInnerText(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedString string
	}{
		{"This is a single line", "This is a single line"},
		{"This has **underlined** characters", "This has <strong>underlined</strong> characters"},
		{"This has *italic* characters", "This has <em>italic</em> characters"},
		{"This has **underlined** characters and *italic ones*", "This has <strong>underlined</strong> characters and <em>italic ones</em>"},
	}
	fmt.Println("\nTest TestInnerText")

	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := GetInnerText(tt.inputString)
			if ans != tt.expectedString {
				t.Errorf("got .%s., wanted .%s.", ans, tt.expectedString)
			}
		})
	}
}

func TestOpenTag(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedString string
	}{
		{"a", "<a>"},
	}
	fmt.Println("\nTest TestOpenTag")
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := OpenTag(tt.inputString)
			if ans != tt.expectedString {
				t.Errorf("got .%s., wanted .%s.", ans, tt.expectedString)
			}
		})
	}
}

func TestCloseTag(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedString string
	}{
		{"a", "</a>"},
	}
	fmt.Println("\nTest TestOpenTag")
	for _, tt := range tests {
		testname := fmt.Sprintf("input: %s", tt.inputString)
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			ans := CloseTag(tt.inputString)
			if ans != tt.expectedString {
				t.Errorf("got .%s., wanted .%s.", ans, tt.expectedString)
			}
		})
	}
}

type firstPassTest struct {
	caseName   string
	inputFile  []string
	outputFile []string
}

func createFirstPassCases() []firstPassTest {
	caseName := []string{
		"1. Single block line",
		"2. Single paragraph line",
		"3. Single ordered list line",
		"4. Multiple block line",
		"5. Multiple block line with italics/bold",
		"6. Multiple block and paragraph line",
		"7. Multiple block and ordered list",
		"8. Multiple block, paragraph, and ordered list with italics and bold",
		"9. Final test case, full document",
	}
	inputCases := [][]string{
		// 1
		{
			"# This is a single block line",
		},
		// 2
		{
			"This is a single paragraph line",
		},
		// 3
		{
			"1. This is the first element of a list",
			"1. This is the second element of a list",
		},
		// 4
		{
			"# This is the first line",
			"## This is the second line",
		},
		// 5
		{
			"# This is the *first* line",
			"## This is the **second** line",
		},
		// 6
		{
			"# This is the first line",
			"This is part of a paragraph",
		},
		// 7
		{
			"# This is the first line",
			"- This is part of a list",
			"- This is part of a list too",
		},
		// 8
		{
			"# This is the *first* line",
			"- This is part of a **list**",
			"This is part of a paragraph",
		},
		{
			"# Cookies or Ice Cream",
			"by *Tyler*",
			"## Introduction",
			"The topic of the  **best** dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.",
			"",
			"These factors include the following:",
			"- taste",
			"- temperature",
			"- price",
			"",
			"Only by considering each of these factors can we truly determine the ultimate dessert.",
			"",
			"### Conclusion",
			"The rankings for best desserts are as follows:",
			"1. Cookies",
			"1. Ice Cream",
			"Thank you",
		},
	}
	expectedOutputs := [][]string{
		{
			// 1
			"<html>",
			"<body>",
			"<h1>This is a single block line</h1>",
			"</body>",
			"</html>",
		}, {
			// 2
			"<html>",
			"<body>",
			"<p>This is a single paragraph line</p>",
			"</body>",
			"</html>",
		}, {
			// 3
			"<html>",
			"<body>",
			"<ol>",
			"\t<li>This is the first element of a list</li>",
			"\t<li>This is the second element of a list</li>",
			"</ol>",
			"</body>",
			"</html>",
		}, {
			// 4
			"<html>",
			"<body>",
			"<h1>This is the first line</h1>",
			"<h2>This is the second line</h2>",
			"</body>",
			"</html>",
		}, {
			// 5
			"<html>",
			"<body>",
			"<h1>This is the <em>first</em> line</h1>",
			"<h2>This is the <strong>second</strong> line</h2>",
			"</body>",
			"</html>",
		}, {
			// 6
			"<html>",
			"<body>",
			"<h1>This is the first line</h1>",
			"<p>This is part of a paragraph</p>",
			"</body>",
			"</html>",
		}, {
			// 7
			"<html>",
			"<body>",
			"<h1>This is the first line</h1>",
			"<ul>",
			"\t<li>This is part of a list</li>",
			"\t<li>This is part of a list too</li>",
			"</ul>",
			"</body>",
			"</html>",
		}, {
			// 8
			"<html>",
			"<body>",
			"<h1>This is the <em>first</em> line</h1>",
			"<ul>",
			"\t<li>This is part of a <strong>list</strong></li>",
			"</ul>",
			"<p>This is part of a paragraph</p>",
			"</body>",
			"</html>",
		}, {
			"<html>",
			"<body>",
			"<h1>Cookies or Ice Cream</h1>",
			"<p>by <em>Tyler</em></p>",
			"<h2>Introduction</h2>",
			"<p>The topic of the  <strong>best</strong> dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.</p>",
			"<p></p>",
			"<p>These factors include the following:</p>",
			"<ul>",
			"\t<li>taste</li>",
			"\t<li>temperature</li>",
			"\t<li>price</li>",
			"</ul>",
			"<p></p>",
			"<p>Only by considering each of these factors can we truly determine the ultimate dessert.</p>",
			"<p></p>",
			"<h3>Conclusion</h3>",
			"<p>The rankings for best desserts are as follows:</p>",
			"<ol>",
			"\t<li>Cookies</li>",
			"\t<li>Ice Cream</li>",
			"</ol>",
			"<p>Thank you</p>",
			"</body>",
			"</html>",
		},
	}
	tests := []firstPassTest{}
	for i := range caseName {
		tests = append(tests, firstPassTest{caseName: caseName[i], inputFile: inputCases[i], outputFile: expectedOutputs[i]})
	}
	return tests
}

func TestFirstPass(t *testing.T) {

	// Input and output are separate variables because the initial <html> tag will make the lines off

	tests := createFirstPassCases()

	fmt.Println("\nTest: TestFirstPass")
	for _, tt := range tests {
		isCorrect := true
		// load buffer
		var readBuffer, writerBuffer bytes.Buffer
		for _, line := range tt.inputFile {
			readBuffer.WriteString(string(line) + "\n")
		}

		testname := fmt.Sprintf("Case: %s", tt.caseName)
		// if tt.caseName == "9. Final test case, full document" {
		// 	continue
		// }
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			FirstPass(&readBuffer, &writerBuffer)
			checkOutput := bufio.NewScanner(&writerBuffer)
			for j, mock := range tt.outputFile {
				checkOutput.Scan()
				if string(mock) != checkOutput.Text() {
					t.Errorf("%v) got .%s., wanted .%s.", j, checkOutput.Text(), string(mock))
					isCorrect = false
					break
				}
			}
		})
		// TODO: FIGURE OUTo WHY THE BUFFER NEEDS TO BE READ AGAIN
		// Look into Tee reader
		if !isCorrect {
			for _, line := range tt.inputFile {
				readBuffer.WriteString(string(line) + "\n")
			}

			FirstPass(&readBuffer, &writerBuffer)
			firstPassOutput := bufio.NewScanner(&writerBuffer)
			fmt.Println("BEGIN FAILED OUTPUT")
			for firstPassOutput.Scan() {
				fmt.Println(firstPassOutput.Text())
			}
			isCorrect = true
			fmt.Println("END FAILED OUTPUT")
		}
	}

}

type secondPassTest struct {
	caseName   string
	inputFile  []string
	outputFile []string
}

func createSecondPassCases() []secondPassTest {
	caseName := []string{
		"1. No paragraphs at all",
		"2. No empty paragraphs(nonempty paragraps exist)",
		"3. Only empty paragraph",
		"4. Final test case, full document",
	}
	inputCases := [][]string{
		{
			// 1
			"<html>",
			"<body>",
			"<h1>This is a single block line</h1>",
			"</body>",
			"</html>",
		}, {
			// 2
			"<html>",
			"<body>",
			"<p>This is a single paragraph line</p>",
			"</body>",
			"</html>",
		}, {
			// 3
			"<html>",
			"<body>",
			"<p></p>",
			"</body>",
			"</html>",
		}, {
			// 4
			"<html>",
			"<body>",
			"<h1>Cookies or Ice Cream</h1>",
			"<p>by <em>Tyler</em></p>",
			"<h2>Introduction</h2>",
			"<p>The topic of the  <strong>best</strong> dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.</p>",
			"<p></p>",
			"<p>These factors include the following:</p>",
			"<ul>",
			"\t<li>taste</li>",
			"\t<li>temperature</li>",
			"\t<li>price</li>",
			"</ul>",
			"<p></p>",
			"<p>Only by considering each of these factors can we truly determine the ultimate dessert.</p>",
			"<p></p>",
			"<h3>Conclusion</h3>",
			"<p>The rankings for best desserts are as follows:</p>",
			"<ol>",
			"\t<li>Cookies</li>",
			"\t<li>Ice Cream</li>",
			"</ol>",
			"<p>Thank you</p>",
			"</body>",
			"</html>",
		},
	}

	expectedOutputs := [][]string{
		{
			// 1
			"<html>",
			"\t<body>",
			"\t\t<h1>This is a single block line</h1>",
			"\t</body>",
			"</html>",
		}, {
			// 2
			"<html>",
			"\t<body>",
			"\t\t<p>This is a single paragraph line</p>",
			"\t</body>",
			"</html>",
		}, {
			// 3
			"<html>",
			"\t<body>",
			"\t</body>",
			"</html>",
		}, {
			// 4
			"<html>",
			"\t<body>",
			"\t\t<h1>Cookies or Ice Cream</h1>",
			"\t\t<p>by <em>Tyler</em></p>",
			"\t\t<h2>Introduction</h2>",
			"\t\t<p>The topic of the  <strong>best</strong> dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.</p>",
			"\t\t<p>These factors include the following:</p>",
			"\t\t<ul>",
			"\t\t\t<li>taste</li>",
			"\t\t\t<li>temperature</li>",
			"\t\t\t<li>price</li>",
			"\t\t</ul>",
			"\t\t<p>Only by considering each of these factors can we truly determine the ultimate dessert.</p>",
			"\t\t<h3>Conclusion</h3>",
			"\t\t<p>The rankings for best desserts are as follows:</p>",
			"\t\t<ol>",
			"\t\t\t<li>Cookies</li>",
			"\t\t\t<li>Ice Cream</li>",
			"\t\t</ol>",
			"\t\t<p>Thank you</p>",
			"\t</body>",
			"</html>",
		},
	}
	tests := []secondPassTest{}
	for i := range caseName {
		tests = append(tests, secondPassTest{caseName: caseName[i], inputFile: inputCases[i], outputFile: expectedOutputs[i]})
	}
	return tests
}

func TestSecondPass(t *testing.T) {

	// Input and output are separate variables because the initial <html> tag will make the lines off

	tests := createSecondPassCases()

	fmt.Println("\nTest: TestSecondPass")
	for _, tt := range tests {
		isCorrect := true
		// load buffer
		var readBuffer, writerBuffer bytes.Buffer
		for _, line := range tt.inputFile {
			readBuffer.WriteString(string(line) + "\n")
		}

		testname := fmt.Sprintf("Case: %s", tt.caseName)
		// if tt.caseName == "9. Final test case, full document" {
		// 	continue
		// }
		fmt.Println(testname)
		t.Run(testname, func(t *testing.T) {
			SecondPass(&readBuffer, &writerBuffer)
			checkOutput := bufio.NewScanner(&writerBuffer)
			for j, mock := range tt.outputFile {
				checkOutput.Scan()
				if string(mock) != checkOutput.Text() {
					t.Errorf("%v) got .%s., wanted .%s.", j, checkOutput.Text(), string(mock))
					isCorrect = false
					break
				}
			}
		})
		// TODO: FIGURE OUTo WHY THE BUFFER NEEDS TO BE READ AGAIN
		// Look into Tee reader
		if !isCorrect {
			fmt.Println("START DEBUG")
			var rb2, wb2 bytes.Buffer
			for _, line := range tt.inputFile {
				rb2.WriteString(string(line) + "\n")
			}

			SecondPass(&rb2, &wb2)
			firstPassOutput := bufio.NewScanner(&wb2)
			fmt.Println("BEGIN FAILED OUTPUT")
			for firstPassOutput.Scan() {
				fmt.Println(firstPassOutput.Text())
			}
			isCorrect = true
			fmt.Println("END FAILED OUTPUT")
			fmt.Println("End DEBUG")
		}
	}

}

type convertTest struct {
	caseName   string
	inputFile  []string
	outputFile []string
}

func createConvertCase() convertTest {
	caseName := "Only case"
	inputCase := []string{
		"# Cookies or Ice Cream",
		"by *Tyler*",
		"## Introduction",
		"The topic of the  **best** dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.",
		"",
		"These factors include the following:",
		"- taste",
		"- temperature",
		"- price",
		"",
		"Only by considering each of these factors can we truly determine the ultimate dessert.",
		"",
		"### Conclusion",
		"The rankings for best desserts are as follows:",
		"1. Cookies",
		"1. Ice Cream",
		"Thank you",
	}

	expectedOutput := []string{
		"<html>",
		"\t<body>",
		"\t\t<h1>Cookies or Ice Cream</h1>",
		"\t\t<p>by <em>Tyler</em></p>",
		"\t\t<h2>Introduction</h2>",
		"\t\t<p>The topic of the  <strong>best</strong> dessert food is one that has never had a definitive answer. This is often because there are too many factors of taste that must be accounted for.</p>",
		"\t\t<p>These factors include the following:</p>",
		"\t\t<ul>",
		"\t\t\t<li>taste</li>",
		"\t\t\t<li>temperature</li>",
		"\t\t\t<li>price</li>",
		"\t\t</ul>",
		"\t\t<p>Only by considering each of these factors can we truly determine the ultimate dessert.</p>",
		"\t\t<h3>Conclusion</h3>",
		"\t\t<p>The rankings for best desserts are as follows:</p>",
		"\t\t<ol>",
		"\t\t\t<li>Cookies</li>",
		"\t\t\t<li>Ice Cream</li>",
		"\t\t</ol>",
		"\t\t<p>Thank you</p>",
		"\t</body>",
		"</html>",
	}
	test := convertTest{caseName: caseName, inputFile: inputCase, outputFile: expectedOutput}
	return test
}

func TestConvert(t *testing.T) {

	// Input and output are separate variables because the initial <html> tag will make the lines off

	test := createConvertCase()

	fmt.Println("\nTest: TestConvert")

	isCorrect := true
	// load buffer
	var readBuffer, writerBuffer bytes.Buffer
	for _, line := range test.inputFile {
		readBuffer.WriteString(string(line) + "\n")
	}

	testname := fmt.Sprintf("Case: %s", test.caseName)

	fmt.Println(testname)
	t.Run(testname, func(t *testing.T) {
		Convert(&readBuffer, &writerBuffer)
		checkOutput := bufio.NewScanner(&writerBuffer)
		for j, mock := range test.outputFile {
			checkOutput.Scan()
			if string(mock) != checkOutput.Text() {
				t.Errorf("%v) got .%s., wanted .%s.", j, checkOutput.Text(), string(mock))
				isCorrect = false
				break
			}
		}
	})
	// TODO: FIGURE OUTo WHY THE BUFFER NEEDS TO BE READ AGAIN
	// Look into Tee reader
	if !isCorrect {
		for _, line := range test.inputFile {
			readBuffer.WriteString(string(line) + "\n")
		}

		FirstPass(&readBuffer, &writerBuffer)
		firstPassOutput := bufio.NewScanner(&writerBuffer)
		fmt.Println("BEGIN FAILED OUTPUT")
		for firstPassOutput.Scan() {
			fmt.Println(firstPassOutput.Text())
		}
		isCorrect = true
		fmt.Println("END FAILED OUTPUT")
	}

}
