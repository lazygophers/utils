//go:build (lang_ja || lang_all) && (country_all || country_americas || country_bm || country_northern_america)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBermuda.RegisterName(xlanguage.Japanese, "バミューダ諸島")
	dataBermuda.RegisterOfficialName(xlanguage.Japanese, "バミューダ諸島")
	dataBermuda.RegisterCapital(xlanguage.Japanese, "ハミルトン")
}
