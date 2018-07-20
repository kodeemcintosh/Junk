package main

import(
	"os"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Expect_WhenFileOrDirectoryNotInCurrentDirectory_ShouldReturnFalse(t *testing.T) {

	testDirectory := "/Users/1kw8742i76k/Repos/GoSpace/src/github.com/kvmac/Junk/file.txt"
	output, doesExist := main.Exists(directory)
	Assert.Equals(t, false, doesExist, )
}

func Expect_WhenFileOrDirectoryNotInCurrentDirectory_ShouldReturnTrue(t *testing.T) {

	testDirectory := "/Users/1kw8742i76k/Repos/GoSpace/src/github.com/kvmac/Junk/testFile.txt"
	output, doesExist := main.Exists(directory)
	Assert.Equals(t, true, doesExist)
}
