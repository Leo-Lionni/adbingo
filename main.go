package main

import (
	"fmt"
	"log"
	"os"
)

type  person struct {
	age int
	name string
	sss  map[string]int
	//c  chan int
}
var p person  //这里只是作出statemeng ；然后给出 编译器是认为 nil
//如果需要初始化；为 非零值；则需要手动指定字面值；
func main() {
	//var names []string
	var dict map[string]int // only statement ;  can be nil;
	dict = make(map[string]int)  // memory  assign
		dict["Changsha"]=12//if no mem ; will be panic
		dict["lisi"] =13
		//for names
//		for k,v :=range dict{
//			fmt.Println("name , age ",k,v) //k ;key ;    v  ,value ;
//			//when one , will key ;
//			names = append(names,k)
//}
//		sort.Strings(names)
//		fmt.Println(sort.StringsAreSorted(names))
//		for _,key := range names{
//			fmt.Println(key,dict[key])
//fmt.Println(dict["zhangsan"])
//		fmt.Println(dict["lisi"])
//		modify(dict)// 并非传递的是  map的副本 ； 与切片相似；传递的是"一个引用吧"
//	fmt.Println(dict["lisi"])
//		fmt.Printf("%p,%#v\n",&p,p)
//}

 jim := person{10,"chenhuiying",map[string]int{"张三": 20}}
 fmt.Println(jim)
 modify(jim)
 fmt.Println(jim)
 file ,err := os.Open("/usr/tmp")
 if err != nil{
 	log.Fatal(err)
	 return // 直接返回了函数
 }
 fmt.Println(file)
}
//func modify(dict map[string]int)  {
//	dict["lisi"]= 22
//}
//显示传递指针
//；  其实也传递过来指针的副本；
func print_all(a  ...interface{} ) {
	for _,v := range a{
		fmt.Print(v)  // value ; if only one ,will be key;
	}
	fmt.Println()
}
func  modify(p person)  {
	p.age = p.age+10 // 传递过来的是  type person 类型 本身和 里面的值的拷贝
	p.sss["张三"]=33
}
//func add(a,b int)(int ,error)  {
//	//c := 10
//	return
//}
/*
1. 函数传递的是值类型，
2. 传递过来的是 结构体的本身和内部值的复制
3. 如果遇到 结构体内部有"引用类型"，即便是传递的是值类型，也会修改内容。
4.  可以发出
 */