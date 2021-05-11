package main

import (
	"fmt"
	"log"
)

type Visitor interface {
	Visit(VisitorFunc) error
}

type VisitorFunc func(info *Info) error

type Info struct {
	Name      string
	NameSpace string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info)
}

type LogVisitor struct {
	visitor Visitor
}

func (l LogVisitor) Visit(fn VisitorFunc) error {
	return l.visitor.Visit(func(info *Info) error {
		log.Println("before")
		err := fn(info)
		if err != nil {
			return err
		}

		log.Println(info)
		log.Println("after")
		return nil
	})
}

// decorator visitor
type DecoratorVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratorVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}

	return DecoratorVisitor{v, fn}
}

func (v DecoratorVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info) error {
		if err := fn(info); err != nil {
			return err
		}

		for _, decorator := range v.decorators {
			if err := decorator(info); err != nil {
				return err
			}
		}

		return nil
	})
}

func main() {
	info := Info{}
	var v Visitor = &info
	// v = LogVisitor{v}

	// v.Visit(func(info *Info) error {
	// 	info.Name = "sakura"
	// 	info.NameSpace = "whirlwinder"
	// 	return nil
	// })

	decoratorVisitor := NewDecoratorVisitor(v, func(info *Info) error {
		log.Println("where")
		fmt.Println(info)
		return nil
	})

	decoratorVisitor.Visit(func(info *Info) error {
		info.Name = "sakura"
		info.NameSpace = "whirlwinder"
		return nil
	})
}
