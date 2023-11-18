package main

import "fmt"

// 抽象产品接口
type Shape interface {
	Draw() string
}

type Color interface {
	Fill() string
}

// 具体产品类型A
type Circle struct{}

func (c *Circle) Draw() string {
	return "Draw Circle"
}

type Red struct{}

func (r *Red) Fill() string {
	return "Fill with Red color"
}

// 具体产品类型B
type Square struct{}

func (s *Square) Draw() string {
	return "Draw Square"
}

type Blue struct{}

func (b *Blue) Fill() string {
	return "Fill with Blue color"
}

// 抽象工厂接口
type AbstractFactory interface {
	CreateShape() Shape
	CreateColor() Color
}

// 具体工厂类型A
type ShapeAndColorFactoryA struct{}

func (f *ShapeAndColorFactoryA) CreateShape() Shape {
	return &Circle{}
}

func (f *ShapeAndColorFactoryA) CreateColor() Color {
	return &Red{}
}

// 具体工厂类型B
type ShapeAndColorFactoryB struct{}

func (f *ShapeAndColorFactoryB) CreateShape() Shape {
	return &Square{}
}

func (f *ShapeAndColorFactoryB) CreateColor() Color {
	return &Blue{}
}

// 客户端
func main() {
	// 选择具体工厂类型
	factoryA := &ShapeAndColorFactoryA{}
	factoryB := &ShapeAndColorFactoryB{}

	// 使用工厂A创建产品
	shapeA := factoryA.CreateShape()
	colorA := factoryA.CreateColor()

	// 使用工厂B创建产品
	shapeB := factoryB.CreateShape()
	colorB := factoryB.CreateColor()

	// 展示产品信息
	fmt.Println("Factory A products:")
	fmt.Println(shapeA.Draw())
	fmt.Println(colorA.Fill())

	fmt.Println("\nFactory B products:")
	fmt.Println(shapeB.Draw())
	fmt.Println(colorB.Fill())
}
