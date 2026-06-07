//go:build country_all || country_americas || country_caribbean || country_mq

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMartinique.RegisterName(xlanguage.Chinese, "马提尼克")
	dataMartinique.RegisterOfficialName(xlanguage.Chinese, "马提尼克")
	dataMartinique.RegisterCapital(xlanguage.Chinese, "法兰西堡")
}
