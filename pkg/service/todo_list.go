package service

import (
	todo "github.com/akiyamart/restAPIGo"
	"github.com/akiyamart/restAPIGo/pkg/repository"
)

type TodoListSevice struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListSevice {
	return &TodoListSevice{repo: repo}
}

func (s *TodoListSevice) Create(userId int, list todo.TodoList) (int, error) { 
	return s.repo.Create(userId, list)
}

func (s *TodoListSevice) GetAll(userId int) ([]todo.TodoList, error) { 
	return s.repo.GetAll(userId)
}

func (s *TodoListSevice) GetById(userId, listId int) (todo.TodoList, error) { 
	return s.repo.GetById(userId, listId)
}

func (s *TodoListSevice) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}


func (s *TodoListSevice) Update(userId, listId int, input todo.UpdateListInput) error { 
	if err := input.Validate(); err != nil { 
		return err
	}
	
	return s.repo.Update(userId, listId, input)
}