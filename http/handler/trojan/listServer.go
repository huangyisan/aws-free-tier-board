package trojan

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trojan-dashboard/pkg/configs"
)

// 列出所有server信息
func ListServer(c *gin.Context) {
	sl := serverList{}
	tjc := configs.TrojanConf
	for _, v := range *tjc {
		if _, ok := sl[v.Group]; !ok {
			sl[v.Group] = append(sl[v.Group], server{Tag: v.Tag, Ip: v.Ip})
		}
	}
	c.JSON(http.StatusOK, sl)
}

// 列出所有server tag
func ListServerTag(c *gin.Context) {
	tags := make([]string, 0)
	tjc := configs.TrojanConf
	for _, v := range *tjc {
		tags = append(tags, v.Tag)
	}
	c.JSON(http.StatusOK, tags)
}

// 列出所有server group
func ListServerGroup(c *gin.Context) {
	gp := make([]string, 0)
	tjc := configs.TrojanConf
	for _, v := range *tjc {
		gp = append(gp, v.Group)
	}
	c.JSON(http.StatusOK, gp)
}

// 根据给定group信息列出该group中所有tag
func ListServerByGroup(c *gin.Context) {
	tags := make([]string, 0)
	tjc := configs.TrojanConf
	for _, v := range *tjc {
		if v.Group == c.Param("group") {
			tags = append(tags, v.Tag)
		}
	}
	c.JSON(http.StatusOK, tags)
}

type server struct {
	Tag string `json:"tag"`
	Ip  string `json:"ip"`
}

type serverList map[string][]server
