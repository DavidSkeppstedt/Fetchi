package user

import (
	"fmt"
)

type User struct {
	AccessToken string `json:"access_token"`
}
type TestUsers struct {
	Users []User `json:"data"`
}

func (this TestUsers) FormatPrint() {
	for _, v := range this.Users {
		fmt.Println(v.AccessToken)
	}
}
