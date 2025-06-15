package firstPart

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getDirFiles(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	var res []string
	for _, e := range entries {
		res = append(res, dir+"/"+e.Name())
	}
	return res
}
func TestHello(t *testing.T) {
	inputFiles := getDirFiles("./input")
	outputFiles := getDirFiles("./output")

	for i := range inputFiles {
		start := time.Now()
		input, err := os.ReadFile(inputFiles[i])
		assert.Nil(t, err)
		if len(outputFiles) <= i {
			break
		}
		output, err := os.ReadFile(outputFiles[i])
		assert.Nil(t, err)

		instructions := Initialize(input)
		Execute(instructions)

		assert.Equal(t, computer.Output, string(output))
		fmt.Println("Execution time: ", inputFiles[i], time.Since(start))
	}
	fmt.Println("yo", inputFiles, outputFiles)
}
