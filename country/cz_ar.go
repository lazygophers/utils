//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataCzechia.RegisterName(xlanguage.Arabic, "التشيك")
	dataCzechia.RegisterOfficialName(xlanguage.Arabic, "الجمهورية التشيكية")
	dataCzechia.RegisterCapital(xlanguage.Arabic, "براغ")
}
