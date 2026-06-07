//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.Japanese, "モーリシャス")
	dataMauritius.RegisterOfficialName(xlanguage.Japanese, "モーリシャス共和国")
	dataMauritius.RegisterCapital(xlanguage.Japanese, "ポートルイス")
}
