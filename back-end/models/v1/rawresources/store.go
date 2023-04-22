package rawresources

import (
	"time"

	"github.com/guioliunb/Ledger-API/back-end/models"
)

func Store(name string, typeId string, weight int, arrivalTime *time.Time)(rawresource *models.RawResource, err error){

	rawresource, err = models.NewRawResource(name, typeId, weight, arrivalTime)

	if err != nil{
		return
	}

	mockRawResources = append(mockRawResources, *rawresource)

	return
}