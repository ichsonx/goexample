package main

import (
	"fmt"
)

func main() {
	rangeRemove()

}

//当slice删除一个元素，并且知道下标的时候，可以用以下方式很方便
func singleRemove()  {
	s := []int{1,2,3,4,5,6,7,8,9,10}
	s = append(s[:9], s[10:]...)
}

/*
	实际中可能会对slice进行遍历，以筛选出需要的结果slice，这时会用到for循环。
	这时就不能用singleRemove()中的方法了，因为它会改变slice的index

	例如：删除了值=2后，值=3的index会前移（从index=2变成了index=1，占了原来值=2的位置），
	而for循环的index不会因此改变，在index=1的时候删除了值=2，下一次的循环将访问（index=2，值=4），错过了（值=3）

	所以如果要遍历slice来筛选删除元素，用以下方法将会是最佳实践
*/
func rangeRemove()  {
	s := []int{1}

	k := 0
	for _, val := range s{
		if val !=1{
			s[k] = val
			k++
		}
	}
	s = s[:k]
	fmt.Println(s)
}
