package aocio 

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

func panic_on_err(err error) {
	if err != nil {
		panic(err)
	}
}

// Like strconv.Atoi but returns a default value on error
func atoi(s string, fallback int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return fallback
	}
	return v
}

func read_file(file_path string) string {
	file, err := os.Open(file_path)
	panic_on_err(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	contents, err := ioutil.ReadAll(reader)
	panic_on_err(err)

	return strings.TrimSuffix(string(contents), "\n")
}
