package etcfs

import consulapi "github.com/hashicorp/consul/api"

type KVClient struct {
	client *consulapi.Client
}

func NewKVClient() (*KVClient, error) {
	if client, err := consulapi.NewClient(consulapi.DefaultConfig()); err != nil {
		return nil, err
	} else {
		kvclient := &KVClient{
			client: client,
		}
		return kvclient, nil
	}
}

func (c *KVClient) Put(key, value string) error {
	kv := &consulapi.KVPair{
		Key:   key,
		Value: []byte(value),
	}
	_, err := c.client.KV().Put(kv, nil)
	return err
}

func (c *KVClient) Get(key string) (value string, err error) {
	if kv, _, err := c.client.KV().Get(key, nil); err != nil {
		return "", err
	} else {
		return string(kv.Value), nil
	}
}
