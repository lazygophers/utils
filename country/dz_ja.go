//go:build (lang_ja || lang_all) && (country_africa || country_all || country_dz || country_northern_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataAlgeria.RegisterName(xlanguage.Japanese, "アルジェリア")
	dataAlgeria.RegisterOfficialName(xlanguage.Japanese, "アルジェリア民主人民共和国")
	dataAlgeria.RegisterCapital(xlanguage.Japanese, "アルジェ")
}
