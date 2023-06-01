package main

import (
	"fmt"

	"github.com/emahiro/il/protobuf/pb/tutrialpb"
	"google.golang.org/protobuf/proto"
)

func main() {

	person := tutrialpb.Person{
		Name:  "Taro",
		Id:    1,
		Email: "taro@examle.com",
		Phones: []*tutrialpb.Person_PhoneNumber{
			{
				Number: "090-1111-1111",
				Type:   tutrialpb.Person_MOBILE,
			},
		},
	}
	b, err := proto.Marshal(&person)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
