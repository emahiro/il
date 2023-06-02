package main

import (
	"fmt"

	pb "github.com/emahiro/il/protobuf/pb/proto"
	"google.golang.org/protobuf/proto"
)

func main() {

	person := pb.Person{
		Name:  "Taro",
		Id:    1,
		Email: "taro@examle.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "090-1111-1111",
				Type:   pb.Person_PHONE_TYPE_MOBILE_UNSPECIFIED,
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
