//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAustria.RegisterName(xlanguage.Japanese, "オーストリア")
	dataAustria.RegisterOfficialName(xlanguage.Japanese, "オーストリア共和国")
	dataAustria.RegisterCapital(xlanguage.Japanese, "ウィーン")
}
