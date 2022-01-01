// @Title : score
// @Description ://TODO: Add Description
// @Author : MX
// @Update : 2021/12/30 16:39 

package service

import (
	"errors"
	"log"
	"strconv"

	"classPai/dao"
	"classPai/model"
	"github.com/gin-gonic/gin"
)

func UploadScore(c *gin.Context) (err error) {
	var score model.Score

	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid is not exist")
	}
	user, err := dao.GetUser(uid.(uint))
	if err != nil {
		return
	}
	// 判断用户身份是否为教师
	if user.Role != 1 {
		return errors.New("权限不够")
	}

	score.Username = c.PostForm("username")
	score.Uid, err = dao.GetUid(score.Username)
	score.Subject = c.PostForm("subject")

	scStr := c.PostForm("score")
	sc ,err := strconv.ParseFloat(scStr,64)
	if err != nil {
		return
	}
	score.Score = sc

	err = dao.UploadScore(score)
	return err
}

func ShowScore(c *gin.Context) (scoreResps []model.ScoreResp, err error) {
	uid, exists := c.Get("uid")
	if !exists {
		return nil, errors.New("uid is not exist")
	}

	scores,err := dao.ShowScore(uid.(uint))
	if err != nil {
		return
	}

	for _, score := range scores {
		scoreResp := model.ScoreResp{
			Username: score.Username,
			Score:    score.Score,
			Subject:  score.Subject,
		}
		scoreResps  = append(scoreResps,scoreResp)
	}
	return
}

func AlterScore(c *gin.Context) (err error) {
	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid is not exist")
	}
	user, err := dao.GetUser(uid.(uint))
	if err != nil {
		return
	}
	// 判断用户身份是否为教师
	if user.Role != 1 {
		return errors.New("权限不够")
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}

	var score model.Score
	score.Subject = c.PostForm("subject")
	scStr := c.PostForm("score")
	sc ,err := strconv.ParseFloat(scStr,64)
	if err != nil {
		return
	}
	score.Score = sc

	err = dao.AlterScore(id, score)
	return err
}

func DeleteScore(c *gin.Context) (err error) {
	uid, exists := c.Get("uid")
	if !exists {
		return errors.New("uid is not exist")
	}
	user, err := dao.GetUser(uid.(uint))
	if err != nil {
		return
	}
	// 判断用户身份是否为教师
	if user.Role != 1 {
		return errors.New("权限不够")
	}

	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		log.Printf("atoi err:%v\n", err)
		return
	}
	err = dao.DeleteScore(id)
	return err
}
