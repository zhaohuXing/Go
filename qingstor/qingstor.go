package qingstor

import (
	"github.com/yunify/qingstor-sdk-go/config"
	"github.com/yunify/qingstor-sdk-go/service"
)

// GetService returns QingStor's Service.
func GetService() (*service.Service, error) {
	conf, err := config.New("MNNIWFONAWPSKIRYDTMQ",
		"JwIYxVA0ahTDhnQ5NQ96FgkSTGsNlT3hC8Db80Q3")
	if err != nil {
		return nil, err
	}
	return service.Init(conf)
}
