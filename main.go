package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func generateMatrix(n int) [][]int{
	//构造n x n的矩阵
	res:=make([][]int,n)
	for i:=0;i<n;i++{
		res[i]=make([]int,n)
	}
	num:=1
	left,right,top,down:=0,n-1,0,n-1
	for num<=n*n{
		//top,left--->top,right
		for i:=left;i<=right;i++{
			res[top][i]=num
			num++
		}
		top++
		//top,right--->down,right
		for i:=top;i<=down;i++{
			res[i][right]=num
			num++
		}
		right--
		//down,right---->down,left
		for i:=right;i>=left;i--{
			res[down][i]=num
			num++
		}
		down--
		//down,left---->top,left
		for i:=down;i>=top;i--{
			res[i][left]=num
			num++
		}
		left++
	}
	return res
}
func generateMatrixTest(){
	input:=bufio.NewScanner(os.Stdin)
	input.Scan()//读取N
	n,_:=strconv.Atoi(input.Text())
	res:=generateMatrix(n)
	fmt.Println(res)
}
func main(){
	generateMatrixTest()
}
