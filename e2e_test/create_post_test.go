package e2e_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CreatePostSuite struct {
	suite.Suite
}

func TestCreatePostPostSuite(t *testing.T) {
	suite.Run(t, new(CreatePostSuite))
}

// This will always failed, because cannot predict the timestamps
func (s *CreatePostSuite) TestCreatePost() {

	values := map[string]string{"title": "This is title", "description": "This is description"}
	json_data, _ := json.Marshal(values)

	resp, _ := http.Post("http://localhost:8080/posts", "application/json",
		bytes.NewBuffer(json_data))

	body, _ := ioutil.ReadAll(resp.Body)

	s.Equal(http.StatusOK, resp.StatusCode)
	s.JSONEq(`{
		"status": 201,
		"message": "success",
		"data": {
			"id": 777,
			"title": "This is title",
			"description": "This is description",
			"created_at": "created_at timestamp",
			"updated_at": "updated_at timestamp"
		}
	}`, string(body))
}
