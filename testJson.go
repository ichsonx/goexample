package main
/*
	测试json编码、解码
	解码是通过从文件读取内容解码，主要是用[]byte解码json内容称golang的struct结构
*/
import (
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
	_ "goexample/mycommon"
	"goexample/mycommon"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

//类型名、变量名需要首字母大写，否则其他函数内访问不了，缺一不可
type Targetlist struct {
	Names []string `json:"names"`
} 

func main() {
	//Jiema()
	Bianma()
}

//从json文件读取内容，解码成golang结构对象
func Jiema()  {
	filePath := "./config/targetlist.json"
	fi, err := os.Open(filePath)
	mycommon.Check(err)
	defer fi.Close()

	//json解码需要[]byte内容，所以使用ioutil.ReadAll来一次性加载文件内所有内容并返回[]byte
	filedata, err := ioutil.ReadAll(fi)
	mycommon.Check(err)

	target := Targetlist{}
	err = json.Unmarshal(filedata, &target)
	mycommon.Check(err)

	//遍历slice
	for _, name := range target.Names {
		fmt.Println(name)
	}

	//将对象打印成字符串输出;
	fmt.Println(target)
}

//从实例对象转换成json（二进制）
func Bianma()  {
	data, err := json.Marshal(movies)
	mycommon.Check(err)
	//打印[]byte，需要用printf来告诉golang是打印string
	fmt.Printf("%s\n", data)

	//这种方法进行编码，打印出来的json字符串带格式
	data2, err := json.MarshalIndent(movies, "", "    ")
	fmt.Printf("%s\n", data2)
}

