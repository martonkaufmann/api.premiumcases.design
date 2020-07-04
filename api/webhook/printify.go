package webhook

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"api.premiumcases.design/pkg/hasura"
	"api.premiumcases.design/pkg/printifyapi"
	"api.premiumcases.design/pkg/utils"
	"github.com/bugsnag/bugsnag-go"
	"github.com/labstack/echo/v4"
	"github.com/thoas/go-funk"
)

func PrintifyProductPublish(ctx echo.Context) error {
	req := new(struct {
		Resource struct {
			ID   string `json:"id" validate:"required"`
			Data struct {
				ShopID int `json:"shop_id" validate:"required"`
			} `json:"data" validate:"required"`
		} `json:"resource" validate:"required"`
	})

	if err := ctx.Bind(req); err != nil {
		bugsnag.Notify(err)

		return ctx.NoContent(http.StatusUnprocessableEntity)
	}

	if err := ctx.Validate(req); err != nil {
		bugsnag.Notify(err)

		return ctx.NoContent(http.StatusUnprocessableEntity)
	}

	p, err := printifyapi.GetProduct(strconv.Itoa(req.Resource.Data.ShopID), req.Resource.ID)

	if err != nil {
		bugsnag.Notify(err)

		return ctx.NoContent(http.StatusInternalServerError)
	}

	c, err := hasuraCaseFromPrintifyProduct(*p)

	if err != nil {
		bugsnag.Notify(err)

		return ctx.NoContent(http.StatusInternalServerError)
	}

	if err := c.Save(); err != nil {
		bugsnag.Notify(err)

		return ctx.NoContent(http.StatusInternalServerError)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func hasuraCaseFromPrintifyProduct(p printifyapi.Product) (*hasura.Case, error) {
	c := new(hasura.Case)
	c.Name = p.Title
	c.Cost = p.Variants[0].Cost
	c.Price = p.Variants[0].Price

	for _, v := range p.Variants {
		if !v.IsEnabled {
			continue
		}

		isSurfaceGlossy := funk.ContainsInt(v.OptionIDs, printifyapi.SurfaceGlossyID)
		printifyDeviceID, isSupportedDevice := funk.FindInt(v.OptionIDs, func(id int) bool {
			return funk.ContainsInt(printifyapi.DeviceIDs, id)
		})

		if !isSurfaceGlossy || !isSupportedDevice {
			continue
		}

		d := new(hasura.Device)
		hasuraDeviceID, err := getHasuraDeviceIDForPrintifyDeviceID(printifyDeviceID)

		if err != nil {
			return nil, err
		}

		d.ID = hasuraDeviceID

		for _, i := range p.Images {
			if i.VariantIDs[0] == v.ID {
				remoteURL := strings.Replace(i.RemoteURL, "-api", "", -1)

				fn, err := utils.DownloadRemoteFile(os.Getenv("IMAGE_PATH"), remoteURL)

				if err != nil {
					return nil, err
				}

				d.Image = fn
			}
		}

		c.Devices = append(c.Devices, *d)
	}

	return c, nil
}

func getHasuraDeviceIDForPrintifyDeviceID(printifyDeviceID int) (int, error) {
	switch printifyDeviceID {
	case printifyapi.DeviceiPhone11ID:
		return hasura.DeviceiPhone11ID, nil
	case printifyapi.DeviceiPhone11ProID:
		return hasura.DeviceiPhone11ProID, nil
	case printifyapi.DeviceiPhone11ProMaxID:
		return hasura.DeviceiPhone11ProMaxID, nil
	case printifyapi.DeviceiPhone8ID:
		return hasura.DeviceiPhone8ID, nil
	case printifyapi.DeviceiPhone8PlusID:
		return hasura.DeviceiPhone8PlusID, nil
	case printifyapi.DeviceiPhoneXID:
		return hasura.DeviceiPhoneXID, nil
	case printifyapi.DeviceiPhoneXRID:
		return hasura.DeviceiPhoneXRID, nil
	case printifyapi.DeviceiPhoneXSID:
		return hasura.DeviceiPhoneXSID, nil
	case printifyapi.DeviceiPhoneXSMaxID:
		return hasura.DeviceiPhoneXSMaxID, nil
	}

	return 0, errors.New("Unsupported device")
}
