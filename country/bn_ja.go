//go:build (lang_ja || lang_all) && (country_all || country_asia || country_bn || country_south_eastern_asia)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.Japanese, "ブルネイ")
	dataBrunei.RegisterOfficialName(xlanguage.Japanese, "ブルネイ・ダルサラーム国")
	dataBrunei.RegisterCapital(xlanguage.Japanese, "バンダルスリブガワン")
}
