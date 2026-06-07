//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPakistan.RegisterName(xlanguage.Japanese, "パキスタン")
	dataPakistan.RegisterOfficialName(xlanguage.Japanese, "パキスタン・イスラム共和国")
	dataPakistan.RegisterCapital(xlanguage.Japanese, "イスラマバード")
}
