package main

/*
使用的库及参考用法：https://github.com/go-yaml/yaml
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

	err = yaml.Unmarshal([]byte(data), &yml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("%v\n", yml)
	for _, developer := range yml.Developers{
		fmt.Printf(" %s\n", developer.Languages[0])
	}

	d, err := yaml.Marshal(&yml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))

}
