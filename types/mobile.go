package types

import "github.com/Elytrium/elling/elling"

type Mobile struct {
	UserID uint64 `gorm:"primaryKey"`
	Number int64
}

type Request struct {
	Code   string
	Number string
}

func (m *Mobile) Save() {
	elling.DB.Save(m)
}
