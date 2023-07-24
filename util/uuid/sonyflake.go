package uuid

import (
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

func (s *sonyFlake) gen() int64 {
	// 這邊err 只有在超過使用年限後會報錯 不考慮
	id, _ := s.sf.NextID()

	return int64(id)
}
