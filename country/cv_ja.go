//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCaboVerde.RegisterName(xlanguage.Japanese, "カーボベルデ")
	dataCaboVerde.RegisterOfficialName(xlanguage.Japanese, "カーボベルデ共和国")
	dataCaboVerde.RegisterCapital(xlanguage.Japanese, "プライア")
}
