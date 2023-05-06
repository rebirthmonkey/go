package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	baseURL string
}

func (s *UserService) GetUserByID(id int) (*User, error) {
	url := s.baseURL + "/users/" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	user := &User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
