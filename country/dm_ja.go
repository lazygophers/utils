//go:build (lang_ja || lang_all) && (country_all || country_americas || country_caribbean || country_dm)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDominica.RegisterName(xlanguage.Japanese, "ドミニカ国")
	dataDominica.RegisterOfficialName(xlanguage.Japanese, "ドミニカ国")
	dataDominica.RegisterCapital(xlanguage.Japanese, "ロゾー")
}
