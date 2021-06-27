package main

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error) {
	//Open file
	f, err := os.Open(filename)
	//Handle the error of there is any
	if err != nil {
		return 0, err
	}
	//Close the file
	defer f.Close()

	//create a hasher
	h := crc32.NewIEEE()
	_, err := io.Copy(h, f)
	//Handle the error if there is any
	if err != nil {
		return 0, nil
	}

	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("text1.txt")

	if err != nil {
		return
	}

	h2, err := getHash("text2.txt")

	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}