package main

import (
	"encoding/json"
	"fmt"
)

// creating aliases with ``
type course struct{
	Name string `json:"coursename"`
	Price int	`json:"courseprice"`
	Platform string
	Password string `json:"-"`  // - ignores the password from display
	Tags []string `json:"tags,omitempty"` // if nil is present it will ignore
}

func main()  {
	fmt.Println("Getting into JSON 🙏")
	EncodeJson()
}

// converting any sort of data to json
func EncodeJson()  {
	aniketCourses := []course{
		{"Pepejs",0,"youtube","abc",[]string{"web","front"}},
		{"Lalajs",100,"youtube","abc",nil},
		{"Hehejs",5000,"youtube","abc",[]string{"web","front"}},
	}

	// package this data into JSON
	finalJson, err := json.Marshal(aniketCourses)

	// for increasing readability of json
	// we use Marshal.Indent
	// json.MarshalIndent(aniketCourses,"","\t")


	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n",finalJson)
}

func DecodeJson()  {
	jsonDataFromWeb := []byte(`
		{
			"fname":"aniket",
			"lname":"pal"
		}
	`)

	var acourse course 
	checkValid := json.Valid(jsonDataFromWeb) // return bool

	// unmarshal decode's json and tell how it should look
	if checkValid{
		fmt.Println("JSON was valid")	
		json.Unmarshal(jsonDataFromWeb,&acourse)
		fmt.Printf("%#v\n",acourse)
	} else{
		fmt.Println("JSON was not valid")
	}
	
	// cases where you just want to add data to key value
	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&onlineData)
	fmt.Printf("%v\n",onlineData)

	for k , v := range onlineData {
		fmt.Printf("key is %v and value is %v",k,v)
	}
}