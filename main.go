package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func getDataList(directory string) []string {
	var list []string

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		list = append(list, f.Name())
		//fmt.Println(f.Name())

	}
	return list
}

func folderList(list []string) []string {

	//fmt.Println("list = ", list)

	var dataList []string

	for _, item := range list {
		result := item[:7]

		fmt.Println("current list element[", item, "], result =", result)

		createFolder(result)

		err := os.Chdir(`F:\ЯндексДиск.Безлимит\`)
		if err != nil {
			fmt.Println("Ошибка смены директории ")
		}
		in, err := os.Open(item)
		if err != nil {
			fmt.Println("Ошибка открытия файла для копирования")
		}
		defer in.Close()

		err = os.Chdir(`f:\photo\` + result)
		if err != nil {
			fmt.Println("Ошибка смены директории куда копироват")
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

func getFileList() {

}

func createFolder(folderName string) {

	err := os.Chdir(`f:\photo\`)
	if err != nil {
		fmt.Println("Ошибка смены директории")
	}
	err = os.Mkdir(folderName, 0777)
	if err != nil {
		fmt.Println("Ошибка создания папки")
	}

}

func main() {
	fmt.Println("run")

	list := getDataList("F:\\ЯндексДиск.Безлимит")

	folders := folderList(list)
	for _, folder := range folders {
		fmt.Println("Folderlist= ", folder)

	}

	//createFolder("07.07.2020")

}
