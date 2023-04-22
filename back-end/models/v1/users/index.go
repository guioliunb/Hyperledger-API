package users

import (
	"github.com/guioliunb/Ledger-API/back-end/models"
)

func Index() (users * models.Users, err error){
	users = &mockUsers

	return
}