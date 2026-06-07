//go:build country_all || country_ck || country_oceania || country_polynesia

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCookIslands.RegisterName(xlanguage.Chinese, "库克群岛")
	dataCookIslands.RegisterOfficialName(xlanguage.Chinese, "库克群岛")
	dataCookIslands.RegisterCapital(xlanguage.Chinese, "阿瓦鲁阿")
}
