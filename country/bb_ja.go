//go:build (lang_ja || lang_all) && (country_all || country_americas || country_bb || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBarbados.RegisterName(xlanguage.Japanese, "バルバドス")
	dataBarbados.RegisterOfficialName(xlanguage.Japanese, "バルバドス")
	dataBarbados.RegisterCapital(xlanguage.Japanese, "ブリッジタウン")
}
