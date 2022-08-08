package trojan

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trojan-dashboard/pkg/common"
	"trojan-dashboard/pkg/configs"
	"trojan-dashboard/pkg/meInfluxdb"
	"trojan-dashboard/pkg/meTrojan"
)

func GetAllTraffic(c *gin.Context) {
	i := meInfluxdb.NewInfluxDBClient(
		configs.InfluxDBConf.Url,
		configs.InfluxDBConf.AuthToken,
		configs.InfluxDBConf.Org,
		configs.InfluxDBConf.Bucket)
	res := meTrojan.QueryAllTraffic(i, common.GetMonthFirstDay())
	c.JSON(http.StatusOK, res)
}
