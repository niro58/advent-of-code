package main

import (
	"fmt"
	"os"
	"regexp"
	"testing"
)
func TestFirstPart(t *testing.T) {
    input,err := os.ReadFile("inputs\\input_first_part.txt")
    if err != nil {
        fmt.Println(err)
    }
    output,err := os.ReadFile("outputs\\output_first_part.txt")
    if err != nil{
        fmt.Println(err)
    }

    want := regexp.MustCompile(`\b`+string(output)+`\b`)

    msg := fmt.Sprint(calculateFirstPart(string(input)))

    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`calculate(%q) = %q, %v, want match for %#q, nil`, string(input), msg, err, want)
    }
}
func TestSecondPart(t *testing.T) {
    input,err := os.ReadFile("inputs\\input_second_part.txt")
    if err != nil {
        fmt.Println(err)
    }
    output,err := os.ReadFile("outputs\\output_second_part.txt")
    if err != nil{
        fmt.Println(err)
    }

    want := regexp.MustCompile(`\b`+string(output)+`\b`)

    msg := fmt.Sprint(calculateFirstPart(string(input)))

    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`calculate(%q) = %q, %v, want match for %#q, nil`, string(input), msg, err, want)
    }
}