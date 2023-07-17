package uuid

import "gorder/model"

type uuid interface {
	Gen() model.UUID
}

var defaultUUID uuid

func init() {
	defaultUUID = NewSnowFlake()
}

func GenUUID() model.UUID {
	return defaultUUID.Gen()
}
