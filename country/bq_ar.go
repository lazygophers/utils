//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.Arabic, "بونير وسينت أوستاتيوس وسابا")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.Arabic, "هولندا الكاريبية")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.Arabic, "كرالنديك")
}
