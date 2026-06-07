//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMongolia.RegisterName(xlanguage.Arabic, "منغوليا")
	dataMongolia.RegisterOfficialName(xlanguage.Arabic, "منغوليا")
	dataMongolia.RegisterCapital(xlanguage.Arabic, "أولان باتور")
}
