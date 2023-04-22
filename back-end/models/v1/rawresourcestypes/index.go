package rawresourcetypes

import (
	"github.com/guioliunb/Ledger-API/back-end/models"
)

func Index() (rawresourcetypes * models.RawResourceTypes, err error){
	rawresourcetypes = &mockRawResourceTypes	
	return
}