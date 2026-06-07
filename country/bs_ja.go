//go:build (lang_ja || lang_all) && (country_all || country_americas || country_bs || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBahamas.RegisterName(xlanguage.Japanese, "バハマ")
	dataBahamas.RegisterOfficialName(xlanguage.Japanese, "バハマ国")
	dataBahamas.RegisterCapital(xlanguage.Japanese, "ナッソー")
}
