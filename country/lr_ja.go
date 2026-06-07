//go:build (lang_ja || lang_all) && (country_africa || country_all || country_lr || country_western_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataLiberia.RegisterName(xlanguage.Japanese, "リベリア")
	dataLiberia.RegisterOfficialName(xlanguage.Japanese, "リベリア共和国")
	dataLiberia.RegisterCapital(xlanguage.Japanese, "モンロビア")
}
