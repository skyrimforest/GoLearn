/*
@File    :   main.go
@Time    :   2023/11/11 14:47:32
@Author  :   Skyrim
@Version :   1.0
@Site    :   https://github.com/skyrimforest
@Desc    :   工厂方法（Factory Method），又名虚拟构造函数。在父类中提供一个创建对象的方法，允许子类决定实例化对象的类型。
*/

package main

import "fmt"

/*
工厂方法模式建议使用特殊的工厂方法代替对于对象构造函数的直接调用，返回对象称为产品。
仅当产品有共同的基类/接口，才能返回不同类型的产品。
*/

type Product interface {
	doStuff()
}

type ProductA struct {
}

func (a *ProductA) doStuff() {
	fmt.Println("AdoStuff")
}

type ProductB struct {
}

func (b *ProductB) doStuff() {
	fmt.Println("BdoStuff")
}

type Creater interface {
	createProduct() Product
}

type ProductCreater struct {
}

func (c *ProductCreater) createProduct() Product {
	return nil
}

func (c *ProductCreater) doSth() {
	fmt.Println("233")
}

type ProductACreater struct {
	Creater
}

func (ac *ProductACreater) createProduct() Product {
	return &ProductA{}
}

type ProductBCreater struct {
	Creater
}

func (bc *ProductBCreater) createProduct() Product {
	return &ProductB{}
}

func main() {
	//下面这个，为了实现Creater，要么在ProductCreater实现createProduct前变成(c Product)，要么就在ProductCreater前加&号即可。
	var Master1 Creater = &ProductACreater{}
	var Master2 Creater = &ProductBCreater{}

	var product1 = Master1.createProduct()
	var product2 = Master2.createProduct()

	product1.doStuff()
	product2.doStuff()
}

/*
适用场景：
	无法预知对象确切类别及其依赖关系时
	易于扩展时
	可复用现有对象时
可能会演化为抽象工厂/原型/生成器等
/*