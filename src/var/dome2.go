package main


//定义类型
type Color int


const(
	Black Color=iota
	Red
	Blue
)

func test(color Color)  {

}
func main() {
	c:=Black
	test(c)


	i:=1
	test(i)


}
