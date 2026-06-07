//go:build country_africa || country_all || country_eastern_africa || country_zw

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Chinese, "津巴布韦")
	dataZimbabwe.RegisterOfficialName(xlanguage.Chinese, "津巴布韦共和国")
	dataZimbabwe.RegisterCapital(xlanguage.Chinese, "哈拉雷")
}
