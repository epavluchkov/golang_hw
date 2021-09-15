package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment, len(files))

	for _, fileInfo := range files {
		if fileInfo.Size() == 0 {
			env[fileInfo.Name()] = EnvValue{"", true}
		} else {
			file, err := os.Open(dir + "/" + fileInfo.Name())
			if err != nil {
				return nil, err
			}

			scan := bufio.NewScanner(file)
			scan.Scan()
			val := scan.Text()

			if err := scan.Err(); err != nil {
				return nil, err
			}

			val = strings.TrimRight(val, " \t")
			val = string(bytes.Replace([]byte(val), []byte(string(rune(0x00))), []byte("\n"), 1))

			env[fileInfo.Name()] = EnvValue{val, false}

			file.Close()
		}
	}

	return env, nil
}
