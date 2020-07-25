package z

import (
	"strings"
	"time"

	"github.com/vhaoran/yi/service"
	"github.com/vhaoran/yi/z/model"
)

//----------------------------------------------------
// auth: whr  date:2020/7/2216:28--------------------------
// ####请勿擅改此功能代码####
// 用途：起運計算
//---------------------------------------------
type (
	QiYunGet struct {
	}
)

//起动的年、月数
func (r *QiYunGet) Call(z *model.SiZhuModel) (nianShu, yueShu int) {
	asc := r.IsAsc(z)
	//
	//得到第月标记性的12个节
	jie := service.GetJie(z.Nian)

	//得到节 意义上的月
	//-1时表示在上一年
	yueOfJie := r.locateYue(jie, z.Solar())

	var data *service.JieData
	//顺排
	if asc {
		data := jie[z.Yue]
		//12月向前数日期，再用30减
		if yueOfJie == 12 {
			data := jie[len(jie)-1]
			offset := int(z.Solar().Sub(data.Date).Hours() / 24)
			if offset < 0 {
				offset = -offset
			}

			//腊月共30天
			nianShu = (30 - offset + 1) / 3
			yueShu = (30 - offset + 1) % 3
			return
		}

		if yueOfJie == -1 {
			data = jie[0]
		} else { //如：1月，取索引为2的节
			data = jie[z.Yue]
		}

		offset := int(data.Date.Sub(z.Solar()).Hours() / 24)
		//腊月共30天
		nianShu = (offset + 1) / 3
		yueShu = (offset + 1) % 3 * 4
		return
	}

	//--------倒排，计算到上一个节的天数 -----------------------------
	//算到上一年的小寒的时间（因为一节30天，故计算到立春用30减）
	if yueOfJie == -1 {
		data := jie[0]
		offset := int(z.Solar().Sub(data.Date).Hours() / 24)
		if offset < 0 {
			offset = -offset
		}

		//腊月共30天
		nianShu = (offset + 1) / 3
		yueShu = (offset + 1) % 3 * 4
		return
	}

	//如：1月，取索引为2的节
	data = jie[yueOfJie-1]
	offset := int(data.Date.Sub(z.Solar()).Hours() / 24)
	if offset < 0 {
		offset = -offset
	}

	//腊月共30天
	nianShu = (offset + 1) / 3
	yueShu = (offset + 1) % 3 * 4
	return
}

func (r *QiYunGet) locateYue(l []*service.JieData, t time.Time) (month int) {
	//排好序的24节气
	for i, v := range l {
		if i < len(l)-1 {
			if t.After(v.Date) && t.Before(l[i+1].Date) {
				return i + 1
			}
		}
	}

	//------if not found------------
	if t.After(l[len(l)-1].Date) {
		return 12
	}

	//
	return -1
}

//是：大运要顺排,阳男或阴女--,
func (r *QiYunGet) IsAsc(z *model.SiZhuModel) bool {
	gan := z.NianGan
	//甲乙丙丁戊己庚辛壬癸
	//"子丑寅卯辰巳午未申酉戌亥"}
	//子寅辰午申戌
	const yangGan = "甲丙戊庚壬"
	isYanGan := strings.Contains(yangGan, gan)
	//阳男 或阴女 反加真
	return (isYanGan && z.IsMale) || (!isYanGan && !z.IsMale)
}
