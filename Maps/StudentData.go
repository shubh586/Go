package main
import ("fmt")


type student struct{
	name string
	rollNo int
	marks map[string]int
}

func loopDepartment(department map[string]map[int]student){
	for departmentName,students:=range department{
		fmt.Println("Department Name:", departmentName)
		for rollNp,studentInfo:=range students{
			fmt.Println("Roll No:", rollNp)
			fmt.Println("Name:", studentInfo.name)
			for subjects,marks:=range studentInfo.marks{
				fmt.Println("Subject:", subjects, "Marks:", marks)
			}
			fmt.Println("===================================")
		}
	}
}


func main2(){
	//var m map[string]int this is just a declaration you need to initialize it m=make(map[string]int) like this or
	// directly initialize like between below  line
	//m:=make(map[string]int) or m:=map[string]int{};
	//m:=map[string]int; yese nahi chalega
	//value, ok := m["foo"]
	//delete(m, "foo")
// for key, value := range m {
//     fmt.Printf("Key: %s, Value: %d\n", key, value)
// }

// maps are reference types
// maps are not thread safe
//You cannot compare two maps with == (except comparing to nil):
//Keys must be types that are comparable with == and != (like strings, integers, pointers, structs without slices or maps).

 department :=make(map[string]map[int]student)

 cseStudent:=make(map[int]student)
 cseStudent[100]=student{
	name:"john",
	rollNo: 100,
	marks: map[string]int{
		"maths":100,
		"english":90,
		"science":95,
	},
}
cseStudent[101] = student{
	name:   "john",
	rollNo: 101,
	marks: map[string]int{
		"maths":   99,
		"english": 95,
		"science": 94,
	},
}
 department["CSE"]=cseStudent
fmt.Println("Printing the CSE department")
 loopDepartment(department);
 
 delete(cseStudent, 100)
 fmt.Println("Printing the CSE department after deletion")
 loopDepartment(department)
}

//delete in maps?