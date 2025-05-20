// in one folder only one package can be there. if you want to use the same package name in 
// another file then you have to create a new folder and put the file in that folder.
// basically one floder=one package.



package main
import "fmt"

// type person struct{
// 	name string
// 	age int
// }  it also getting error because of the same name in the other file . redeclaration
// func (p person) display()string{
// 	return fmt.Sprintf(("name %s age %d"), p.name,p.age)
// }

type company struct{
	name string
	salary int32
}
func (c company) display()string{
	return fmt.Sprintf(("name %s age %d"), c.name,c.salary)
}
//  type display interface{
// 	display() string
//  }  redeclaration 
 func showDetails(d display) {
	fmt.Println(d.display())
}

func displayInfo(){
	p:=person{"shubham",20,1000.50}
	c:=company{"tesla",1000}
	showDetails(p)
	showDetails(c)
}

