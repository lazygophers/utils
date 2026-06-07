//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPortugal.RegisterName(xlanguage.Japanese, "ポルトガル")
	dataPortugal.RegisterOfficialName(xlanguage.Japanese, "ポルトガル共和国")
	dataPortugal.RegisterCapital(xlanguage.Japanese, "リスボン")
}
