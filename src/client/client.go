package client

import (
	"crypto/sha1"
	"math/rand"
	"strconv"
)

var character string = "qwertyuikasdnhshbdsjadbjhasndjnaskdasdandaksdsddknasdkasdkaskndaskdadkasnkdknaksnd"

type Client struct {
	ID             int
	Username       string
	ContractAdress []byte
	Balance        float64
	Contracts      []string
}

func NewUser(Username string) Client {
	h := sha1.New()
	u := Client{}
	u.ID = rand.Intn(99) * 9999
	u.Username = character + Username + character + strconv.Itoa(u.ID)
	h.Write([]byte(Username))
	bs := h.Sum(nil)
	u.ContractAdress = bs
	u.Balance = 0
	return u
}
