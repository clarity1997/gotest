package main

import "fmt"

//如果类名首字母大写则表示其他包可以访问，如果类的属性大写，则该属性可以在类外部访问
//go语言中的大小写是有区别的
type Hero struct {
	name string
	age  int
}

func (this Hero) setName(name string) {
	//this 是调用该方法的对象的一个副本
	this.name = name
}
func (this *Hero) setNameTrue(name string) {
	this.name = name
}
func (this Hero) getName() {
	fmt.Println(this.name)
}
func main() {
	var hero Hero
	hero.name = "liuhao"
	hero.age = 23
	hero.setName("mahuan")
	hero.getName()
	hero.setNameTrue("xiaojuju")
	hero.getName()
}
