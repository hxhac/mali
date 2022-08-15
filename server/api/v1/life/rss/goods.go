package rss

import (
	"fmt"
	"github.com/apcera/termtables"
	resp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/os/gtime"
	"github.com/golang-module/carbon"
)

var goodsEvaluationService = service.ServiceGroupApp.GoodsServiceGroup.GoodsEvaluationService
var goodsLabelService = service.ServiceGroupApp.GoodsServiceGroup.GoodsLabelService

// GoodsRss 根据goods_evaluation直接生成feed
func (r RssApi) GoodsRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "复购/清洗/更换物品",
		},
		Author:      "lry",
		URL:         GetURL(ctx.Request),
		UpdatedTime: time.GetToday(),
	}, labelGoods())

	resp.SendXML(ctx, res)
}

// 不同label的物品
func labelGoods() []rss.Item {
	ret := []rss.Item{}
	const GoodsLabelHighScore = 4
	// 获取所有label>4的标签
	err, labelList, _ := goodsLabelService.GetGoodsLabelInfoListByScore(GoodsLabelHighScore)
	if err != nil {
		return nil
	}

	// 直接根据查询结果生成md文件，每个label一个feed
	for _, label := range labelList {
		// 根据label筛选，以及cron字段是否有数据？
		// cron是否触发？
		_, goodsList, _ := goodsEvaluationService.GetGoodsEvaluationByLabel(label.ID)

		// 筛掉所有没有cron的label
		if len(goodsList) != 0 {
			t := termtables.CreateTable()
			t.AddHeaders("物品名称", "复购周期", "清洁周期/更换周期")
			for _, goodsInfo := range goodsList {
				t.AddRow(goodsInfo.GoodsName, checkCronRes(goodsInfo.BuyCron), checkCronRes(goodsInfo.CleanCron))
			}
			t.SetModeHTML()

			// 通过表格形式列出
			ct := t.Render()
			uuid, _ := gmd5.EncryptString(fmt.Sprintf("%s%s", time.GetToday().String(), ct))
			title := fmt.Sprintf("[%s] - %s", gtime.Date(), label.LabelName)
			ret = append(ret, rss.Item{
				Title:       title,
				Contents:    ct,
				UpdatedTime: time.GetToday(),
				ID:          uuid,
			})
		}
	}
	return ret
}

func checkCronRes(cron string) string {
	checkCron := CheckCron(cron, carbon.Now().IsSaturday())
	if checkCron {
		return "✅"
	}
	return ""
}
