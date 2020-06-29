package hasura

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type query struct {
	Query string `json:"query"`
}

func MakeRequest(q string) error {
	b, err := json.Marshal(query{Query: q})

	if err != nil {
		return err
	}

	c := new(http.Client)
	req, err := http.NewRequest("POST", os.Getenv("HASURA_URL"), bytes.NewBuffer(b))

	if err != nil {
		return err
	}

	req.Header.Add("x-hasura-admin-secret", os.Getenv("HASURA_SECRET"))
	req.Header.Add("Content-Type", "application/json")

	_, err = c.Do(req)

	return err
}
