//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.Arabic, "نيوزيلندا")
	dataNewZealand.RegisterOfficialName(xlanguage.Arabic, "نيوزيلندا")
	dataNewZealand.RegisterCapital(xlanguage.Arabic, "ولينغتون")
}
