//go:build country_africa || country_all || country_cd || country_middle_africa

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.Chinese, "刚果民主共和国")
	dataDrCongo.RegisterOfficialName(xlanguage.Chinese, "刚果民主共和国")
	dataDrCongo.RegisterCapital(xlanguage.Chinese, "金沙萨")
}
