package main

import "fmt"

type Foo interface {
	HelloFoo()
}

type FooImpl struct{}

func (f *FooImpl) HelloFoo() {
	fmt.Println("i'm fooo")
}

func NewFoo() *FooImpl {
	return &FooImpl{}
}

type Boo interface {
	HelloBoo()
}

type BooImpl struct {
}

func (f *BooImpl) HelloBoo() {
	fmt.Println("i'm booo")
}

func NewBoo() *BooImpl {
	return &BooImpl{}
}

type Adapter struct {
	boo BooImpl
}

func (a *Adapter) HelloFoo() {
	a.boo.HelloBoo()
}

func main() {
	adapter := &Adapter{boo: BooImpl{}}

	var foo Foo = adapter

	foo.HelloFoo()
}
