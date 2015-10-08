package client

import (
	"ojapi/model"

	"fmt"
)

type SolutioNservice struct {
	apiClient *ApiClient
}

func (s *SolutioNservice) List() ([]*model.Solution, error) {
	u := "/solutions"

	req, err := s.apiClient.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	solutions := new([]*model.Solution)
	_, err = s.apiClient.Do(req, solutions)
	if err != nil {
		return nil, err
	}

	return *solutions, nil
}

// Get a single solution.
func (s *SolutioNservice) Get(id string) (*model.Solution, error) {
	u := fmt.Sprintf("/solutions/%s", id)

	req, err := s.apiClient.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	solution := new(model.Solution)
	_, err = s.apiClient.Do(req, solution)
	if err != nil {
		return nil, err
	}

	return solution, nil
}

// Create a new solution
func (s *SolutioNservice) Create(code string, compilerId int) (*model.Solution, error) {
	u := "/solutions"
	out := struct {
		Code       string `json:"code"`
		CompilerId int    `json:"compiler_id"`
	}{
		Code:       code,
		CompilerId: compilerId,
	}

	req, err := s.apiClient.NewRequest("POST", u, out)
	if err != nil {
		return nil, err
	}

	solution := new(model.Solution)
	_, err = s.apiClient.Do(req, solution)
	if err != nil {
		return nil, err
	}

	return solution, nil
}
