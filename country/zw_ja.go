//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eastern_africa || country_zw)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataZimbabwe.RegisterName(xlanguage.Japanese, "ジンバブエ")
	dataZimbabwe.RegisterOfficialName(xlanguage.Japanese, "ジンバブエ共和国")
	dataZimbabwe.RegisterCapital(xlanguage.Japanese, "ハラレ")
}
