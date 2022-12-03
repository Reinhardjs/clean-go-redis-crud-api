package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetSinglePostSuite struct {
	suite.Suite
}

func TestGetSinglePostSuite(t *testing.T) {
	suite.Run(t, new(GetSinglePostSuite))
}

func (s *GetSinglePostSuite) TestGetPostThatDoesNotExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/posts/1")
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusNotFound, r.StatusCode)
	s.JSONEq(`{
		"detail": "record not found"
	}`, string(body))
}

func (s *GetSinglePostSuite) TestGetPostThatDoesExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/posts/53")
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusOK, r.StatusCode)
	s.JSONEq(`{
		"status": 200,
		"message": "success",
		"data": {
			"id": 53,
			"title": "This is",
			"description": "This is description",
			"created_at": "2022-12-03T19:15:19.923628Z",
			"updated_at": "2022-12-03T19:15:19.923628Z"
		}
	}`, string(body))
}
