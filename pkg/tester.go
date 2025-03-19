package tester

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"unicode"
)

type Tester struct {
	Path  string
	tasks []*Task
}

type Task struct {
	TestFile string
	In       string
	Out      string
}

func New(path string) *Tester {
	return &Tester{Path: filepath.Join(path, "testdata")}
}

func (t *Tester) GetTasks() (tasks []*Task) {
	nr := 0
	for {
		testFile := fmt.Sprintf("test.%s", strconv.Itoa(nr))
		fileIn := fmt.Sprintf("%s.%s", testFile, "in")
		fileOut := fmt.Sprintf("%s.%s", testFile, "out")

		if !t.fileExists(fileIn) || !t.fileExists(fileOut) {
			break
		}

		task := &Task{
			TestFile: testFile,
			In:       t.getFileContent(fileIn),
			Out:      t.getFileContent(fileOut),
		}
		tasks = append(tasks, task)
		nr++
	}
	return
}

func (t *Tester) getFileContent(filename string) string {
	fromFileContent, _ := os.ReadFile(filepath.Join(t.Path, filename))
	formatted := ""
	for _, r := range string(fromFileContent) {
		if !unicode.IsSpace(r) {
			formatted += string(r)
		}
	}
	return formatted
}

func (t *Tester) fileExists(filename string) bool {
	info, err := os.Stat(filepath.Join(t.Path, filename))
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
