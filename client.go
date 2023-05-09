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

func (c Client) Create(ctx context.Context, id string, object interface{}) error {
	_, err := c.getDocRef(id).Create(ctx, object)

	return err
}

func (c Client) GetById(ctx context.Context, id string, toObject *interface{}) error {
	doc, err := c.getDocRef(id).Get(ctx)
	if err != nil {
		return err
	}

	err = doc.DataTo(toObject)
	return err
}

func (c Client) Put(ctx context.Context, id string, object interface{}) error {
	_, err := c.getDocRef(id).Set(ctx, object)
	return err
}

func (c Client) Delete(ctx context.Context, id string) error {
	_, err := c.getDocRef(id).Delete(ctx)
	return err
}

func (c Client) getCollectionRef() *firestore.CollectionRef {
	return c.client.Collection(c.collection)
}

func (c Client) getDocRef(id string) *firestore.DocumentRef {
	return c.getCollectionRef().Doc(id)
}
