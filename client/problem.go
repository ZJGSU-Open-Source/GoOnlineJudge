package client

import (
	"ojapi/model"

	"fmt"
)

type ProblemService struct {
	apiClient *ApiClient
}

func (s *ProblemService) List() ([]*model.Problem, error) {
	u := "/problems"

	req, err := s.apiClient.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	problems := new([]*model.Problem)
	_, err = s.apiClient.Do(req, problems)
	if err != nil {
		return nil, err
	}

	return *problems, nil
}

// Get a single news.
func (s *ProblemService) Get(id string) (*model.Problem, error) {
	u := fmt.Sprintf("/problems/%s", id)

	req, err := s.apiClient.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	problem := new(model.Problem)
	_, err = s.apiClient.Do(req, problem)
	if err != nil {
		return nil, err
	}

	return problem, nil
}

// Create a new problem
func (s *ProblemService) Create(problem *model.Problem) (*model.Problem, error) {
	u := "/problems"

	req, err := s.apiClient.NewRequest("POST", u, problem)
	if err != nil {
		return nil, err
	}

	_, err = s.apiClient.Do(req, problem)
	if err != nil {
		return nil, err
	}

	return problem, nil
}
