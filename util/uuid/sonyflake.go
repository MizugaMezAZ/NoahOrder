package uuid

import (
	"gorder/model"

	"github.com/bwmarrin/snowflake"
	"github.com/sony/sonyflake"
)

type sonyFlake struct {
	sf *sonyflake.Sonyflake
}

func newSonyFlake() uuid {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (uint16, error) { return 0, nil },
	})

	return &sonyFlake{sf}
}

func (s *sonyFlake) gen() model.UUID {
	// 這邊err 只有在超過使用年限後會報錯 不考慮
	uintid, _ := s.sf.NextID()

	id := snowflake.ID(uintid)

	return model.UUID{
		ID:     id.Int64(),
		Base58: id.Base58(),
	}
}
