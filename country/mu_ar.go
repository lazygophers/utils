//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMauritius.RegisterName(xlanguage.Arabic, "موريشيوس")
	dataMauritius.RegisterOfficialName(xlanguage.Arabic, "جمهورية موريشيوس")
	dataMauritius.RegisterCapital(xlanguage.Arabic, "بورت لويس")
}
