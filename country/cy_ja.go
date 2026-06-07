//go:build (lang_ja || lang_all) && (country_all || country_cy || country_europe || country_western_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCyprus.RegisterName(xlanguage.Japanese, "キプロス")
	dataCyprus.RegisterOfficialName(xlanguage.Japanese, "キプロス共和国")
	dataCyprus.RegisterCapital(xlanguage.Japanese, "ニコシア")
}
