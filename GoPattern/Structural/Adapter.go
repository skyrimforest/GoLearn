/*
@File    :   main.go
@Time    :   2023/11/15 23:30:23
@Author  :   Skyrim
@Version :   1.0
@Site    :   https://github.com/skyrimforest
@Desc    :   适配器（Wrapper或Adapter）是一种结构型设计模式，它能使接口不兼容的对象能够相互合作。
*/

package main

import (
	"fmt"
	"math"
)

/*
适配器模式通过封装对象将复杂的转换过程隐藏于幕后。
1. 适配器实现与其中一个现有对象兼容的接口。
2. 现有对象可以使用该接口安全地调用适配器方法。
3. 适配器方法被调用后将以另一个对象兼容的格式和顺序将请求传递给该对象。、
适配器分为对象适配器与类适配器两种。
下面以方钉圆孔问题为例实现一个类适配器。
非常遗憾，go的组合式继承不能实现多态，必须使用接口。
*/

type sqareInterface interface {
	getWidth() float64
}

type SquarePeg struct {
	width float64
}

func (sp *SquarePeg) getWidth() float64 {
	return sp.width
}

type roundInterface interface {
	getRadius() float64
}

type RoundPeg struct {
	radius float64
}

func (rp *RoundPeg) getRadius() float64 {
	return rp.radius
}

type SquarePegAdapter struct {
	RoundPeg
	peg SquarePeg
}

func (spa *SquarePegAdapter) getRadius() float64 {
	return float64(spa.peg.getWidth()) * math.Sqrt2 / 2.0
}

type RoundHole struct {
	radius float64
}

func (rh *RoundHole) fits(peg roundInterface) bool {
	return rh.radius >= peg.getRadius()
}

func main() {
	hole := &RoundHole{
		radius: 5,
	}
	rpeg := &RoundPeg{
		radius: 5,
	}
	fmt.Println(hole.fits(rpeg))

	small_sqpeg := &SquarePeg{
		width: 5,
	}
	big_sqpeg := &SquarePeg{
		width: 10,
	}
	small_sqpeg_adapter := &SquarePegAdapter{
		peg: *small_sqpeg,
	}
	big_sqpeg_adapter := &SquarePegAdapter{
		peg: *big_sqpeg,
	}
	fmt.Println(hole.fits(small_sqpeg_adapter))
	fmt.Println(hole.fits(big_sqpeg_adapter))

}
