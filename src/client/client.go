package client

import (
	"math/rand"

	"github.com/ulvimemmeedov/p2p/tools"
)

type Client struct {
	ID             int
	Username       string
	ContractAdress []byte
	Balance        float64
	Contracts      []string
	Password       string
}

func FindClientContractAdress(adress string) string {
	return adress
}
func NewUser(Username, Password string) Client {
	u := Client{}
	hash := tools.HashPassword(Password)
	u.ID = rand.Intn(99) * 9999
	u.Username = Username
	bs := tools.GenerateRandomBytes(Username, hash, u.ID)
	u.ContractAdress = bs
	u.Balance = 0
	u.Password = hash
	return u
}
