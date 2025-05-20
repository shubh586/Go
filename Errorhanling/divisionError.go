package main

import (
    "fmt"
)

// error is a interface in the go which have the method Error() string
type divisionError struct{
	divisior int
	dividend int
	message string
}
func (d divisionError) Error() string{
	return fmt.Sprintf("Error: %s, Divisor: %d, Dividend: %d", d.message, d.divisior, d.dividend)

}
func divide(a ,b int)(float32,error){
	if(b==0){
		return 0, divisionError{
			divisior: b,
			dividend: a,
			message: "Division by zero is not allowed",
		}

	}
	return float32(a)/float32(b), nil
}

func main2(){
     result,err:=divide(10,0)
	fmt.Println("Result:", err);
	//Error() method automatically called on err because it implements the error interface:
	// Go automatically calls err.Error() here
	//You’re telling Go:"Hey, this divisionError struct is an error because it has an Error() string method."
	//so no need to call err.Error() explicitly

	// if we have used the err instead of Dicerr showding may have occur as this check returns the err of tye divsion
	// err if failed then it will return the 0 . \
	//Shadowing happens when you declare a new variable with the same name as an existing variable. 
	// The new variable hides or overwrites access to the outer one — but only inside the smaller block.
	 if err!=nil{
		if Diverr,ok:=err.(divisionError); ok{
			fmt.Println("Error occurred:", Diverr)
		 }else{
			fmt.Println("Result unkkown error occured:", err)
		 }
	 }else{
		fmt.Println("Result:", result);
	 }

}