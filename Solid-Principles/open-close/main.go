package main

import (
	"fmt"
)

/*
	Suppose we have a product list.
	We want to implement filter over it.
		- filter by color
		- filter by size
		- filter by color and size
		- ..
		- .. many more.
*/

type Size int

const (
	small Size = iota
	medium
	large
)

type Color int

const (
	green Color = iota
	red
	blue
)

type Product struct {
	name  string
	size  Size
	color Color
}

// Naive approach
type Filter struct {
}

// filter by color
func (f *Filter) FilterByColor(product []Product, color Color) []Product {
	var result []Product
	for i, v := range product {
		if v.color == color {
			result = append(result, product[i])
		}
	}
	return result
}

// filter by size
func (f *Filter) FilterBySize(product []Product, size Size) []Product {
	// Implementation
	return []Product{}
}

// filter by color and size
func (f *Filter) FilterByColorAndSize(product []Product, size Size, color Color) []Product {
	// Implementation
	return []Product{}
}

// We can see more requirement will increase more complexity and will end up in ugly code.
// ***********************************************************************

// Better Approach.
type Specification interface {
	IsSatisfied(p Product) bool
}

type ColorSpecification struct {
	Color Color
}

func (cs *ColorSpecification) IsSatisfied(p Product) bool {
	return p.color == cs.Color
}

type SizeSpecification struct {
	Size Size
}

func (ss *SizeSpecification) IsSatisfied(p Product) bool {
	return p.size == ss.Size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(product []Product, spec Specification) []Product {
	var result []Product
	for i, v := range product {
		if spec.IsSatisfied(v) {
			result = append(result, product[i])
		}
	}

	return result
}

func main() {
	fmt.Println("*****************Open close Principle*****************")

	tree := Product{"tree", large, green}
	apple := Product{"apple", small, green}
	phone := Product{"Motorola", medium, red}

	productList := []Product{
		tree, apple, phone,
	}
	fmt.Println("Test With Naive Approach")

	var f Filter
	list := f.FilterByColor(productList, green)
	for _, p := range list {
		fmt.Println(p.name)
	}

	fmt.Println("Test With Better Approach")
}
