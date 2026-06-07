//go:build (lang_ar || lang_all) && (country_africa || country_all || country_cd || country_middle_africa)

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataDrCongo.RegisterName(xlanguage.Arabic, "جمهورية الكونغو الديمقراطية")
	dataDrCongo.RegisterOfficialName(xlanguage.Arabic, "جمهورية الكونغو الديمقراطية")
	dataDrCongo.RegisterCapital(xlanguage.Arabic, "كينشاسا")
}
