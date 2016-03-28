package repositories

import (
	"app/config"
	"app/models"
	"encoding/json"
	riak "github.com/basho/riak-go-client"
	"log"
)

type Repository interface {
	Bucket() string
	Client() *riak.Client
	Model() models.Model
	Find(string) (models.Model, error)
	Save(models.Model) error
}

type RepositoryImpl struct {
	client *riak.Client
}

func (repo *RepositoryImpl) Client() *riak.Client {
	if repo.client == nil {
		opts := &riak.NewClientOptions{
			RemoteAddresses: config.RIAK_ADDRESSES,
		}

		client, err := riak.NewClient(opts)
		if err != nil {
			log.Fatalln(err)
		}

		repo.client = client
	}

	return repo.client
}

func FetchModel(key string, repo Repository) (models.Model, error) {
	client := repo.Client()
	bucket := repo.Bucket()

	cmd, err := riak.NewFetchValueCommandBuilder().
		WithBucket(bucket).
		WithKey(key).
		WithNotFoundOk(true).
		Build()

	if err != nil {
		return nil, err
	}

	if err := client.Execute(cmd); err != nil {
		return nil, err
	}

	fcmd := cmd.(*riak.FetchValueCommand)
	model := repo.Model()
	model.SetKey(key)

	if len(fcmd.Response.Values) > 0 {
		if err := json.Unmarshal(fcmd.Response.Values[0].Value, &model); err != nil {
			return nil, err
		}
	}

	return model, nil
}

func SaveModel(model models.Model, repo Repository) error {
	client := repo.Client()
	bucket := repo.Bucket()
	jval, _ := model.Json()

	obj := &riak.Object{
		Bucket:      bucket,
		Key:         model.Key(),
		ContentType: "application/json",
		Value:       jval,
	}

	cmd, err := riak.NewStoreValueCommandBuilder().
		WithContent(obj).
		WithReturnBody(true).
		Build()

	if err != nil {
		return err
	}

	if err := client.Execute(cmd); err != nil {
		return err
	}

	return nil
}
