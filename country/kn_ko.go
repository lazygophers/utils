//go:build (lang_ko || lang_all) && (country_all || country_americas || country_caribbean || country_kn)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSaintKittsAndNevis.RegisterName(xlanguage.Korean, "세인트키츠 네비스")
	dataSaintKittsAndNevis.RegisterOfficialName(xlanguage.Korean, "세인트키츠 네비스 연방")
	dataSaintKittsAndNevis.RegisterCapital(xlanguage.Korean, "바스테르")
}
