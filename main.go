package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/tidwall/gjson"
	"github.com/xuri/excelize/v2"
	"go_study/libs"
	"gopkg.in/yaml.v3"
	"io/fs"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"time"
)

func main() {
	d := []byte(`{"name":"aaa","gender":1, "address": {"country":"中a国","cityp":"used"},"email":["aa","bb",3,4,5]}`)
	var j map[string]any
	err := json.Unmarshal(d, &j)
	if err != nil {
		return
	}
	var gj string = gjson.GetBytes(d, "address.city").String()
	if gj=="" {
		fmt.Println("empty")
	}

	readFile()
	writeFile()
	readFileLine()
	writeFileLine()
	readYaml()
	writeYaml()
	readXlsx()
	writeXlsx()
	pathOperation()
	studyDatetime()
	studyConcat()
	studyIpNetwork()
	studyStr2num()
	user := &study.User{Name: "赵钱孙李", Age: 32}
	fmt.Println(user.MyAge())
	var a study.Action = user
	a.SetAge(45)
	fmt.Println(user.MyAge())
}

func readFile() {
	file, err := os.ReadFile(`D:\OneDrive\python\tool.py`)
	if err != nil {
		return
	}

	fmt.Println(string(file))
}

func readFileLine() {
	file, err := os.Open(`D:\OneDrive\python\tool.py`)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		fmt.Println(reader.Text())
	}

}

func writeFile() {
	file, err := os.Create("data.txt")
	if err != nil {
		return
	}
	defer file.Close()

	_, _ = file.Write([]byte("大口大口的"))
}

func writeFileLine() {
	file, err := os.Create("data.txt")
	if err != nil {
		return
	}
	defer file.Close()

	for i := 1; i <= 10; i++ {
		_, _ = file.Write([]byte("大口大口的\n"))
	}
}

func readYaml() {
	file, err := os.ReadFile(`C:\Users\sharp\AppData\Local\Programs\clash_win\config.yaml`)
	if err != nil {
		return
	}
	var data map[string]any
	err = yaml.Unmarshal(file, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data["dns"].(map[string]any)["nameserver"])
}

func writeYaml() {
	file, err := os.Create("data.txt")
	if err != nil {
		return
	}
	defer file.Close()

	fileIn, _ := os.ReadFile(`C:\Users\sharp\AppData\Local\Programs\clash_win\config.yaml`)
	var data map[string]any
	if err = yaml.Unmarshal(fileIn, &data); err != nil {
		fmt.Println(err)
	}

	out, _ := yaml.Marshal(data)
	if _, err = file.Write(out); err != nil {
		fmt.Println(err)
	}
}

func readXlsx() {
	f, err := excelize.OpenFile(`C:\Users\sharp\Desktop\data\2023-04-21-plan2-all-f11.xlsx`)
	if err != nil {
		return
	}
	defer func(f *excelize.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	sheet, _ := f.GetRows("全国")
	for _, row := range sheet {
		fmt.Println(row)
	}
}

func writeXlsx() {
	file := excelize.NewFile()
	defer func(file *excelize.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	_, _ = file.NewSheet("sheet1")

	err := file.SetCellValue("sheet1", "A1", 509)
	if err != nil {
		return
	}
	if err = file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func pathOperation() {
	_ = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path, d.IsDir())
		return nil
	})

	files, _ := filepath.Glob("*.go")
	fmt.Println(files)
	if _, err := os.Stat("/path/to/whatever"); errors.Is(err, os.ErrNotExist) {
		fmt.Println(false)
	}
}

func studyConcat() {
	a := "sss"
	b := "ddd"
	fmt.Println(a + b)
}

func studyIpNetwork() {
	net.ParseIP("1.2.3.4")

}

func studyStr2num() {
	ia := 10
	var iaaa int64
	ib := "50000000000000000000"

	fa := 100.4
	fb := "50.6"

	iaa := strconv.Itoa(ia)
	ibb, _ := strconv.Atoi(ib)
	ibb2, _ := strconv.ParseInt(ib, 10, 64)

	faa := strconv.FormatFloat(fa, 'f', -1, 32)
	fbb, _ := strconv.ParseFloat(fb, 32)
	fbb2 := float32(fbb)

	fmt.Println(new(big.Int).SetString(ib, 10))
	fmt.Println(iaa, reflect.TypeOf(iaa))
	fmt.Println(ibb, reflect.TypeOf(ibb))
	fmt.Println(ibb2, reflect.TypeOf(ibb2))
	fmt.Println(faa, reflect.TypeOf(faa))
	fmt.Println(fbb, reflect.TypeOf(fbb))
	fmt.Println(fbb2, reflect.TypeOf(fbb2))
	fmt.Println(iaaa)
}

func studyDatetime() {
	now := time.Now()
	utc := time.Now().UTC()
	tz, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().Unix())
	fmt.Println(now.Format(time.DateTime))
	fmt.Println(utc.In(tz).Add(-2 * time.Hour))
	fmt.Println(time.Unix(1683275206, 0))
	var dt1 = time.Date(2013, 12, 20, 0, 0, 0, 0, tz)
	var dt2 = time.Date(2014, 1, 20, 0, 0, 0, 0, tz)
	for dt := dt1; dt.Before(dt2); dt = dt.Add(time.Hour * 24) {
		fmt.Println(dt)
	}

	var nowc = carbon.CreateFromStdTime(now)
	// 时区转换
	fmt.Println(nowc.SetTimezone("utc"), "xxxxxxxx")
	// 修改日期
	fmt.Println(nowc.SetDay(25), "xxxxxxxx")
	fmt.Println(nowc.SubDays(30).Date())
	fmt.Println(carbon.CreateFromTimestamp(1683275206).ToDateTimeString())
	var dtc1 = carbon.CreateFromDate(2021, 12, 20)
	for dt := dtc1; dt.Lte(dtc1.AddDays(30)); dt = dt.AddDay() {
		fmt.Println(dt.ToIso8601String())
		fmt.Println(dt.Format("Y-m-d H:i:s P"))
	}
}
