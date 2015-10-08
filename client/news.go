package client

import (
	"ojapi/model"

	"fmt"
)

type NewsService struct {
	apiClient *ApiClient
}

func (s *NewsService) List() ([]*model.News, error) {
	u := "/news"

	req, err := s.apiClient.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	news := new([]*model.News)
	_, err = s.apiClient.Do(req, news)
	if err != nil {
		return nil, err
	}

	return *news, nil
}

// Get a single news.
func (s *NewsService) Get(id string) (*model.News, error) {
	u := fmt.Sprintf("/news/%s", id)

	req, err := s.apiClient.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	news := new(model.News)
	_, err = s.apiClient.Do(req, news)
	if err != nil {
		return nil, err
	}

	return news, nil
}
