//go:build (lang_ko || lang_all) && (country_ag || country_all || country_americas || country_caribbean)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAntiguaAndBarbuda.RegisterName(xlanguage.Korean, "앤티가 바부다")
	dataAntiguaAndBarbuda.RegisterOfficialName(xlanguage.Korean, "앤티가 바부다")
	dataAntiguaAndBarbuda.RegisterCapital(xlanguage.Korean, "세인트존스")
}
