//go:build (lang_ar || lang_all) && (country_all || country_australia_and_new_zealand || country_nz || country_oceania)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNewZealand.RegisterName(xlanguage.Arabic, "نيوزيلندا")
	dataNewZealand.RegisterOfficialName(xlanguage.Arabic, "نيوزيلندا")
	dataNewZealand.RegisterCapital(xlanguage.Arabic, "ولينغتون")
}
