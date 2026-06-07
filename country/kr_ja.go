//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataSouthKorea.RegisterName(xlanguage.Japanese, "大韓民国")
	dataSouthKorea.RegisterOfficialName(xlanguage.Japanese, "大韓民国")
	dataSouthKorea.RegisterCapital(xlanguage.Japanese, "ソウル")
}
