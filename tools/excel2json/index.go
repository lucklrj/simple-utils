package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
	stringsHelper "github.com/lucklrj/simple-utils/helper/strings"
	systemHelper "github.com/lucklrj/simple-utils/helper/system"
	"github.com/xuri/excelize/v2"
	_ "gopkg.in/ffmt.v1"
)

type SingleTitle struct {
	Title string
	Desc  string
}

func main() {

	flag.Parse()
	sp := strings.Repeat("-", 100)

	binDir, err := systemHelper.GetBinDir()
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}
	//遍历xlsx文件
	var xlsxFiles []string
	err = filepath.Walk(binDir, func(path string, info os.FileInfo, err error) error {
		if path[len(path)-5:] == ".xlsx" {
			xlsxFiles = append(xlsxFiles, path)
		}
		return nil
	})
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}

	for _, path := range xlsxFiles {
		color.Green(path)
		allTitle, allBody, err := ParseExcel(path)
		if err != nil {
			color.Red(err.Error())
			color.Red(sp)
			continue
		}
		err = Verification(allTitle, allBody)
		if err != nil {
			color.Red(err.Error())
			color.Red(sp)
			continue
		}

		err = WriteFile(path, allBody)
		if err != nil {
			color.Red(err.Error())
			color.Red(sp)
			continue
		}
		color.Green(sp)
	}
	color.Green("处理完毕")
	fmt.Scanln()
}

func ParseExcel(path string) (map[int]SingleTitle, []map[string]string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return nil, nil, err
	}
	//读取所有行
	rows, err := f.GetRows(f.GetSheetList()[0])
	if err != nil {
		return nil, nil, err
	}
	if len(rows) == 0 {
		return nil, nil, errors.New("empty")
	}

	//遍历标题
	allTitle := map[int]SingleTitle{}
	for index, content := range rows[0] {
		if strings.Contains(content, "|") == false {
			errMsg := fmt.Sprintf("line:%d,%s lost |", 1, content)
			return nil, nil, errors.New(errMsg)
		} else {
			singleTitle := strings.Split(content, "|")
			title := titleFormat(singleTitle[0])
			if title == "" {
				errMsg := fmt.Sprintf("line:%d,%s lost key", 1, content)
				return nil, nil, errors.New(errMsg)
			}
			allTitle[index] = SingleTitle{
				title,
				singleTitle[1],
			}
		}
	}
	//解析内容
	allBody := make([]map[string]string, 0)

	for line, row := range rows {
		if line == 0 {
			continue
		}
		body := map[string]string{}
		for index, content := range row {
			body[allTitle[index].Title] = content
		}
		allBody = append(allBody, body)
	}

	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			color.Red(err.Error())
		}
	}()
	return allTitle, allBody, nil
}
func titleFormat(title string) string {
	reg, _ := regexp.Compile(`\s+`)
	title = reg.ReplaceAllString(title, "_")
	return stringsHelper.UnderscoreToCamelCase(title, true)
}
func Verification(allTitle map[int]SingleTitle, allBody []map[string]string) error {
	//检查id是否重复
	isHaveTitle := false
	for _, singleTitle := range allTitle {
		if singleTitle.Title == "id" {
			isHaveTitle = true
			break
		}
	}
	if isHaveTitle {
		var ids = make(map[string]bool)
		for index, line := range allBody {
			id := strings.TrimSpace(line["id"])
			if id == "" {
				errMsg := fmt.Sprintf("line:%d lost id", index+2)
				return errors.New(errMsg)
			}
			_, ok := ids[id]
			if ok {
				errMsg := fmt.Sprintf("line:%d id exists", index+2)
				return errors.New(errMsg)
			} else {
				ids[id] = true
			}
		}
	}
	return nil
}
func WriteFile(path string, allBody []map[string]string) error {
	targetPath := path[0:len(path)-5] + ".json"

	content, err := json.MarshalIndent(allBody, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(targetPath, content, 0644)
}
