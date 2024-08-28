package mongo

import "github.com/VishwasMallikarjuna/go-mongo-application/common/config"

func NewHandler(config config.Config) Handler {
	return &theHandler{
		config: config,
	}
}
