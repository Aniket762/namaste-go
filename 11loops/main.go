package main

import "fmt"

func main()  {
	fmt.Println("Getting into Loops!")
	
	days := [] string {"Sun","Sat","Fri","Wed","Mon"}
	fmt.Println(days)

	for index:=0;index<len(days);index++{
		
		if index%2 ==0{
			fmt.Println(days[index])
		}
	}

	// i will be index and we need
	for i := range days{
		fmt.Println(days[i])
	}

	// while loop
	var index2 int64 = 1 	

	for index2<10{
		fmt.Println("Value is: ", index2)

		// break
		if index2 ==5{
			break
		}

		// goto 
		if index2 ==4{
			goto hehe
		}

		// continue
		if index2 ==6{
			continue
		}
		index2++
	}
  hehe:
  fmt.Println("This is our goto label")

}