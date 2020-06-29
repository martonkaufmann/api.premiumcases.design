package hasura

import (
	"fmt"
)

type Case struct {
	Name    string
	Price   int
	Cost    int
	Image   string
	Devices []Device
}

func (c *Case) Save() error {
	dq := ""

	for _, d := range c.Devices {
		dq += fmt.Sprintf("{image: \"%s\" device_id: %d}", d.Image, d.ID)
	}

	err := MakeRequest(fmt.Sprintf(
		"mutation{insert_cases_one(object:{name: \"%s\" price: %d cost: %d cases_devices:{data: [%s]}}) {id}}",
		c.Name,
		c.Price,
		c.Cost,
		dq,
	))

	return err
}
