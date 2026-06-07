//go:build country_africa || country_all || country_cg || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCongo.RegisterName(xlanguage.Chinese, "刚果共和国")
	dataCongo.RegisterOfficialName(xlanguage.Chinese, "刚果共和国")
	dataCongo.RegisterCapital(xlanguage.Chinese, "布拉柴维尔")
}
