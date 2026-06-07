//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAfghanistan.RegisterName(xlanguage.Japanese, "アフガニスタン")
	dataAfghanistan.RegisterOfficialName(xlanguage.Japanese, "アフガニスタン・イスラム首長国")
	dataAfghanistan.RegisterCapital(xlanguage.Japanese, "カブール")
}
