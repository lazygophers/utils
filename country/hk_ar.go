//go:build lang_ar || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataHongKong.RegisterName(xlanguage.Arabic, "هونغ كونغ")
	dataHongKong.RegisterOfficialName(xlanguage.Arabic, "منطقة هونغ كونغ الإدارية الخاصة بجمهورية الصين الشعبية")
	dataHongKong.RegisterCapital(xlanguage.Arabic, "هونغ كونغ")
}
