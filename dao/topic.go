// @Title : topic
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/27 15:48 

package dao

import "classPai/model"

// UploadTopic 将Topic存入数据库
func UploadTopic(topic model.Topic) (err error) {
	if err = DB.Create(&topic).Error; err != nil {
		return
	}
	return
}

// LikeTopic 根据Topic id查询like, like+1后更新数据库
func LikeTopic(id int) (err error) {
	var topic model.Topic
	if err = DB.Where("id = ?", id).First(&topic, id).Error; err != nil {
		return
	}

	// 点赞数+1
	topic.Likes++
	if err = DB.Save(&topic).Error; err != nil {
		return
	}

	return
}

// ShowTopic 根据id查询blog
func ShowTopic(id int) (topic model.Topic, err error) {
	if err = DB.First(&topic, id).Error; err != nil {
		return
	}
	return topic, nil
}

// AlterTopic
func AlterTopic(id int, topic model.Topic) (err error) {
	if err = DB.Model(&topic).Where("id = ?", id).Updates(&model.Topic{
		Title:   topic.Title,
		Content: topic.Content,
	}).Error; err != nil {
		return
	}
	return
}

// DeleteTopic 根据id删除blog
func DeleteTopic(id int) (err error) {
	if err = DB.Delete(&model.Topic{}, id).Error; err != nil {
		return
	}
	return
}

//func GetTopic(uid uint) (topics []model.Topic, err error) {
//	if err = DB.Where("uid = ?", uid).Find(&topics).Error; err != nil {
//		return
//	}
//	return
//}