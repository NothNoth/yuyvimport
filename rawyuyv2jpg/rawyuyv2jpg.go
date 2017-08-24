package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"

	"strconv"

	"github.com/NothNoth/yuyvimport"
)

func help() {
	fmt.Println("Usage: " + os.Args[0] + " <width> <height> <raw yuyv file(s)>")
}

func main() {
	if len(os.Args) < 4 {
		help()
		return
	}

	w, err := strconv.Atoi(os.Args[1])
	if err != nil {
		help()
		return
	}
	h, err := strconv.Atoi(os.Args[2])
	if err != nil {
		help()
		return
	}
	for i := 3; i < len(os.Args); i++ {
		data, err := ioutil.ReadFile(os.Args[i])
		if err != nil {
			fmt.Println("Failed to read" + os.Args[i] + err.Error())
			continue
		}

		img := yuyvimport.Import(w, h, data)
		jpgName := fmt.Sprintf("%s_converted.jpg", os.Args[i])
		f, _ := os.Create(jpgName)
		defer f.Close()
		err = jpeg.Encode(f, img, nil)
		if err != nil {
			fmt.Println("Error while encoding " + jpgName + err.Error())
		} else {
			fmt.Println("Created " + jpgName)
		}
	}
}
