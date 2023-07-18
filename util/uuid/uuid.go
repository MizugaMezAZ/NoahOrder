package uuid

import "gorder/model"

type uuid interface {
	gen() model.UUID
}

var defaultUUID uuid

func init() {
	defaultUUID = newSnowFlake()
}

func GenUUID() model.UUID {
	return defaultUUID.gen()
}
