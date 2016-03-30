package repositories

import (
	"app/config"
	"app/models"
	"encoding/json"
	riak "github.com/basho/riak-go-client"
	//"github.com/l3mncakes/tt-back/config"
	//"github.com/l3mncakes/tt-back/models"
	"log"
)

type RiakRepository interface {
	Repository
	Client() *riak.Client
}

type RiakRepositoryImpl struct {
	RepositoryImpl
	client *riak.Client
}

func (repo *RiakRepositoryImpl) Client() *riak.Client {
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

func (repo *RiakRepositoryImpl) Model() models.Model {
	return repo.model
}

func (repo *RiakRepositoryImpl) Find(key string) error {
	client := repo.Client()
	bucket := repo.model.Bucket()

	cmd, err := riak.NewFetchValueCommandBuilder().
		WithBucket(bucket).
		WithKey(key).
		WithNotFoundOk(true).
		Build()

	if err != nil {
		return err
	}

	if err := client.Execute(cmd); err != nil {
		return err
	}

	fcmd := cmd.(*riak.FetchValueCommand)
	model := repo.model

	if len(fcmd.Response.Values) > 0 {
		if err := json.Unmarshal(fcmd.Response.Values[0].Value, &model); err != nil {
			return err
		}
	}

	repo.model = model

	return nil
}

func (repo *RiakRepositoryImpl) Save() error {
	client := repo.Client()
	bucket := repo.model.Bucket()
	jval, err := json.Marshal(repo.model)

	if err != nil {
		return err
	}

	obj := &riak.Object{
		Bucket:      bucket,
		Key:         repo.model.Key(),
		ContentType: "application/json",
		Value:       jval,
	}

	for k, v := range repo.model.SecondaryIndexes() {
		obj.AddToIndex(k, v)
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

func (repo *RiakRepositoryImpl) FindBySecondaryIndex(indx_name string, indx_val string) error {
	client := repo.Client()
	bucket := repo.model.Bucket()

	cmd, err := riak.NewSecondaryIndexQueryCommandBuilder().
		WithBucket(bucket).
		WithIndexName(indx_name).
		WithIndexKey(indx_val).
		Build()

	if err != nil {
		return err
	}

	if err := client.Execute(cmd); err != nil {
		return err
	}

	scmd := cmd.(*riak.SecondaryIndexQueryCommand)

	if len(scmd.Response.Results) > 0 {
		return repo.Find(string(scmd.Response.Results[0].ObjectKey))
	}

	return nil
}
