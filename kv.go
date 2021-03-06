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

func (c *KVClient) GetCurrent(key, host string) (value string, err error) {
	return c.Get("current/" + key)
}

func (c *KVClient) PutCurrent(key, host, value string) error {
	return c.Put("current/" + key, value)
}

func (c *KVClient) GetVersion(key string, version uint64) (value string, err error) {
	return "", nil
}

func (c *KVClient) PutVersion(key string, version uint64) error {
	return nil
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
