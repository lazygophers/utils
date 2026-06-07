//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCuracao.RegisterName(xlanguage.Japanese, "キュラソー")
	dataCuracao.RegisterOfficialName(xlanguage.Japanese, "キュラソー")
	dataCuracao.RegisterCapital(xlanguage.Japanese, "ウィレムスタット")
}
