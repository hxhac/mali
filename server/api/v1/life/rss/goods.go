package rss

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/alecthomas/template"

	resp "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/time"
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

		// 筛掉掉没有cron的label
		if len(goodsList) != 0 {
			items := []TableRes{}

			// TODO 如果没有匹配的cron，就直接移除该feed
			for _, goodsInfo := range goodsList {
				isBuy := CheckCronDefault(goodsInfo.BuyCron)
				isClean := CheckCronDefault(goodsInfo.CleanCron)
				if isBuy || isClean {
					table := TableRes{
						GoodsName: fmt.Sprintf("%s-%s", goodsInfo.GoodsBrand.BrandName, goodsInfo.GoodsName),
						BuyCron:   checkCronToIcon(goodsInfo.BuyCron),
						CleanCron: checkCronToIcon(goodsInfo.CleanCron),
					}
					items = append(items, table)
				}
			}

			// 通过表格形式列出
			tpl := template.Must(template.New("").Parse(HTML))
			var buf bytes.Buffer
			if err = tpl.Execute(&buf, items); err != nil {
				log.Fatal(err)
			}
			ct := buf.String()
			// 判断是否有数据，如果没有就说明没有匹配到的cron
			if len(items) > 0 {
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
	}
	return ret
}

func checkCronToIcon(cron string) string {
	checkCron := CheckCronDefault(cron)
	if checkCron {
		return "✅"
	}
	return ""
}

// 生成html
func GenerateHTML(html string, datas map[string]any) string {
	for key, data := range datas {
		rDataKey := reflect.TypeOf(data)
		rDataVal := reflect.ValueOf(data)
		fieldNum := rDataKey.NumField()
		for i := 0; i < fieldNum; i++ {
			fName := rDataKey.Field(i).Name
			rValue := rDataVal.Field(i)

			var fValue string
			switch rValue.Interface().(type) {
			case string:
				fValue = rValue.String()
			case []string:
				fValue = strings.Join(rValue.Interface().([]string), "<br>")
			}

			mark := fmt.Sprintf("{{%s.%s}}", key, fName)
			html = strings.ReplaceAll(html, mark, fValue)
		}
	}
	return html
}

const HTML = `
<iframe frameborder="0" scrolling="no" frameborder="0" border="0" style= "border:none;">
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
</iframe>
`
