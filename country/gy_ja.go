//go:build (lang_ja || lang_all) && (country_all || country_americas || country_gy || country_south_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataGuyana.RegisterName(xlanguage.Japanese, "ガイアナ")
	dataGuyana.RegisterOfficialName(xlanguage.Japanese, "ガイアナ協同共和国")
	dataGuyana.RegisterCapital(xlanguage.Japanese, "ジョージタウン")
}
