package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GetAllPostSuite struct {
	suite.Suite
}

func TestGetAllPostSuite(t *testing.T) {
	suite.Run(t, new(GetAllPostSuite))
}

func (s *GetAllPostSuite) TestGetAllPostThatDoesExist() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/posts")
	body, _ := ioutil.ReadAll(r.Body)

	s.Equal(http.StatusOK, r.StatusCode)
	s.JSONEq(`{
		"status": 200,
		"message": "success",
		"data": [
			{
				"id": 49,
				"title": "This is",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T17:38:14.739374Z"
			},
			{
				"id": 35,
				"title": "This is title",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T16:29:12.582286Z"
			},
			{
				"id": 46,
				"title": "This is title 2",
				"description": "This is description 2",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T19:00:57.78321Z"
			},
			{
				"id": 42,
				"title": "This is updated title 5",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T19:01:06.346782Z"
			},
			{
				"id": 43,
				"title": "This is updated title 5",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T19:03:28.097134Z"
			},
			{
				"id": 52,
				"title": "This is",
				"description": "This is description",
				"created_at": "2022-12-03T19:15:05.704656Z",
				"updated_at": "2022-12-03T19:15:05.704656Z"
			},
			{
				"id": 53,
				"title": "This is",
				"description": "This is description",
				"created_at": "2022-12-03T19:15:19.923628Z",
				"updated_at": "2022-12-03T19:15:19.923628Z"
			},
			{
				"id": 41,
				"title": "This is title",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T16:41:56.866906Z"
			},
			{
				"id": 39,
				"title": "This is title",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T16:53:37.766581Z"
			},
			{
				"id": 44,
				"title": "This is",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T17:37:38.873348Z"
			},
			{
				"id": 47,
				"title": "This is",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T17:38:02.103052Z"
			},
			{
				"id": 48,
				"title": "This is",
				"description": "This is description",
				"created_at": "0001-01-01T00:00:00Z",
				"updated_at": "2022-12-03T17:38:03.114936Z"
			}
		]
	}`, string(body))
}
