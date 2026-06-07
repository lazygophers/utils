//go:build country_all || country_ba || country_europe || country_southern_europe

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBosniaAndHerzegovina.RegisterName(xlanguage.Chinese, "波斯尼亚和黑塞哥维那")
	dataBosniaAndHerzegovina.RegisterOfficialName(xlanguage.Chinese, "波斯尼亚和黑塞哥维那")
	dataBosniaAndHerzegovina.RegisterCapital(xlanguage.Chinese, "萨拉热窝")
}
