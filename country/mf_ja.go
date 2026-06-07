//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_mf)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintMartin.RegisterName(xlanguage.Japanese, "サン・マルタン（フランス領）")
	dataSaintMartin.RegisterOfficialName(xlanguage.Japanese, "サン・マルタン")
	dataSaintMartin.RegisterCapital(xlanguage.Japanese, "マリゴ")
}
