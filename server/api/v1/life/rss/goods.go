package rss

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/alecthomas/template"
	resp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	htime "github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/rss"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/os/gtime"
)

var (
	goodsEvaluationService = service.ServiceGroupApp.GoodsServiceGroup.GoodsEvaluationService
	goodsLabelService      = service.ServiceGroupApp.GoodsServiceGroup.GoodsLabelService
)

type TableRes struct {
	GoodsName string
	BuyCron   string
	CleanCron string
}

// GoodsRss 根据goods_evaluation直接生成feed
func (r RssApi) GoodsRss(ctx *gin.Context) {
	res := rss.Rss(&rss.Feed{
		Title: rss.Title{
			Prefix: "life",
			Name:   "复购/清洗/更换物品",
		},
		Author:      "lry",
		URL:         GetURL(ctx.Request),
		UpdatedTime: htime.GetToday(),
	}, labelGoods())

	resp.SendXML(ctx, res)
}

// GoodsTableTpl 用来生成表格的模版，GoodsRss只需要生成这个方法的url，就会请求该方法，生成HTML
func (RssApi) GoodsTableTpl(ctx *gin.Context) {
	label := ctx.Query("label")
	tt := ctx.Query("time")
	labelID, _ := strconv.ParseUint(label, 10, 64)
	// 查出数据
	items := []TableRes{}
	_, goodsList, _ := goodsEvaluationService.GetGoodsEvaluationByLabel(uint(labelID))
	specifiedTime := htime.StrToTime(tt, "Y-m-d")
	for _, goodsInfo := range goodsList {
		// isBuy := CheckCronSpecifiedTime(specifiedTime, goodsInfo.BuyCron)
		// isClean := CheckCronSpecifiedTime(specifiedTime, goodsInfo.CleanCron)
		// if isBuy || isClean {
		//
		// }
		table := TableRes{
			GoodsName: fmt.Sprintf("[%s] %s", goodsInfo.GoodsBrand.BrandName, goodsInfo.GoodsName),
			BuyCron:   checkCronToIcon(specifiedTime, goodsInfo.BuyCron),
			CleanCron: checkCronToIcon(specifiedTime, goodsInfo.CleanCron),
		}
		items = append(items, table)
	}
	// 通过表格形式列出
	tpl := template.Must(template.New("").Parse(HTML))
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, items); err != nil {
		log.Fatal(err)
	}
	ct := buf.String()

	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.String(200, ct, nil)
}

