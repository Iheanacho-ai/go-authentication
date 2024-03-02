package initializers

import (
	"github.com/Iheanacho-ai/go-authentication.git/models"
)
func SyncDatabase(){
	DB.AutoMigrate(&models.User{})
}
