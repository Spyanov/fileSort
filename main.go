package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	src = "F:\\Исходная папка"
	dst = "F:\\Целевая папка\\"
)

func getDataList(directory string) []string {
	/*
		Принимаем директорию для поиска вложенных файлов
		Отдаем список файлов в директории
	*/
	var fileList []string

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fileList = append(fileList, f.Name())
	}
	return fileList
}

func folderList(list []string) []string {

	var dataList []string

	for _, item := range list {
		result := item[:7]

		fmt.Println("current list element[", item, "], result =", result)

		createFolder(result)

		err := os.Chdir(src)
		if err != nil {
			fmt.Println("Ошибка смены директории ")
		}
		in, err := os.Open(item)
		if err != nil {
			fmt.Println("Ошибка открытия файла для копирования")
		}
		defer in.Close()

		err = os.Chdir(dst + result)
		if err != nil {
			fmt.Println("Ошибка смены директории куда копировать", err)
		}
		out, err := os.Create(item)
		if err != nil {
			fmt.Println("Ошибка создания нового файла", err)
		}
		defer out.Close()

		_, err = io.Copy(out, in)
		if err != nil {
			fmt.Println("Ошибка копирования ", err)
		}
	}

	fmt.Println("datalist = ", dataList)

	return dataList
}

func createFolder(folderName string) {
	err := os.Chdir(dst)
	if err != nil {
		fmt.Println("Ошибка смены директории")
	}

	if _, err := os.Stat(folderName); err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(folderName, 0777)
			if err != nil {
				fmt.Println("Ошибка создания папки")
			}
		} else {
			fmt.Println("Папка существует")
		}
	}
}

func main() {

	list := getDataList(src)

	folders := folderList(list)
	for _, folder := range folders {
		fmt.Println("Folderlist= ", folder)
	}
}
