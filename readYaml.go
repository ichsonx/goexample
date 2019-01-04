package main

/*
使用的库及参考用法：https://github.com/go-yaml/yaml
以下提供了反序列化方法2个，序列化方法2个
在实践中，读取配置文件都是一次性读取的。
*/

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type developer struct {
	Name      string
	Age       int
	Weight    string
	Languages []string
}

type config struct {
	Fruits     []string          `yaml:"fruits"`
	Websites   map[string]string `yaml:"websites"`
	Developers []developer       `yaml:"developers"`
}

func main() {
	filepath := "./config/yamltext.yml"
	data, err := ioutil.ReadFile(filepath)
	yml := config{}

	//反序列化方法（1）
	//将读取到的[]byte内容 data，反序列化到变量对象 yml 中。
	err = yaml.Unmarshal([]byte(data), &yml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%v\n", yml)
	fmt.Printf("weight: %d",len(yml.Developers[0].Weight))

	//序列化方法（1）
	//将变量对象 yml 序列化到字符串变量 ymlstr 中。
	//ymlstr, err := yaml.Marshal(&yml)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- ymlstr dump:\n%s\n\n", string(ymlstr))
	//
	////反序列化方法（2）
	////将读取到的[]byte内容 data，反序列化到map变量 ymlMap 中。
	//ymlMap := make(map[interface{}]interface{})
	//err = yaml.Unmarshal([]byte(data), &ymlMap)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- ymlMap:\n%v\n\n", ymlMap)
	//
	////序列化方法（2）
	////将map变量 ymlMap 序列化到字符串变量 ymlstr 中。
	//ymlstr, err = yaml.Marshal(&ymlMap)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//fmt.Printf("--- ymlMap dump:\n%s\n\n", string(ymlstr))

}
