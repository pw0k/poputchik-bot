package service

import (
	"fmt"
	"stranger-bot/internal/model"
)

type UserService struct {
	Users map[string]*model.User
	Queue []string
}

func NewUserService() *UserService {
	return &UserService{
		Users: make(map[string]*model.User),
		Queue: make([]string, 0),
	}
}

func (s *UserService) ProcessUser(username string) {
	if _, exists := s.Users[username]; !exists {
		s.Users[username] = &model.User{Username: username, Status: model.Ready}
		fmt.Printf("New user added: %s\n", username)
	} else {
		for i, u := range s.Queue {
			if u == username {
				s.Queue = append(s.Queue[:i], s.Queue[i+1:]...)
				fmt.Printf("User %s removed from queue\n", username)
				return
			}
		}
		s.Queue = append(s.Queue, username)
		fmt.Printf("User %s added to queue\n", username)
	}
}
