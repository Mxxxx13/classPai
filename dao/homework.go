// @Title : homework
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/29 15:52 

package dao

import "classPai/model"

// UploadHomework 存入数据库
func UploadHomework(homework model.Homework) (err error) {
	if err = DB.Create(&homework).Error; err != nil {
		return
	}
	return
}

// GetHomework 根据id查询
func GetHomework(id int) (homework model.Homework, err error) {
	if err = DB.First(&homework, id).Error; err != nil {
		return
	}
	return homework, nil
}

// AlterHomework 根据id修改
func AlterHomework(id int, homework model.Homework) (err error) {
	if err = DB.Model(&homework).Where("id = ?", id).Updates(model.Homework{
		Name:     homework.Name,
		Deadline: homework.Deadline,
		Content:  homework.Content,
	}).Error; err != nil {
		return
	}
	return
}

// DeleteHomework 根据id删除
func DeleteHomework(id int) (err error) {
	if err = DB.Delete(&model.Homework{}, id).Error; err != nil {
		return
	}
	return
}
