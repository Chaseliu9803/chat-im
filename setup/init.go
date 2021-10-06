package setup

import (
	"chat-im/config"
	"chat-im/model"
)

func SetUpServer() {
	config.SetConfig()
	model.SetDB()
}
