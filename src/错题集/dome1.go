package main

import "fmt"

type Lili struct {

}

func (li *Lili)fmtPointer()  {
	fmt.Println("pointer")
}

func (li Lili)fmtReference()  {
  fmt.Println("reference")
}

//main主函数中的“li”是一个变量，li的虽然是类型Lili，但是li是可以寻址的，&li的类型是Lili，因此可以调用Lili的方法。
func main() {
	li:=Lili{}
	li.fmtPointer()

	//Lili{}.fmtPointer() //编译报错
}
