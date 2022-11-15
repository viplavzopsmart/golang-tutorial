package main

import (
	"fmt"
	testModel "github.com/modTutorial/models"
	"github.com/pborman/uuid"
	"strings"
)

func main() {
	fmt.Println("Package tutorial")
	fmt.Println(testModel.Name)
	fmt.Println(testModel.GetMyPhoneNumber())
	testModel.PrintTestVariable()

	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
}
