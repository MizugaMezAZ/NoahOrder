package uuid

import (
	"github.com/bwmarrin/snowflake"
)

type snowFlake struct {
	sf *snowflake.Node
}

func newSnowFlake() uuid {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil
	}

	return &snowFlake{node}
}

func (s *snowFlake) gen() int64 {
	id := s.sf.Generate()

	return id.Int64()
}
