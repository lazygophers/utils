//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataWesternSahara.RegisterName(xlanguage.Japanese, "西サハラ")
	dataWesternSahara.RegisterOfficialName(xlanguage.Japanese, "サハラ・アラブ民主共和国")
	dataWesternSahara.RegisterCapital(xlanguage.Japanese, "ラユーン")
}
