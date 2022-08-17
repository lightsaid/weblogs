package validator

import (
	"fmt"
	"testing"
)

func TestNewJsonValidator(t *testing.T) {

	type SS struct {
		Id   int `json:"id"`
		Name string
	}

	id := SS{Id: 100, Name: "zhangsan"}
	_, err := NewJsonValidator(id)
	if err != nil {
		t.Error(err)
	}
	jvalid, err := NewJsonValidator(&id)
	if err != nil {
		t.Error(err)
	}

	ok := jvalid.Includes("Register", "Register", "Login")
	jvalid.Check(ok, "formType", "formType 选项必须是Register或Login")
	fmt.Println(ok)
	fmt.Println(jvalid.Errors)
	fmt.Println(jvalid.Errors.Get("formType"))

	var name = "sss"
	_, err = NewJsonValidator(name)
	if err == nil {
		t.Errorf("expect err not nil but %s", err)
	}

	_, err = NewJsonValidator(&name)
	if err == nil {
		t.Errorf("expect err not nil but %s", err)
	}

}

func TestEmail(t *testing.T) {
	type SS struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
	}
	id := SS{Id: 100, Email: "zhangsan"}
	jvalid, _ := NewJsonValidator(id)
	jvalid.IsEmail("email")
	if len(jvalid.Errors.Get("email")) == 0 {
		t.Error("邮箱应该是错误的")
	}
	jvalid.Errors["email"] = []string{}
	id.Email = "zhangsan@qq.com"
	jvalid.IsEmail("email")
	if len(jvalid.Errors.Get("email")) == 0 {
		t.Error("邮箱应该正确的")
	}
}
