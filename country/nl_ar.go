//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataNetherlands.RegisterName(xlanguage.Arabic, "هولندا")
	dataNetherlands.RegisterOfficialName(xlanguage.Arabic, "مملكة هولندا")
	dataNetherlands.RegisterCapital(xlanguage.Arabic, "أمستردام")
}
