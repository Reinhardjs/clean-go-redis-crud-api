package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type DeletePostSuite struct {
	suite.Suite
}

func TestDeletePostPostSuite(t *testing.T) {
	suite.Run(t, new(DeletePostSuite))
}

func (s *DeletePostSuite) TestDeletePostDoesNotExist() {

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest("DELETE", "http://localhost:8080/posts/-1", nil)

	// Fetch Request
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	s.Equal(http.StatusNotFound, resp.StatusCode)
	s.JSONEq(`{
		"detail": "record not found"
	}`, string(body))
}

func (s *DeletePostSuite) TestDeletePost() {

	// Create client
	client := &http.Client{}

	// Create request
	req, _ := http.NewRequest("DELETE", "http://localhost:8080/posts/44", nil)

	// Fetch Request
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	s.Equal(http.StatusOK, resp.StatusCode)
	s.JSONEq(`{
		"status": 200,
		"message": "success",
		"data": {
			"rows_affected": 1
		}
	}`, string(body))
}
