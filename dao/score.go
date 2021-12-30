// @Title : score
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 16:25 

package dao

import "classPai/model"


// UploadScore
func UploadScore(score model.Score) (err error) {
	if err = DB.Create(&score).Error; err != nil {
		return
	}
	return
}

// ShowScore
func ShowScore(id uint) (scores []model.Score, err error) {
	if err = DB.Find(&scores, id).Error; err != nil {
		return
	}
	return scores, nil
}

// AlterScore
func AlterScore(id int, score model.Score) (err error) {
	if err = DB.Model(&score).Where("id = ?", id).Updates(model.Score{
		Uid: score.Uid,
		Username: score.Username,
		Score:   score.Score,
		Subject: score.Subject,
	}).Error; err != nil {
		return
	}
	return
}

// DeleteScore
func DeleteScore(id int) (err error) {
	if err = DB.Delete(&model.Score{}, id).Error; err != nil {
		return
	}
	return
}
