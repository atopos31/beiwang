package main

import (
	"fmt"
	//"go/build"
	"os"
	"strings"

	// "sync"
	"encoding/csv"
	"io/ioutil"
	"strconv"
)

func main() {
	fmt.Println("-----------------欢迎来到hackerxiao的借贷备忘录程序-----------------\n---------->版本 V1.0\n访问我的个人网站----点击 hackerxiao.online\n\n")
	for true {
		var i int
		fmt.Printf("请输入1读取记录，输入2写入记录,输入3新建对象，输入4退出程序\n")
		fmt.Scan(&i)

		if i == 1 {
			readCSV()
		} else if i == 2 {
			writerCSV()
		} else if i == 3 {
			buildCSV()
		} else if i == 4 {
			break
		}
		fmt.Println("\n")

	}

}

func readCSV() int {
	fmt.Printf("请输入文件名查询记录，输入0退出\n") 
	var nam string
	fmt.Scan(&nam)
	e := ".csv"
	name := nam + e
	if name == "0" {
		return 0
	}
	f, _ := os.Open(name)
	defer f.Close()
	content, err := ioutil.ReadAll(f)

	cntb, err := newFunction(name)
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)

	var n [1000]string
	var p [1000]int

	for i := 0; i < sz; i++ {
		n[i] = ss[i][0]
		p[i], _ = strconv.Atoi(ss[i][1])
	}

	num := 0
	for i := 0; i < sz-1; i++ {
		if n[i+1] == "+" {
			num += p[i+1]
		} else {
			num -= p[i+1]
		}
	}

	if nil != err {
		fmt.Println("读取", name, "失败！")

	}

	fmt.Println(string(content))
	fmt.Println("余额", num, "元")
	fmt.Println("查询完毕")

	return 0
}

func newFunction(name string) ([]byte, error) {
	cntb, err := ioutil.ReadFile(name)
	return cntb, err
}

func writerCSV() int {
	fmt.Printf("请输入文件名写入记录，输入0退出\n")
	var nam string
	fmt.Scan(&nam)
	e := ".csv"
	name := nam + e
	if name == "0" {
		return 0
	}

	File, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败！")
	}
	defer File.Close()

	fmt.Printf("请输入要写入的数据--------格式为 +/-(代表还/借)(空格)(金额）\n")

	var slice []string
	for i := 0; i < 2; i++ {
		var element string
		fmt.Scan(&element)
		slice = append(slice, element)

	}

	//创建写入接口
	WriterCsv := csv.NewWriter(File)

	//写入一条数据，传入数据为切片(追加模式)
	err1 := WriterCsv.Write(slice)
	if err1 != nil {
		fmt.Println("写入文件失败")
	}
	WriterCsv.Flush() //刷新一下，不刷新无法写入的
	fmt.Println("数据写入成功.......")

	//下面是判断余额为0的时候是否想要删除记录
	f, _ := os.Open(name)
	defer f.Close()

	cntb, err := newFunction(name)
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()  
	sz := len(ss)

	var n [1000]string
	var p [1000]int

	for i := 0; i < sz; i++ {
		n[i] = ss[i][0]
		p[i], _ = strconv.Atoi(ss[i][1])
	}

	num := 0
	for i := 0; i < sz-1; i++ {
		if n[i+1] == "+" {
			num += p[i+1]
		} else {
			num -= p[i+1]
		}
	}

	if num==0 {
		fmt.Println(name,"的借款已还清，是否删除借款记录？\n","输入yes删除")

		var stz string
	    fmt.Scan(&stz)

		if stz=="yes"{
			writeContent := "on lan,balance\n"
			err = ioutil.WriteFile(name, []byte(writeContent), os.ModePerm)
			//此处并不是清楚的操作，而是用一个先清空再写入的函数，实现效果。
			if nil != err {
    		fmt.Println(name,"删除记录失败")
			}

			fmt.Println(name,"的记录已删除")

		}
	}
	

	return 0
}

func buildCSV() int {
	fmt.Printf("请输入文件名以创建新的csv文件，输入0退出\n")

	var nam string
	fmt.Scan(&nam)
	e := ".csv"
	name := nam + e
	if name == "0" {
		return 0
	}

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")  //写入UTF-8 防止word打开乱码
	f.WriteString("on lan,balance\n")

	fmt.Printf("文件创建成功\n")
	return 0
}
