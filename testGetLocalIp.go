/*
	获取本地ip、mac地址
*/

package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	list, _ := net.Interfaces()
	for _, v := range list{
		fmt.Println(v.Name)
	}
}

func getlocalIPV4()  {
	//linux中以太网网卡名称为eth0，windows中第一网卡名为“以太网”
	inter, err := net.InterfaceByName("eth0")
	if err != nil {
		fmt.Println(err)
	}
	//mac地址
	fmt.Println(inter.HardwareAddr.String())
	addrs, err := inter.Addrs()
	if err != nil {
		fmt.Println(err)
	}
	//ip地址一个ip4一个ip6，这样后面带子网掩码数字
	for _, addr := range addrs {
		fmt.Println(addr.String())
	}

	//这样就只显示纯正的ipv4地址了，循环中做了判断
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

//获取所有本地ipv4的ip地址，如果装了虚拟机就会显示n个ip地址
func getLocalIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				fmt.Println(address)
			}
		}
	}
}

func getLocalIp2() {
	//以太网网卡名称为eth0
	inter, err := net.InterfaceByName("以太网")
	if err != nil {
	}
	//mac地址
	fmt.Println(inter.HardwareAddr.String())
	addrs, err := inter.Addrs()
	if err != nil {
	}
	//ip地址一个ip4一个ip6
	//for _, addr := range addrs {
	//	fmt.Println(addr.String())
	//}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}