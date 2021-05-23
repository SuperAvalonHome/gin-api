package sodexo_service

import (
	"github.com/SuperAvalonHome/gin-api/models"
)

type SodexoUser struct {
	Id         	int
	Uid        	int
	Uuid  		 	string
	Username 		string
	Mobile      int64
	CardInfo		string
	DefaultCard string
	IdSuffix		string
	BindTime		string
	BindStatus	int
}

func (u *SodexoUser) Add() error {
	sodexoUser := map[string]interface{}{
		"id":          	u.Id,
		"uid":          u.Uid,
		"uuid":         u.Uuid,
		"username":     u.Username,
		"mobile":      	u.Mobile,
		"card_info": 		u.CardInfo,
		"default_card": u.DefaultCard,
		"id_suffix":    u.IdSuffix,
		"bind_time":    u.BindTime,
		"bind_status":  u.BindStatus,
	}
	if err := models.AddSodexoUser(sodexoUser); err != nil {
		return err
	}
	return nil
}

func (u *SodexoUser) Get() (*models.SodexoUser, error) {
	var sodexoUser *models.SodexoUser
	sodexoUser, err := models.GetSodexoUser(u.Id)
	if err != nil {
		return nil, err
	}
	return sodexoUser, nil
}
