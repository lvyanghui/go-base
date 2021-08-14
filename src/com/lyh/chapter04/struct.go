package chapter04

import (
	"time"
	"fmt"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func Job() {
	dilbert.Salary -= 5000
	position := &dilbert.Position
	*position = "Senior" + *position

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	fmt.Println(EmployeeById(dilbert.ManagerID).Position)

	id := dilbert.ID
	EmployeeById(id).Salary = 0

	seen := make(map[string]struct{})
	if _, ok := seen["aa"]; !ok {
		seen["aa"] = struct{}{}
	}
	fmt.Printf("seen=%q\n", seen)

	p := Point{X: 1, Y: 2}
	fmt.Println(p)

	fmt.Println(Scale(p, 10))

	pp := &Point{1, 2}
	fmt.Println(pp)

	q := Point{2, 1}
	fmt.Println(p == q)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)

	/*var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20*/

	/*var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20*/

	var w Wheel
	/*w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20*/

	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Printf("%#v\n", w)
}

func EmployeeById(id int) *Employee {
	return &dilbert
}

type tree struct {
	value       int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values := appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

type Point struct {
	X, Y int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

type address struct {
	hostname string
	port     int
}

/*type Circle struct{
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}*/

/*type Circle struct{
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}*/

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
