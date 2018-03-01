package main

import (
	"github.com/BurntSushi/toml"
	"fmt"
	"os"
)

func main() {
	ReadToml()
}

type tomlconfig struct {
	Title		string
	Constraints	constraints
}

type constraints struct {
	Constraint	[]constraint
}

type constraint struct {
	Branch 	string
	Name	string
	Version	string
}


//read toml by struct way
func ReadToml()  {
	var conf tomlconfig
	filepath := "./config/tomlconfig.toml"
	if _, err := toml.DecodeFile(filepath, &conf); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	for _, v := range conf.Constraints.Constraint{
		fmt.Println(v.Name)
	}
}
