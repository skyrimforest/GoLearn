/*
@File    :   main.go
@Time    :   2023/11/11 21:32:58
@Author  :   Skyrim
@Version :   1.0
@Site    :   https://github.com/skyrimforest
@Desc    :   抽象工厂（Abstract Factory）创建一系列相关的对象，无需指定具体类。
*/

package main

import "fmt"

/*
有四种产品：type1ProductA,type1ProductB,type2ProductA,type2ProductB
所以创建了一个抽象工厂接口，俩具体工厂；俩产品接口，四种具体产品。
*/

type AbstractFactory interface {
	createProductA() ProductA
	createProductB() ProductB
}

type Type1Factory struct {
}

func (tf *Type1Factory) createProductA() ProductA {
	return &type1ProductA{}
}
func (tf *Type1Factory) createProductB() ProductB {
	return &type1ProductB{}
}

type Type2Factory struct {
}

func (tf *Type2Factory) createProductA() ProductA {
	return &type2ProductA{}
}
func (tf *Type2Factory) createProductB() ProductB {
	return &type2ProductB{}
}

type ProductA interface {
	doStuffA()
}
type ProductB interface {
	doStuffB()
}

type type1ProductA struct {
}

func (tp *type1ProductA) doStuffA() {
	fmt.Println("type1doStuffA")
}

type type2ProductA struct {
}

func (tp *type2ProductA) doStuffA() {
	fmt.Println("type2doStuffA")
}

type type1ProductB struct {
}

func (tp *type1ProductB) doStuffB() {
	fmt.Println("type1doStuffB")
}

type type2ProductB struct {
}

func (tp *type2ProductB) doStuffB() {
	fmt.Println("type2doStuffB")
}

type Client struct {
	af AbstractFactory
	pa ProductA
	pb ProductB
}

func NewClient(factory AbstractFactory) *Client {
	return &Client{
		af: factory,
	}
}
func (c *Client) getPa() {
	c.pa = c.af.createProductA()
}
func (c *Client) getPb() {
	c.pb = c.af.createProductB()
}
func (c *Client) usePa() {
	c.pa.doStuffA()
}
func (c *Client) usePb() {
	c.pb.doStuffB()
}

func main() {

	var ClientA *Client = NewClient(&Type1Factory{})
	var ClientB *Client = NewClient(&Type2Factory{})

	ClientA.getPa()
	ClientA.getPb()
	ClientA.usePa()
	ClientA.usePb()

	ClientB.getPa()
	ClientB.getPb()
	ClientB.usePa()
	ClientB.usePb()
}
