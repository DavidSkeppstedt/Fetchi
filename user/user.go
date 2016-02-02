package user

import (
	"fmt"
	"strconv"
)

type User struct {
	AccessToken string `json:"access_token"`
}
type TestUsers struct {
	Users []User `json:"data"`
}

func (this TestUsers) FormatPrint() {
	for i, v := range this.Users {
		fmt.Println("tokens[" + strconv.Itoa(i) + "]=" + v.AccessToken)
	}
}
