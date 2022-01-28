package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/tealeg/xlsx"
)

var dirCurrent, err = os.Getwd()

type Header struct {
	EN string `json:"EN"`
	CN string `json:"CN"`
}

func errPrint(msg string) {
	fmt.Printf(msg)
}

func convert(excelFileName string) []*Header {
	excelFileName = strings.Replace(excelFileName, "\\", "/", -1)
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if r := recover(); r != nil {
			errPrint(fmt.Sprintf("%s 解析错误 panic: %v\n", excelFileName, r))
		}
	}()

	res := make([]*Header, 0)
	for _, sheet := range xlFile.Sheets {
		// if strings.Index(sheet.Name, "_") != 0 || len(sheet.Rows) == 0 {
		//     continue
		// }
		if len(sheet.Rows) == 0 {
			continue
		}
		rows := sheet.Rows

		for index, row := range rows {
			if index == 0 {
				continue
			}
			r := Header{}
			r.EN = row.Cells[0].String()
			r.CN = row.Cells[1].String()
			res = append(res, &r)
		}
	}
	// jsons, errs := json.Marshal(res)
	// if err != nil {
	// 	fmt.Println(errs.Error())
	// }
	// fmt.Println(string(jsons))
	return res

	// // 写json文件
	// filePtr, err := os.Create("data.json")
	// if err != nil {
	// 	fmt.Println("文件创建失败", err.Error())
	// 	return
	// }
	// defer filePtr.Close()

	// // 创建Json编码器
	// encoder := json.NewEncoder(filePtr)

	// err = encoder.Encode(res)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}

func walk(dir string) []*Header {
	fmt.Println(dir)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	header := make([]*Header, 0)
	for _, file := range files {
		filepath := path.Join(dir, file.Name())
		header = append(header, convert(filepath)...)
	}
	return header
}

func Converter() []*Header {
	// dirExcel := path.Join(dirCurrent, "data")
	dirExcel := dirCurrent + "\\data"
	fmt.Println(dirExcel)
	if info, err := os.Stat(dirExcel); !os.IsNotExist(err) && info.IsDir() {
		header := walk(dirExcel)
		return header
	}
	// fmt.Println("Press any key to exit...")
	// b := make([]byte, 1)
	// os.Stdin.Read(b)
	return nil
}
