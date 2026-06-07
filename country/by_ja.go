//go:build (lang_ja || lang_all) && (country_all || country_by || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBelarus.RegisterName(xlanguage.Japanese, "ベラルーシ")
	dataBelarus.RegisterOfficialName(xlanguage.Japanese, "ベラルーシ共和国")
	dataBelarus.RegisterCapital(xlanguage.Japanese, "ミンスク")
}
