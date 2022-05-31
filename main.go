package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func moveZeroes(nums []int){
	slowIndex:=0
	for fastIndex:=0;fastIndex<len(nums);fastIndex++{
		if nums[fastIndex]!=0{
			nums[slowIndex]=nums[fastIndex]
			slowIndex++
		}
	}
	for i:=slowIndex;i<len(nums);i++{
		nums[i]=0
	}
}
func moveZeroesTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()
	var nums []int
	str:=strings.Split(input.Text(),",")
	for _,s:=range str{
		v,_:=strconv.Atoi(s)
		nums=append(nums,v)
	}

	fmt.Println("移除前的数组为：",nums)
	moveZeroes(nums)
	fmt.Println("移除后的数组为：",nums)
}
func main(){
	moveZeroesTest()
}