// 不同label的物品
//
//	func labelGoods() []*rss.Item {
//		ret := []*rss.Item{}
//		const GoodsLabelHighScore = 4
//		// 获取所有label>4的标签
//		err, labelList, _ := goodsLabelService.GetGoodsLabelInfoListByScore(GoodsLabelHighScore)
//		if err != nil {
//			return nil
//		}
//
//		// 直接根据查询结果生成md文件，每个label一个feed
//		for _, label := range labelList {
//			// 根据label筛选，以及cron字段是否有数据？
//			// cron是否触发？
//			_, goodsList, _ := goodsEvaluationService.GetGoodsEvaluationByLabel(label.ID)
//
//			// 筛掉掉没有cron的label
//			if len(goodsList) != 0 {
//				items := []TableRes{}
//				//如果没有匹配的cron，就直接移除该feed
//				for _, goodsInfo := range goodsList {
//					isBuy := CheckCronNowDefault(goodsInfo.BuyCron)
//					isClean := CheckCronNowDefault(goodsInfo.CleanCron)
//					if isBuy || isClean {
//						table := TableRes{
//							GoodsName: fmt.Sprintf("%s-%s", goodsInfo.GoodsBrand.BrandName, goodsInfo.GoodsName),
//						}
//						items = append(items, table)
//					}
//				}
//
//				// 剔除掉不需要的label，判断是否有数据，如果没有就说明没有匹配到的cron
//				if len(goodsList) > 0 {
//					ct := fmt.Sprintf(IFrame, uint64(label.ID), gtime.Now().Format("Y-m-d"))
//					uuid, _ := gmd5.EncryptString(fmt.Sprintf("%s%s", htime.GetToday().String(), ct))
//					title := fmt.Sprintf("[%s] - %s", gtime.Date(), label.LabelName)
//					ret = append(ret, &rss.Item{
//						Title:       title,
//						Contents:    ct,
//						UpdatedTime: htime.GetToday(),
//						ID:          uuid,
//					})
//				}
//			}
//		}
//		return ret
//	}
func labelGoods() []*rss.Item {
	ret := []*rss.Item{}
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

		// 筛掉掉没有cron的label
		if len(goodsList) != 0 {
			items := []TableRes{}
			// 如果没有匹配的cron，就直接移除该feed
			for _, goodsInfo := range goodsList {
				isBuy := CheckCronNowDefault(goodsInfo.BuyCron)
				isClean := CheckCronNowDefault(goodsInfo.CleanCron)
				if isBuy || isClean {
					table := TableRes{
						GoodsName: fmt.Sprintf("%s-%s", goodsInfo.GoodsBrand.BrandName, goodsInfo.GoodsName),
					}
					items = append(items, table)
				}
			}

			// 剔除掉不需要的label，判断是否有数据，如果没有就说明没有匹配到的cron
			if len(items) > 0 {
				ct := fmt.Sprintf(IFrame, uint64(label.ID), gtime.Now().Format("Y-m-d"))
				uuid, _ := gmd5.EncryptString(fmt.Sprintf("%s%s", htime.GetToday().String(), ct))
				title := fmt.Sprintf("[%s] - %s", gtime.Date(), label.LabelName)
				ret = append(ret, &rss.Item{
					Title:       title,
					Contents:    ct,
					UpdatedTime: htime.GetToday(),
					ID:          uuid,
				})
			}
		}
	}
	return ret
}

// checkCronToIcon 检查cron是否触发，如果触发就返回对应的icon
func checkCronToIcon(tt time.Time, cron string) string {
	// checkCron := CheckCronNowDefault(cron)
	checkCron := CheckCronSpecifiedTime(tt, cron)
	if checkCron {
		// 处理weekly
		if cron != "@weekly" {
			return fmt.Sprintf("⭐#%s", cron)
		}
		// 处理monthly
		// 处理yearly
		return "✅"
	}
	return ""
}

const IFrame = `
<iframe src='https://mali-api.wrss.top/rss/goods/tpl?label=%d&time=%s' frameborder='0' width='640' height='390'></iframe>
`

const HTML = `
<!DOCTYPE html>
<html>
    <head>
        <title></title>
        <style type="text/css">
            /*表格样式*/            
            table {
                width: 90%;
                background: #ccc;
                margin: 10px auto;
                border-collapse: collapse;
                /*border-collapse:collapse合并内外边距
                (去除表格单元格默认的2个像素内外边距*/ 
            }
/*             table tr:nth-child(odd){background: #e3e3e3;} */
            th,td {
                height: 25px;
                line-height: 25px;
                text-align: center;
                border: 1px solid #ccc;
            }       
            th {
                background: #eee;
                font-weight: normal;
            }       
            tr {
                background: #fff;
            }       
            tr:hover {
                background: #cc0;
            }       
            td a {
                color: #06f;
                text-decoration: none;
            }       
            td a:hover {
                color: #06f;
                text-decoration: underline;
            }
        </style>
    </head>
    <body>
        <table>
            <tr>
                <th>物品名称</th>
                <th>复购周期</th>
                <th>清洁周期/更换周期</th>
            </tr>
{{ range . }}
            <tr>
                <td>{{ .GoodsName }}</td>
                <td>{{ .BuyCron }}</td>
                <td>{{ .CleanCron }}</td>
            </tr>
{{ end }}
        </table>
    </body>
</html>
`
