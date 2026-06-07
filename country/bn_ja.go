//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrunei.RegisterName(xlanguage.Japanese, "ブルネイ")
	dataBrunei.RegisterOfficialName(xlanguage.Japanese, "ブルネイ・ダルサラーム国")
	dataBrunei.RegisterCapital(xlanguage.Japanese, "バンダルスリブガワン")
}
