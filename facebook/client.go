package facebook

import (
	"encoding/json"
	"github.com/DavidSkeppstedt/fetchi/user"
	"io/ioutil"
	"net/http"
)

var url string = "/accounts/test-users/"
var baseUrl string = "https://graph.facebook.com/v2.5/"

func TestUsers(id, token string) user.TestUsers {
	response, err := http.Get(baseUrl + id + url + "?access_token=" + token)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}
	var testUsers user.TestUsers
	json.Unmarshal(data, &testUsers)
	return testUsers
}
