package main

import (
	"fmt"

	"github.com/diegoclair/microservices-codeeducation/tree/master/devops/docker/challenge03-CICD-Go/function"
)

func main() {
	result := function.AddTwoNumbers(5, 5)
	fmt.Println(result)

}
