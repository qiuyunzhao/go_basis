package main

import "fmt"

//作用：defer的引入，为了在函数执行完毕后及时释放资源
//说明：
//1.当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
//2.当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
//3.defer将语句入栈时，也会将语句相关的值进行拷贝，同时入栈

func sum(n1 int, n2 int) int {
	//defer 后边的会在该函数执行完成后再执行
	defer fmt.Println("ok1 n1=", n1) //执行顺序 3. ok1 n1 = 10
	defer fmt.Println("ok2 n2=", n2) //执行顺序 2. ok2 n2= 20
	//增加一句话
	n1++                         // n1 = 11
	n2++                         // n2 = 21
	res := n1 + n2               // res = 32
	fmt.Println("ok3 res=", res) // 执行顺序1. ok3 res= 32
	return res
}

func test() {
	fmt.Println("test()") // 执行顺序4. res= 32
}

func main() {
	res := sum(10, 20)
	test()
	fmt.Println("res=", res) // 执行顺序5. res= 32
}
