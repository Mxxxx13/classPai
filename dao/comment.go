// @Title : comment
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/27 16:08 

package dao

import "classPai/model"

func UploadComment(comment model.Comment) (err error) {
	if err = DB.Create(&comment).Error;err != nil {
		return
	}
	return
}

func GetComments(tid int) (comments []model.Comment, err error) {
	if err = DB.Where("tid = ?",tid).Find(&comments).Error;err != nil {
		return
	}
	return
}

func DeleteComment(cid uint) (err error){
	if err = DB.Delete(&model.Comment{}, cid).Error; err != nil {
		return
	}
	return
}