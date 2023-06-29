// package main

// import (
// 	"fmt"
// 	"math"
// )

// type geometry interface {
// 	before()
// 	area() float64
// 	perim() float64
// 	// after(beforeStr)
// }

// type rect struct {
// 	width, height float64
// }
// type circle struct {
// 	radius float64
// }

// func (r rect) area() float64 {
// 	return r.width * r.height
// }
// func (r rect) perim() float64 {
// 	return 2*r.width + 2*r.height
// }

// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }
// func (c circle) perim() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func measure(g geometry) {
// 	// beforeStr := before()
// 	// defer after(beforeStr)
// 	fmt.Println(g)
// g.before()
// 	fmt.Println(g.area())
// 	// log.Fatal("Just for test")
// 	fmt.Println(g.perim())
// }

// func main() {
// 	r := rect{width: 3, height: 4}
// 	// c := circle{radius: 5}

// 	measure(r)
// 	// measure(c)
// }

// func (r rect) before() {
// 	before := "before"
// 	fmt.Println(before)
// 	defer fmt.Println("after")

// }

// func after(before string) {
// 	fmt.Println(before)
// 	fmt.Println("now it is after")

// }

package main

import (
	"fmt"
	"log"
)

type DBConnector interface {
	connectDB() error
	func1()
	func10()
	Close()
}

type MyDBConnector struct {
	db       string
	isClosed bool
}

func (m *MyDBConnector) connectDB() error {
	// Connect to the database
	db := "db opened"

	// Store the database connection
	m.db = db
	m.isClosed = false

	return nil
}

func (m *MyDBConnector) func1() {
	if m.isClosed {
		log.Fatal("Database connection is closed")
	}
	panic("any panic")
	// Use m.db for func1 logic
}

func (m *MyDBConnector) func10() {
	if m.isClosed {
		log.Fatal("Database connection is closed")
	}
	// Use m.db for func10 logic
	// Use m.db for func1 logic

}

func (m *MyDBConnector) Close() {
	if m.db != "" && !m.isClosed {
		m.db = "Closed"
		fmt.Println(m.db)
		m.isClosed = true
	}
}

func main() {
	connector := &MyDBConnector{}
	if err := connector.connectDB(); err != nil {
		log.Fatal(err)
	}
	defer connector.Close()

	// Call func1
	connector.func1()

	// Call func10
	connector.func10()
	fmt.Println(connector.db)
}
