package aocio

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"io/ioutil"
)

func panic_on_err(err error) {
	if err != nil {
		panic(err)
	}
}

func InputToIntArr(input string) []int {
	var arr = strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	var res = []int{}

	for _, i := range arr {
		j, err := strconv.Atoi(i)
		panic_on_err(err)
		res = append(res, j)
	}

	return res
}

// Like strconv.Atoi but returns a default value on error
func Atoi(s string, fallback int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return v
}

func ReadFile(file_path string) string {
	file, err := os.Open(file_path)
	panic_on_err(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := ioutil.ReadAll(reader)
	panic_on_err(err)

	return strings.TrimSuffix(string(contents), "\n")
}
