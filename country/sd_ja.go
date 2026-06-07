//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSudan.RegisterName(xlanguage.Japanese, "スーダン")
	dataSudan.RegisterOfficialName(xlanguage.Japanese, "スーダン共和国")
	dataSudan.RegisterCapital(xlanguage.Japanese, "ハルツーム")
}
