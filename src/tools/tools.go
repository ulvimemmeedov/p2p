package tools

import (
	"crypto/sha1"
	"fmt"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var character string = "qwertyuikasd1233123113#@$%^&*^&%$#@!#$%^&&??????D/SA'/'A/'/SA'/DD'//'SA'/'/'/'/'/ADS'/'/SA'/D'/AS/'D'/DSnhshbdsjadbjhasasnjdkjas;';';''A'D'AD'A;SDkdnaskdasjkdajsdnjD.SD.S.D>>>DS,D,D,D,D,,<<<k12123asndkanknasjdnsajnkndkAJKNS@@#3212ndjnaskdasdandaksdsddknasdkasdkaskndaskdadkasnkdsdskmsdkmskddknaksnd"

func GenerateRandomBytes(Username, hash string, ID int) []byte {
	var h = sha1.New()
	h.Write([]byte(character + Username + character + hash + strconv.Itoa(ID)))
	var bs = h.Sum(nil)
	return bs
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
