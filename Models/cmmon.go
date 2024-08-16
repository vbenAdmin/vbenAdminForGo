package Models

import (
	"gorm.io/plugin/soft_delete"
)

type Common struct {
	ID      uint                  `form:"id" gorm:"primaryKey" json:"id"`
	Deleted soft_delete.DeletedAt `gorm:"softDelete:;column:deleted;"json:"deleted,omitempty"`
	Updated uint                  `gorm:"column:updated;autoUpdateTime" json:"updated,omitempty"`
	Created uint                  `gorm:"column:created;autoCreateTime" json:"created,omitempty"`
}
