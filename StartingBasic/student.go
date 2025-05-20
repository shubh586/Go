package main
import "fmt"




type address struct{
	city string
	state string 
	pinCode int
}

type  student struct{
	name string
	age int
	rollNo int
	address
}



type teacher struct {
	name   string
	subject string
}


type displySchool interface{
	displySchool() string
}

func (t teacher) displySchool() string {
	return fmt.Sprintf("Teacher Name: %s, Subject: %s", t.name, t.subject)
}

func (s student ) displySchool( ) string{
	return fmt.Sprintf("Name: %s, Age: %d, Roll No: %d, City: %s, State: %s, PinCode: %d", s.name, s.age, s.rollNo, s.city, s.state, s.pinCode)
}

func main2(){
	p:= student{

		name : "shubham",
		rollNo: 56,
		age:20,
		address:address{
			city: "pune",
			state: "maharashtra",
			pinCode: 411057,
		},
		
	}
	t:=teacher{
		name:"shubh",
		subject:"math",
	}
	k1:=t.displySchool()
	fmt.Println(k1)
	fmt.Println("In the main function 2")
	k:=p.displySchool()
	fmt.Println(k)

	var members []displySchool = []displySchool{p, t}// this is a slice of interface type

	for _, member := range members {
		fmt.Println(member.displySchool()) // Polymorphism: one function call, different results
	}

	//var arr [5]int=[5]int{1,2,3,4,5};  valid declaration
	// arr:=[5]int{1,2,3,4,5}   valid declaration


}



// go mod init <module-name> — Go Modules (Short Notes)
// Initializes a Go module in your project.
// Creates a go.mod file to:
// Define the module name.
// Track dependencies.
// Required for using go run ., go build, and go get.

