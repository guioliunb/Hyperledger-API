package rawresources

import (
	"github.com/guioliunb/Ledger-API/back-end/models"
)

func Index() (rawresources * models.RawResources, err error){
	rawresources = &mockRawResources

	return
}