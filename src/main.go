package main

import (
	"fmt"

	"github.com/ulvimemmeedov/p2p/client"
	"github.com/ulvimemmeedov/p2p/contracts"
)

func main() {
	u := client.NewUser("UlviMemmeedov", "0519770884")
	contracts.GenerateContract("rasim", "test")
	fmt.Println(u)

}
