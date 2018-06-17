package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main(){
	s := "Is this really magic?"

	s64 := base64.StdEncoding.EncodeToString([]byte(s))
	
	fmt.Println("Encoding...")
	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)


	fmt.Println("Decoding...")
	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("It's all gone horribly wrong!", err)
	}

	fmt.Println(string(bs))
}