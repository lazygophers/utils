package fake

import (
	"fmt"
	"strings"

	"github.com/lazygophers/utils/randx"
)

// Name 生成完整姓名
func (f *Faker) Name() string {
	f.incrementCallCount()
	
	firstName := f.FirstName()
	lastName := f.LastName()
	
	// 根据语言决定姓名顺序
	switch f.language {
	case LanguageChineseSimplified, LanguageChineseTraditional:
		return lastName + firstName
	default:
		return firstName + " " + lastName
	}
}

// FirstName 生成名字
func (f *Faker) FirstName() string {
	f.incrementCallCount()
	
	var dataType string
	switch f.gender {
	case GenderMale:
		dataType = "first_male"
	case GenderFemale:
		dataType = "first_female"
	default:
		// 随机选择性别
		if randx.Bool() {
			dataType = "first_male"
		} else {
			dataType = "first_female"
		}
	}
	
	values, weights, err := getWeightedItems(f.language, "names", dataType)
	if err != nil {
		// 如果当前语言不支持，回退到英语
		if f.language != LanguageEnglish {
			values, weights, err = getWeightedItems(LanguageEnglish, "names", dataType)
		}
		if err != nil {
			// 最终回退方案
			return getDefaultFirstName(f.gender)
		}
	}
	
	f.incrementGeneratedData()
	return randx.WeightedChoose(values, weights)
}

// LastName 生成姓氏
func (f *Faker) LastName() string {
	f.incrementCallCount()
	
	values, weights, err := getWeightedItems(f.language, "names", "last")
	if err != nil {
		// 如果当前语言不支持，回退到英语
		if f.language != LanguageEnglish {
			values, weights, err = getWeightedItems(LanguageEnglish, "names", "last")
		}
		if err != nil {
			// 最终回退方案
			return getDefaultLastName()
		}
	}
	
	f.incrementGeneratedData()
	return randx.WeightedChoose(values, weights)
}

// FullName 生成带中间名的完整姓名（仅英语支持）
func (f *Faker) FullName() string {
	f.incrementCallCount()
	
	firstName := f.FirstName()
	lastName := f.LastName()
	
	// 只有英语才有中间名概念
	if f.language == LanguageEnglish {
		// 30% 概率添加中间名
		if randx.Float32() < 0.3 {
			middleName := f.FirstName()
			return fmt.Sprintf("%s %s %s", firstName, middleName, lastName)
		}
	}
	
	return f.Name()
}

// NamePrefix 生成姓名前缀（如 Mr., Ms., Dr. 等）
func (f *Faker) NamePrefix() string {
	f.incrementCallCount()
	
	switch f.language {
	case LanguageEnglish:
		switch f.gender {
		case GenderMale:
			return randx.Choose([]string{"Mr.", "Dr.", "Prof."})
		case GenderFemale:
			return randx.Choose([]string{"Ms.", "Mrs.", "Miss", "Dr.", "Prof."})
		default:
			return randx.Choose([]string{"Mr.", "Ms.", "Mrs.", "Miss", "Dr.", "Prof."})
		}
	case LanguageChineseSimplified:
		switch f.gender {
		case GenderMale:
			return randx.Choose([]string{"先生", "老师", "教授", "博士"})
		case GenderFemale:
			return randx.Choose([]string{"女士", "小姐", "老师", "教授", "博士"})
		default:
			return randx.Choose([]string{"先生", "女士", "小姐", "老师", "教授", "博士"})
		}
	case LanguageChineseTraditional:
		switch f.gender {
		case GenderMale:
			return randx.Choose([]string{"先生", "老師", "教授", "博士"})
		case GenderFemale:
			return randx.Choose([]string{"女士", "小姐", "老師", "教授", "博士"})
		default:
			return randx.Choose([]string{"先生", "女士", "小姐", "老師", "教授", "博士"})
		}
	case LanguageFrench:
		switch f.gender {
		case GenderMale:
			return randx.Choose([]string{"M.", "Dr.", "Prof."})
		case GenderFemale:
			return randx.Choose([]string{"Mme.", "Mlle.", "Dr.", "Prof."})
		default:
			return randx.Choose([]string{"M.", "Mme.", "Mlle.", "Dr.", "Prof."})
		}
	default:
		return ""
	}
}

// NameSuffix 生成姓名后缀（如 Jr., Sr., III 等）
func (f *Faker) NameSuffix() string {
	f.incrementCallCount()
	
	if f.language == LanguageEnglish {
		suffixes := []string{"Jr.", "Sr.", "II", "III", "IV", "V"}
		// 10% 概率有后缀
		if randx.Float32() < 0.1 {
			return randx.Choose(suffixes)
		}
	}
	
	return ""
}

// FormattedName 生成格式化的姓名（包含前缀和后缀）
func (f *Faker) FormattedName() string {
	f.incrementCallCount()
	
	var parts []string
	
	prefix := f.NamePrefix()
	if prefix != "" {
		parts = append(parts, prefix)
	}
	
	name := f.Name()
	parts = append(parts, name)
	
	suffix := f.NameSuffix()
	if suffix != "" {
		parts = append(parts, suffix)
	}
	
	return strings.Join(parts, " ")
}

// BatchNames 批量生成姓名
func (f *Faker) BatchNames(count int) []string {
	f.incrementCallCount()
	
	return f.batchGenerate(count, f.Name)
}

// BatchFirstNames 批量生成名字
func (f *Faker) BatchFirstNames(count int) []string {
	f.incrementCallCount()
	
	return f.batchGenerate(count, f.FirstName)
}

// BatchLastNames 批量生成姓氏
func (f *Faker) BatchLastNames(count int) []string {
	f.incrementCallCount()
	
	return f.batchGenerate(count, f.LastName)
}

// 回退默认值
func getDefaultFirstName(gender Gender) string {
	switch gender {
	case GenderMale:
		return randx.Choose([]string{"John", "James", "Robert", "Michael", "William"})
	case GenderFemale:
		return randx.Choose([]string{"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth"})
	default:
		return randx.Choose([]string{"John", "Mary", "James", "Patricia", "Robert"})
	}
}

func getDefaultLastName() string {
	return randx.Choose([]string{"Smith", "Johnson", "Williams", "Brown", "Jones"})
}

// 全局便捷函数
func Name() string {
	return getDefaultFaker().Name()
}

func FirstName() string {
	return getDefaultFaker().FirstName()
}

func LastName() string {
	return getDefaultFaker().LastName()
}

func FullName() string {
	return getDefaultFaker().FullName()
}

func FormattedName() string {
	return getDefaultFaker().FormattedName()
}

func NamePrefix() string {
	return getDefaultFaker().NamePrefix()
}

func NameSuffix() string {
	return getDefaultFaker().NameSuffix()
}