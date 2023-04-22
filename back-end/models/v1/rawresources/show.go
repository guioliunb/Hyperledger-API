package rawresources

import (
	"errors"

	"github.com/guioliunb/Ledger-API/back-end/models"
)

func Show(id string) (*models.RawResource, error) {

	for index, rawresource := range mockRawResources {
		if rawresource.ID == id {
			return &mockRawResources[index], nil
		}
	}

	return nil, errors.New("unable to find rawresource")
}