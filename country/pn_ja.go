//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataPitcairn.RegisterName(xlanguage.Japanese, "ピトケアン諸島")
	dataPitcairn.RegisterOfficialName(xlanguage.Japanese, "ピトケアン諸島")
	dataPitcairn.RegisterCapital(xlanguage.Japanese, "アダムスタウン")
}
