package internal

import (
	"fmt"
	"server/msg"
)

type PlayerData struct {
	PlayerID uint   `gorm:"primary_key"`
	UserName string `gorm:"not null"`
	role     msg.RoleData
}

func (playerData *PlayerData) getValue(PlayerID uint) error {
	// mysql := mysql.MysqlDB()

	// err := mysql.Where("PlayerID = ?", playerID).Limit(1).Find(&playerInfo).Error
	// if nil != err {
	// 	fmt.Println(err)
	// 	return fmt.Errorf("get   PlayerBaseInfo id error: %v", err)
	// }

	return nil
}

func (playerData *PlayerData) saveValue() error {
	fmt.Println(" playerData saveValue ===3= === ", playerData.role.RoleName)
	return nil
}
func addValue(playerData PlayerData) error {
	return nil
}
func setStatus(status int) error {
	return nil
}
