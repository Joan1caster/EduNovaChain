package api

import (
	"fmt"
	"testing"
)

var (
	URL = "http://localhost:8080"
)

func TestAddUser(t *testing.T) {
	strUrl := fmt.Sprintf("%s/user", URL)
	fmt.Println("test...." + strUrl)
	fmt.Println(strUrl)
}
