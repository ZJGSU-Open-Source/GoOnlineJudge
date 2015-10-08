package client

import (
	"ojapi/model"

	"fmt"
)

type ContestService struct {
	apiClient *ApiClient
}

func (s *ContestService) List() ([]*model.Contest, error) {
	u := "/contests"

	req, err := s.apiClient.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	contests := new([]*model.Contest)
	_, err = s.apiClient.Do(req, contests)
	if err != nil {
		return nil, err
	}

	return *contests, nil
}

// Get a single contest.
func (s *ContestService) Get(id string) (*model.Contest, error) {
	u := fmt.Sprintf("/contests/%s", id)

	req, err := s.apiClient.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	contest := new(model.Contest)
	_, err = s.apiClient.Do(req, contest)
	if err != nil {
		return nil, err
	}

	return contest, nil
}

type TinyContest struct {
	Title    string      `json:"title"bson:"title"`
	Encrypt  int         `json:"encrypt"bson:"encrypt"`
	Argument interface{} `json:"argument"bson:"argument"`

	Start int64 `json:"start"bson:"start"`
	End   int64 `json:"end"bson:"end"`

	List []int `json:"list"bson:"list"` //problem list
}

// Create a new contest.
func (s *ContestService) Create(tinyContest *TinyContest) (*model.Contest, error) {
	u := "/contests"

	req, err := s.apiClient.NewRequest("POST", u, tinyContest)
	if err != nil {
		return nil, err
	}

	contest := new(model.Contest)
	_, err = s.apiClient.Do(req, contest)
	if err != nil {
		return nil, err
	}

	return contest, nil
}
