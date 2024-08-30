package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var strUrl = "http://localhost:8080"

func TestUserHandler_GetSIWEMessage(t *testing.T) {
	strUrl = fmt.Sprintf(strUrl + "/api/v1/siweMessage?wallet=0x0d0d0d0d0d0d0d0d0d0d00d00d0d0d0d0d0d0d0d")
	resp, err := http.Get(strUrl)
	if err != nil {
		t.Error(err)
	} else {
		defer resp.Body.Close()

		// 读取响应体内容
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("Response Status:", resp.Status)
			fmt.Println("Response Body:", string(body))
		}
	}
}

func TestCreateUser(t *testing.T) {
	strUrl = fmt.Sprintf(strUrl + "/api/v1/createUser")
	resp, err := http.Post(strUrl, "application/json", nil)
	if err != nil {
		t.Error(err)
	} else {
		defer resp.Body.Close()

		// 读取响应体内容
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("Response Status:", resp.Status)
			fmt.Println("Response Body:", string(body))
		}
	}
}
