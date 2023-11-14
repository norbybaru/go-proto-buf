package main

import (
	"log"
	"os"

	userpb "github.com/norbybaru/go-proto-buf/gen/go/user/v1"
	"google.golang.org/protobuf/proto"
)

func main() {
  user := userpb.User{
    Uuid: "1a2b3c",
    FullName: "Norby",
    BirthYear: 2001,
  }

  b, err := proto.Marshal(&user)

  if err != nil {
    log.Fatalln("Marshall error: ", err)
  }

  if err := os.WriteFile("user.bin", b, 0644); err != nil {
    log.Fatalln("Write file error: ", err)
  }
}
