//go:build lang_ja || lang_all

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataBonaireSintEustatiusAndSaba.RegisterName(xlanguage.Japanese, "ボネール、シント・ユースタティウスおよびサバ")
	dataBonaireSintEustatiusAndSaba.RegisterOfficialName(xlanguage.Japanese, "ボネール、シント・ユースタティウスおよびサバ")
	dataBonaireSintEustatiusAndSaba.RegisterCapital(xlanguage.Japanese, "クラレンダイク")
}
