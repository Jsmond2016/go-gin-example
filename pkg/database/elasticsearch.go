package database

import (
	"github.com/olivere/elastic/v7"

	"github.com/EDDYCJY/go-gin-example/pkg/config"
)

func NewElasticsearch(cfg *config.ElasticsearchConfig) (*elastic.Client, error) {
	return elastic.NewClient(
		elastic.SetURL(cfg.URL),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(cfg.Username, cfg.Password),
	)
}