//go:build (lang_ja || lang_all) && (country_all || country_eastern_europe || country_europe || country_md)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMoldova.RegisterName(xlanguage.Japanese, "モルドバ")
	dataMoldova.RegisterOfficialName(xlanguage.Japanese, "モルドバ共和国")
	dataMoldova.RegisterCapital(xlanguage.Japanese, "キシナウ")
}
