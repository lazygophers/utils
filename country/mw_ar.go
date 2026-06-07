//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMalawi.RegisterName(xlanguage.Arabic, "مالاوي")
	dataMalawi.RegisterOfficialName(xlanguage.Arabic, "جمهورية مالاوي")
	dataMalawi.RegisterCapital(xlanguage.Arabic, "ليلونغوي")
}
