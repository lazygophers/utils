//go:build country_all || country_mh || country_micronesia || country_oceania

package country

import xlanguage "golang.org/x/text/language"

func init() {
	dataMarshallIslands.RegisterName(xlanguage.Chinese, "马绍尔群岛")
	dataMarshallIslands.RegisterOfficialName(xlanguage.Chinese, "马绍尔群岛共和国")
	dataMarshallIslands.RegisterCapital(xlanguage.Chinese, "马朱罗")
}
