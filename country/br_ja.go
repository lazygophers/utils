//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBrazil.RegisterName(xlanguage.Japanese, "ブラジル")
	dataBrazil.RegisterOfficialName(xlanguage.Japanese, "ブラジル連邦共和国")
	dataBrazil.RegisterCapital(xlanguage.Japanese, "ブラジリア")
}
