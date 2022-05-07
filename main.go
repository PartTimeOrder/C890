package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
	"strings"
	"time"
)

func init() {
	_, err := os.Stat("img")
	if os.IsNotExist(err) == true {
		fmt.Println("开始初始化")
		os.MkdirAll("img", 0666)
	}
}

func main() {
	num := input()
	time.Sleep(5 * time.Second)
	fmt.Println("请最小化此窗口，鼠标选中需要截图的窗口，5 秒后开始截图...")
	var a int
	for a < num {
		// 获取当前的时间
		localTime := getTime()
		// 截取日期
		pathName := strings.Split(localTime, "_")[0]
		// 创建当前日期目录
		fileToFolder(pathName)
		// 开始截图
		screenshots(pathName + "/" + localTime)
		// 下一页
		keyboard()
		a += 1
	}
}

// 截图
func screenshots(fileName string) {

	fileName = "img/" + fileName + ".png"

	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		// 加上一个延迟，避免被覆盖
		time.Sleep(2 * time.Second)

		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}

		//fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}
}

// 获取当前时间
func getTime() string {
	dateTime := time.Now().Format("2006-01-02 15:04:05")
	timeStr := strings.Replace(dateTime, " ", "_", 1)
	return timeStr
}

// 操作键盘
func keyboard() {
	robotgo.KeyTap("pagedown")
}

// 根据日期创建目录，如果存在就不创建
func fileToFolder(pathName string) {
	_, err := os.Stat("img/" + pathName)
	if os.IsNotExist(err) == true {
		os.MkdirAll("img/"+pathName, 0666)
	}
}

// 获取用户输入
func input() int {
	fmt.Print("请输入需要截屏的次数: ")
	var num int
	if _, err := fmt.Scanln(&num); err != nil {
		fmt.Println("输入有误，只允许输入数字!")
	} else {
		return num
	}
	return 0
}
