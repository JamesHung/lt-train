package main

import (
	"encoding/json"
	"fmt"
)

type APIUser struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
}

func parseUser(jsonStr string) (*APIUser, error) {
	var user APIUser
	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		return nil, err
	}
	fmt.Printf("User: %+v\n", user)

	data, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	fmt.Printf("json str: %s\n", string(data))

	return &user, nil
}

func main() {
	data := `{"id":"123"}`

	user, err := parseUser(data)
	if err != nil {
		fmt.Printf("failed to parse user: %v\n", err)
		return
	}

	fmt.Printf("got %+v\n", user)
}
