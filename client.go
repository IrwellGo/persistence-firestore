package persistence_firestore

import (
	"cloud.google.com/go/firestore"
	"context"
)

type Client struct {
	client     *firestore.Client
	collection string
}

func FakeFuncForUnitTesting(pass bool) bool {
	return pass
}

func New(client *firestore.Client, collection string) Client {
	return Client{
		client,
		collection,
	}
}

func NewSimplified(ctx context.Context, projectId string, collection string) (Client, error) {
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		return Client{}, err
	}

	return New(client, collection), nil
}

func (r Client) Create(ctx context.Context, id string, object interface{}) error {
	_, err := r.getDocRef(id).Create(ctx, object)

	return err
}

func (r Client) getCollectionRef() *firestore.CollectionRef {
	return r.client.Collection(r.collection)
}

func (r Client) getDocRef(id string) *firestore.DocumentRef {
	return r.getCollectionRef().Doc(id)
}
