package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type testBlock struct {
	description string
	cases       []testCase
}

type testCase struct {
	in  string
	out string
}

func parseBlockDescription(testBlock string) string {
	re := regexp.MustCompile(`(?s)test\('(.*?)',`)
	match := re.FindStringSubmatch(testBlock)
	if len(match) < 2 {
		return ""
	}
	return match[1]
}

func splitTestCases(testBlock string) []string {
	matches := strings.Split(testBlock, "expect(")
	return matches
}

func parseTestCase(testBlock string) (string, string) {
	re := regexp.MustCompile(`(?s)twMerge\((.*?)\)\).toBe\((.*?)\)`)
	match := re.FindStringSubmatch(testBlock)
	if len(match) < 3 {
		return "", ""
	}
	in := strings.TrimSpace(match[1])
	in = strings.Trim(in, "'")
	in = strings.TrimSpace(in)
	in = strings.ReplaceAll(in, "\n", "")
	in = strings.ReplaceAll(in, `"`, "")

	out := strings.TrimSpace(match[2])
	out = strings.Trim(out, "'")
	out = strings.TrimSpace(out)
	out = strings.ReplaceAll(out, "\n", "")
	out = strings.ReplaceAll(out, `"`, "")
	return in, out
}

func splitBlocks(tests string) []string {
	re := regexp.MustCompile(`(?s)test\((.*?)}\)`)
	matches := re.FindAllString(tests, -1)
	return matches
}

func parse(tests string) []testBlock {
	if !strings.Contains(tests, "twMerge(") {
		return nil
	}

	testBlocks := splitBlocks(tests)

	parsedTests := make([]testBlock, 0)
	for _, block := range testBlocks {
		description := parseBlockDescription(block)
		if description == "" {
			continue
		}
		for _, test := range splitTestCases(block) {
			in, out := parseTestCase(test)
			if in == "" || out == "" {
				continue
			}
			parsedTests = append(parsedTests, testBlock{
				description: description,
				cases: []testCase{
					{in: in, out: out},
				},
			})
		}
	}
	return parsedTests
}

func genHtml(tests []testBlock) string {
	builder := strings.Builder{}
	for _, block := range tests {
		for _, test := range block.cases {
			builder.WriteString(fmt.Sprintf("<div class='%s'></div>\n", test.in))
		}
	}
	return builder.String()
}

func genTestCases(tests []testBlock) string {
	builder := strings.Builder{}
	s := `package rules

import "testing"

func TestTailwindMerge(t *testing.T) {
	tt := []struct {
		in string
		out string
	}{`
	builder.WriteString(s)
	lastDescription := ""
	for _, block := range tests {
		// write the description as a comment
		if lastDescription != block.description {
			builder.WriteString(fmt.Sprintf("\n	// %s\n", block.description))
			lastDescription = block.description
		}
		for _, test := range block.cases {
			builder.WriteString(fmt.Sprintf(`{
				in: "%s",
				out: "%s",
			},`, test.in, test.out))
		}
	}
	builder.WriteString("\n	}\n")
	builder.WriteString(`
	for _, tc := range tt {
		got := twMerge(tc.in)
		if got != tc.out {
			t.Errorf("twMerge returned %v, want %v", got, tc.out)
		}
	}
}`)
	return builder.String()
}

func main() {
	dir := `./tailwind-merge/tests`
	outHtml := `./rules/test_data/test.html`
	outGo := `./rules/twMerge_test.go`

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	tests := make([]testBlock, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), "test.ts") {
			continue
		}
		file, err := os.ReadFile(filepath.ToSlash(dir + "/" + entry.Name()))
		if err != nil {
			fmt.Println(err)
		}
		fileTests := parse(string(file))
		tests = append(tests, fileTests...)
	}
	html := genHtml(tests)
	goFile := genTestCases(tests)
	err = os.WriteFile(outHtml, []byte(html), 0644)
	if err != nil {
		fmt.Println(err)
	}
	err = os.WriteFile(outGo, []byte(goFile), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
