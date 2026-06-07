//go:build (lang_ja || lang_all) && (country_all || country_bg || country_eastern_europe || country_europe)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBulgaria.RegisterName(xlanguage.Japanese, "ブルガリア")
	dataBulgaria.RegisterOfficialName(xlanguage.Japanese, "ブルガリア共和国")
	dataBulgaria.RegisterCapital(xlanguage.Japanese, "ソフィア")
}
