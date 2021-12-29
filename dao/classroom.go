// @Title : classroom'
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/29 14:23 

package dao

import "classPai/model"

// CreateClassroom
func CreateClassroom(classroom model.Classroom) (err error) {
	if err = DB.Create(&classroom).Error; err != nil {
		return err
	}
	return
}

// DeleteClassroom
func DeleteClassroom(tid uint) (err error) {
	if err = DB.Delete(&model.Classroom{}, tid).Error; err != nil {
		return
	}
	return
}
