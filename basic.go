package main

import "fmt"

type M struct {
	name string
}

func (m M)change()  {
	m.name = "aaa"
}
func (m *M)change2()  {
	m.name = "aaa"
}

func main() {
	//==============================指针、值接受者方法的区别================================
	//m := M{"sonx"}
	//m.change()
	//fmt.Println(m.name)	//打印的仍然是sonx
	//m.change2()
	//fmt.Println(m.name)	//打印的是aaa

	m := new(M)
	m.name = "yes"
	m.change()
	fmt.Println(m.name)	//打印的仍然是yes
	m.change2()
	fmt.Println(m.name)	//打印的是aaa
	//结论，无论是指针，还是值接受者，都不影响方法的使用。可以互相使用。值接受者，调用指针接受者方法，仍然可以达到指针效果。

	//===============================new、make、指针的关系==============================
	//var apot *string
	//apot = new(string)
	//*apot = "hello"
	//fmt.Println(*apot)
	//
	//var mapPot *map[string]int
	//
	//fmt.Printf("mapPot: %p %#v \n", &mapPot, mapPot) //  mapPot: 0xc42000c050 (*map[string]int)(nil)
	//// 初始化map指针的地址
	//mapPot = new(map[string]int)
	//
	//fmt.Printf("mapPot: %p %#v \n", &mapPot, mapPot) // mapPot: 0xc42000c050 &map[string]int(nil)
	//
	////(*mapPot)["age"] = 21 // 报错
	//// 初始化map指针指向的map
	//(*mapPot) = make(map[string]int)
	//(*mapPot)["age"] = 21
	//fmt.Printf("mapPot: %p %#v \n", &mapPot, mapPot)
	//=============================================================

}