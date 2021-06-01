package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetRootDomain(fqdn string) (str string) {
	arr := strings.Split(fqdn, ".")

	max := len(arr)

	str = fmt.Sprintf("%s.%s", arr[max-2], arr[max-1])

	return str
}

func ReadFile(path string) string {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return ""
	}
	data, err := ioutil.ReadFile(path)
	CheckErr(err)

	return string(data)
}

func WriteFile(path string, content string) {
	err := ioutil.WriteFile(path, []byte(content), 0700)
	CheckErr(err)
}

func CreateFolderIfNotExists(path string, mode fs.FileMode) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return os.Mkdir(path, mode)
	}
	return nil
}
