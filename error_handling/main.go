package main 
import "fmt"

// func divide(a,b float64) (float64,error) {
// 	if b==0{
// 		return 0 , fmt.Errorf("denominator not to be zero")
// 	}
// 	return a/b , nil
// }

func divide(a,b float64)(float64,string){
	if b==0{
		return 0, "denominator must not be zero"
	}
	return a/b, "nil"
}

func main(){
	fmt.Println("started error handling in functions")
	// ans,err:= divide(10,0);
	// if err!= nil{
	// 	fmt.Println("Error handling")
	// }
	// fmt.Println("Division of 2 numbers:", ans)

	ans,_ := divide(10,0)
	fmt.Println("division of 2 number is",ans)
}