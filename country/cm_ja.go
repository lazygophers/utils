//go:build (lang_ja || lang_all) && (country_africa || country_all || country_cm || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Japanese, "カメルーン")
	dataCameroon.RegisterOfficialName(xlanguage.Japanese, "カメルーン共和国")
	dataCameroon.RegisterCapital(xlanguage.Japanese, "ヤウンデ")
}
