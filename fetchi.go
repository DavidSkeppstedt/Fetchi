package main

import (
	"github.com/DavidSkeppstedt/fetchi/facebook"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var conf = readConfigfile(os.Args[1])
	facebook.TestUsers(conf[0], conf[0]+"|"+conf[1]).FormatPrint()

}

func readConfigfile(path string) []string {
	credentials, ferr := ioutil.ReadFile(path)
	if ferr != nil {
		panic("Can not open and read the file")
	}
	config := strings.Split(string(credentials), "\n")
	return config
}
