package utils

import "testing"

func TestSensitive(t*testing.T) {
	sm := NewSensitiveMarshal("password")
	t.Run("nil", func(t *testing.T) {
		type A struct{}
		t.Log(sm.SafeMarshal(nil))

		var a *A
		t.Log(sm.SafeMarshal(a))
	})

	t.Run("list", func(t *testing.T) {
		t.Log(sm.SafeMarshal([]int{1,2,3}))
	})

	t.Run("struct", func(t *testing.T) {
		type A struct {
			Password int `json:"password"`
		}

		type User struct {
			Name string `json:"name"`
			Password string `json:"password"`
			Age int
			A A
			APtr *A
			AaPtr *A
		}

		user := User{
			Name: "Jack",
			Password: "123456",
			Age: 18,
			A: A{
				Password: 1,
			},
			APtr: &A{
				Password: 2,
			},
			AaPtr: nil,
		}

		t.Log(sm.SafeMarshal(user))

		t.Log(sm.SafeMarshal(&user))
	})
}