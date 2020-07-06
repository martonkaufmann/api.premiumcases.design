package printifyapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Product struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Variants    []struct {
		ID        int   `json:"id"`
		Cost      int   `json:"cost"`
		Price     int   `json:"price"`
		IsEnabled bool  `json:"is_enabled"`
		IsDefault bool  `json:"is_default"`
		OptionIDs []int `json:"options"`
	}
	Images []struct {
		RemoteURL  string `json:"src"`
		VariantIDs []int  `json:"variant_ids"`
		IsDefault  bool   `json:"is_default"`
	}
}

func GetProduct(shopID, productID string) (*Product, error) {
	c := new(http.Client)
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"https://api.printify.com/v1/shops/%s/products/%s.json",
			shopID,
			productID,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add(
		"Authorization",
		"Bearer "+os.Getenv("PRINTIFY_API_KEY"),
	)

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	p := new(Product)

	err = json.Unmarshal(body, &p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func ProductPublishSuccess(shopID, productID string) error {
	c := new(http.Client)
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"https://api.printify.com/v1/shops/%s/products/%s/publishing_succeeded.json",
			shopID,
			productID,
		),
		nil,
	)

	if err != nil {
		return err
	}

	req.Header.Add(
		"Authorization",
		"Bearer "+os.Getenv("PRINTIFY_API_KEY"),
	)

	_, err = c.Do(req)

	return err
}
