package main

import "fmt"
import "net/http"

type Course struct {
	Name		string
	Description	string
	Price		int
}

func soma(x int, y int) (int, bool) {
	if x > 10 {
		return x + y, true
	}
	return x + y, false
}

func main() {
	// var a string
	// a = "Hello, World!"
	fmt.Println("Hello, World!")
	
	// a := "Hello, World!" // string
	// resultado, status := soma(10, 20)
	println(soma(1, 2))

	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)

	course := Course{
		Name: "golang",
		Description: "golang course",
		Price: 100,
	}

	println(course.Name)

	fmt.Println(course.GetFullInfo())
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}