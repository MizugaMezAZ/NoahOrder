package uuid

import "gorder/model"

type uuid interface {
	gen() int64
}

var defaultUUID uuid

func init() {
	defaultUUID = newSnowFlake()
}

func GenUUID() model.UUID {
	id := defaultUUID.gen()

	return model.UUID{
		ID:     id,
		Base58: base58(id),
	}
}
