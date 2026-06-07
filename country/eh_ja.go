//go:build (lang_ja || lang_all) && (country_africa || country_all || country_eh || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.Japanese, "西サハラ")
	dataWesternSahara.RegisterOfficialName(xlanguage.Japanese, "サハラ・アラブ民主共和国")
	dataWesternSahara.RegisterCapital(xlanguage.Japanese, "ラユーン")
}
