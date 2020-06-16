package main

import (
	"fmt"

	"github.com/tddey01/luffy/day004/02package/math_pkg"
)

const Mode = 1

func main() {
	math_pkg.Add(100, 200)
	stu := math_pkg.Student{Name: "豪杰", Age: 19}
	fmt.Println(stu.Name)
	fmt.Println(stu.Age)
}
