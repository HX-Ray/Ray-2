package main

import(
	"fmt"
)

type cuboid struct{
	lenth int
	width int
	hight int
}

func main(){
	var brick cuboid
	fmt.Print("长：")
	fmt.Scan(&brick.lenth)
	fmt.Print("宽：")
	fmt.Scan(&brick.width)
	fmt.Print("高：")
	fmt.Scan(&brick.hight)

	fmt.Printf("体积为%d\n", brick.V())
	fmt.Printf("表面积为%d\n", brick.S())

}

func (b cuboid) V() int {
	return b.lenth * b.width * b.hight
}

func (b cuboid) S() int {
	return (b.lenth * b.width + b.lenth * b.hight + b.width * b.hight) * 2
}
