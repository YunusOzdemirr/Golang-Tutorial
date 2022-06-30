package main

import (
	"errors"
	"log"
	"os"
)

var (
	newFile  *os.File
	err      error
	fileInfo *os.FileInfo
)

func main2() {
	// newFile, err := os.Create("demo.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(newFile)

	// fileInfo, err := os.Stat("demo.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Dosya Adi: ", fileInfo.Name())
	// fmt.Println("Dosya Boyutu: ", fileInfo.Size())
	// fmt.Println("Dosya İzinleri: ", fileInfo.Mode())

	// _, err := RenameFile("demo.txt", "./Zort/asd.txt")
	// fmt.Println(err)
	//ChannelTest()
	// result, err := IsExist("./Zort/a2sd.txt")
	// fmt.Println(result, err)
}
func IsExist(path string) (string, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", errors.New("File does not exist")
		}
	}
	return fileInfo.Name(), nil
}

func RenameFile(oldPath, newPath string) (string, error) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal("Failed to rename file")
		return "Fail", err
	}
	return "Success", nil
}
