//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCameroon.RegisterName(xlanguage.Japanese, "カメルーン")
	dataCameroon.RegisterOfficialName(xlanguage.Japanese, "カメルーン共和国")
	dataCameroon.RegisterCapital(xlanguage.Japanese, "ヤウンデ")
}
