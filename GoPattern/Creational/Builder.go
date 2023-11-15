/*
@File    :   main.go
@Time    :   2023/11/11 22:06:37
@Author  :   Skyrim
@Version :   1.0
@Site    :   https://github.com/skyrimforest
@Desc    :   生成器（Builder）分步骤创建复杂对象，允许使用相同创建代码生成不同类型和形式的对象
*/

package main

import "fmt"

/*
若某房屋需要多种设施，则可以创建多种子类-类的层次结构负责，或者构造函数使用多个参数-构造函数的调用不简洁。
生成器模式建议将构造代码从产品类中抽取出来，并将其放在名为生成器的独立对象中。
不允许其他对象访问正在创建中的产品。需要不同形式的产品时使用不同生成器获得不同类对象。
可以创建一个主管类 定义创建步骤的执行顺序，而生成器则提供这些步骤的实现。
*/

type Builder interface {
	reset()
	buildStep1(num int)
	buildStep2()
	buildStepn()
}

type ConcreteBuilderA struct {
	pa *ProductA
}

func (cb *ConcreteBuilderA) reset() {
	cb.pa = NewProductA()
}
func (cb *ConcreteBuilderA) buildStep1(num int) {
	fmt.Println("First Step for ProductA,num is ", num)
}
func (cb *ConcreteBuilderA) buildStep2() {
	fmt.Println("Second Step for ProductA")
}
func (cb *ConcreteBuilderA) buildStepn() {
	fmt.Println("Nth Step for ProductA")
	fmt.Println()
}
func (cb *ConcreteBuilderA) getResult() *ProductA {
	res := cb.pa
	cb.reset()
	return res
}

type ProductA struct {
}

func NewProductA() *ProductA {
	return &ProductA{}
}
func (p1 *ProductA) doStuff() {
	fmt.Println("doStuffA")
}

type ConcreteBuilderB struct {
	pb ProductB
}

func (cb *ConcreteBuilderB) reset() {
	cb.pb = *NewProductB()
}
func (cb *ConcreteBuilderB) buildStep1(num int) {
	fmt.Println("First Step for ProductB,num is ", num)
}
func (cb *ConcreteBuilderB) buildStep2() {
	fmt.Println("Second Step for ProductB")
}
func (cb *ConcreteBuilderB) buildStepn() {
	fmt.Println("Nth Step for ProductB")
	fmt.Println()
}
func (cb *ConcreteBuilderB) getResult() ProductB {
	res := cb.pb
	cb.reset()
	return res
}

type ProductB struct {
}

func NewProductB() *ProductB {
	return &ProductB{}
}
func (p1 *ProductB) doStuff() {
	fmt.Println("doStuffB")
}

type Director struct {
	builder Builder
}

func NewDirector0() *Director {
	return &Director{}
}
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}
func (d *Director) setBuilder(builder Builder) {
	d.builder = builder
}

//注意下面其实没有用director自带的builder

func (d *Director) buildProductNumIs10(builder Builder) {
	builder.reset()
	builder.buildStep1(10)
	builder.buildStep2()
	builder.buildStepn()
}
func (d *Director) buildProductNumIs233(builder Builder) {
	builder.reset()
	builder.buildStep1(233)
	builder.buildStep2()
	builder.buildStepn()
}

type Client struct {
	director *Director
}

func (c *Client) MakeProduct() {
	c.director = NewDirector0()

	builder := &ConcreteBuilderA{}
	c.director.buildProductNumIs10(builder)
	pa := builder.getResult()

	builder2 := &ConcreteBuilderB{}
	c.director.buildProductNumIs10(builder2)
	pb := builder2.getResult()

	pa.doStuff()
	pb.doStuff()
}

func main() {
	client := Client{}
	client.MakeProduct()
}
