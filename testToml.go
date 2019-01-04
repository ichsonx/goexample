/*
	2018-03-02
	toml配置文件读取方法
*/

package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"time"
)

func main() {
	ReadToml()
	//AnotherWayReadAndShow()
}
type constraint struct {
	Branch	string
	Name	string
	Version	string
}
type tomlconfig struct {
	Title		string
	Constraints	[]constraint `toml:"constraint"`
	Owner   ownerInfo
	DB      database `toml:"database"`
	Servers map[string]server
	Clients clients
}
type ownerInfo struct {
	Name string
	Org  string `toml:"organization"`
	Bio  string
	DOB  time.Time
}

type database struct {
	Server  string
	Ports   []int
	ConnMax int `toml:"connection_max"`
	Enabled bool
}

type server struct {
	IP string
	DC string
}

type clients struct {
	Data  [][]interface{}
	Hosts []string
}


//read toml by struct way
/*
	使用结构类型来从toml文件读取数据
*/
func ReadToml()  {
	var config tomlconfig
	filepath := "./config/tomlconfig.toml"
	if _, err := toml.DecodeFile(filepath, &config); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	for _, v := range config.Constraints{
		fmt.Println(v.Name)
	}

	fmt.Printf("Title: %s\n", config.Title)
	fmt.Printf("Owner: %s (%s, %s), Born: %s\n",
		config.Owner.Name, config.Owner.Org, config.Owner.Bio,
		config.Owner.DOB)
	fmt.Printf("Database: %s %v (Max conn. %d), Enabled? %v\n",
		config.DB.Server, config.DB.Ports, config.DB.ConnMax,
		config.DB.Enabled)
	for serverName, server := range config.Servers {
		fmt.Printf("Server: %s (%s, %s)\n", serverName, server.IP, server.DC)
	}
	fmt.Printf("Client data: %v\n", config.Clients.Data)
	fmt.Printf("Client hosts: %v\n", config.Clients.Hosts)
}

/*
	使用 map 映射类型从toml读取数据
*/
func AnotherWayReadAndShow()  {
	filepath := "./config/tomlconfig.toml"
	data := make(map[string]interface{})
	toml.DecodeFile(filepath, &data)
	fmt.Println(len(data))

	for key, val := range data {
		fmt.Println(key)
		fmt.Println(val)
	}

	port := data["database"]
	fmt.Println(port)
}
