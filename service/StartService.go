package service

import "gin-inject/repository"

type StartService struct {
	Repo repository.IStartRepo `inject:""`
}

func (s *StartService) Say(message string) string {
	return s.Repo.Speak(message)
}

func (s *StartService)GetID(ID int)string{
	return s.Repo.Get(ID)
}
