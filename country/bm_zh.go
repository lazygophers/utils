//go:build country_all || country_americas || country_bm || country_northern_america

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.Chinese, "百慕大")
	dataBermuda.RegisterOfficialName(xlanguage.Chinese, "百慕大")
	dataBermuda.RegisterCapital(xlanguage.Chinese, "汉密尔顿")
}
