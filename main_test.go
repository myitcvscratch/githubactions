package main

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestTest(t *testing.T) {
	fmt.Println(filepath.Join("hello", "*.txt"))
}
