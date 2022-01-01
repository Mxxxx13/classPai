// @Title : sign
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 12:19 

package dao

import (
	"fmt"

	"classPai/model"
)

// UploadSign
func UploadSign(sign model.Sign) (err error) {
	if err = DB.Create(&sign).Error; err != nil {
		return
	}
	return
}

// Sign
func Sign(table model.SignTable) (err error) {
	if err = DB.Create(&table).Error; err != nil {
		return
	}
	return
}

// ShowSign
func ShowSign(sid uint) (signs []model.SignTable,err error) {
	if err = DB.Where("sid = ?",sid).Find(&signs).Error; err != nil {
		return
	}
	fmt.Println(signs)
	return
}