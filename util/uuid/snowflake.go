package uuid

import (
	"gorder/model"

	"github.com/bwmarrin/snowflake"
)

type snowFlake struct {
	sf *snowflake.Node
}

func NewSnowFlake() uuid {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil
	}

	return &snowFlake{node}
}

func (s *snowFlake) Gen() model.UUID {
	id := s.sf.Generate()

	return model.UUID{
		ID:     id.Int64(),
		Base58: id.Base58(),
	}
}
