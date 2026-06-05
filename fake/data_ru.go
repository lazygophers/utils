package fake

func init() {
	registerDataSet("ru", "addresses", "cities", dsRu_0)
	registerDataSet("ru", "addresses", "streets", dsRu_1)
	registerDataSet("ru", "companies", "names", dsRu_2)
	registerDataSet("ru", "companies", "suffixes", dsRu_3)
	registerDataSet("ru", "names", "first_female", dsRu_4)
	registerDataSet("ru", "names", "first_male", dsRu_5)
	registerDataSet("ru", "names", "last", dsRu_6)
	registerDataSet("ru", "texts", "lorem", dsRu_7)
}

var dsRu_0 = &DataSet{
	Language: "ru",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Москва", Weight: 2.0, Tags: []string{"major"}},
		{Value: "Санкт-Петербург", Weight: 1.8, Tags: []string{"major"}},
		{Value: "Новосибирск", Weight: 1.6, Tags: []string{"major"}},
		{Value: "Екатеринбург", Weight: 1.4, Tags: []string{"medium"}},
		{Value: "Казань", Weight: 1.2, Tags: []string{"medium"}},
		{Value: "Нижний Новгород", Tags: []string{"medium"}},
		{Value: "Челябинск", Weight: 0.7999999999999998, Tags: []string{"medium"}},
		{Value: "Самара", Weight: 0.5999999999999999, Tags: []string{"medium"}},
		{Value: "Омск", Weight: 0.5, Tags: []string{"medium"}},
		{Value: "Ростов-на-Дону", Weight: 0.5, Tags: []string{"medium"}},
	},
}

var dsRu_1 = &DataSet{
	Language: "ru",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Тверская улица", Weight: 2.0, Tags: []string{"common"}},
		{Value: "Невский проспект", Weight: 1.7, Tags: []string{"common"}},
		{Value: "Красная площадь", Weight: 1.4, Tags: []string{"common"}},
		{Value: "Арбат", Weight: 1.1, Tags: []string{"common"}},
		{Value: "проспект Мира", Weight: 0.8, Tags: []string{"common"}},
	},
}

var dsRu_2 = &DataSet{
	Language: "ru",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Газпром", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Лукойл", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Роснефть", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Сбербанк", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Яндекс", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "МТС", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Вымпелком", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Норникель", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Магнит", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Тинькофф", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsRu_3 = &DataSet{
	Language: "ru",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "ОАО", Weight: 3.0, Tags: []string{"common"}},
		{Value: "ЗАО", Weight: 2.7, Tags: []string{"common"}},
		{Value: "ООО", Weight: 2.4, Tags: []string{"common"}},
		{Value: "ОДО", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "ИП", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "АО", Weight: 1.5, Tags: []string{"formal"}},
		{Value: "Группа", Weight: 1.2000000000000002, Tags: []string{"formal"}},
		{Value: "Холдинг", Weight: 0.8999999999999999, Tags: []string{"formal"}},
	},
}

var dsRu_4 = &DataSet{
	Language: "ru",
	Country:  "RU",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Анна", Weight: 2.0, Tags: []string{"biblical", "grace"}},
		{Value: "Мария", Weight: 1.8, Tags: []string{"biblical", "bitter"}},
		{Value: "Елена", Weight: 1.6, Tags: []string{"classical", "light"}},
		{Value: "Ольга", Weight: 1.4, Tags: []string{"traditional", "holy"}},
		{Value: "Татьяна", Weight: 1.2, Tags: []string{"traditional", "father"}},
		{Value: "Наталья", Weight: 1.1, Tags: []string{"traditional", "birth"}},
		{Value: "Ирина", Tags: []string{"traditional", "peace"}},
		{Value: "Светлана", Weight: 0.9, Tags: []string{"traditional", "light"}},
		{Value: "Екатерина", Weight: 0.8, Tags: []string{"traditional", "pure"}},
		{Value: "Юлия", Weight: 0.7, Tags: []string{"classical", "youthful"}},
	},
}

var dsRu_5 = &DataSet{
	Language: "ru",
	Country:  "RU",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Александр", Weight: 2.0, Tags: []string{"traditional", "defender"}},
		{Value: "Дмитрий", Weight: 1.8, Tags: []string{"traditional", "earth"}},
		{Value: "Максим", Weight: 1.6, Tags: []string{"traditional", "greatest"}},
		{Value: "Сергей", Weight: 1.4, Tags: []string{"traditional", "servant"}},
		{Value: "Андрей", Weight: 1.2, Tags: []string{"biblical", "brave"}},
		{Value: "Алексей", Weight: 1.1, Tags: []string{"traditional", "defender"}},
		{Value: "Артём", Tags: []string{"modern", "healthy"}},
		{Value: "Илья", Weight: 0.9, Tags: []string{"biblical", "traditional"}},
		{Value: "Кирилл", Weight: 0.8, Tags: []string{"traditional", "lord"}},
		{Value: "Михаил", Weight: 0.7, Tags: []string{"biblical", "traditional"}},
	},
}

var dsRu_6 = &DataSet{
	Language: "ru",
	Country:  "RU",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Смирнов", Weight: 2.0, Tags: []string{"descriptive"}},
		{Value: "Иванов", Weight: 1.8, Tags: []string{"patronymic"}},
		{Value: "Кузнецов", Weight: 1.6, Tags: []string{"occupational"}},
		{Value: "Попов", Weight: 1.4, Tags: []string{"occupational"}},
		{Value: "Соколов", Weight: 1.2, Tags: []string{"animal"}},
		{Value: "Лебедев", Weight: 1.1, Tags: []string{"animal"}},
		{Value: "Козlov", Tags: []string{"animal"}},
		{Value: "Новиков", Weight: 0.9, Tags: []string{"descriptive"}},
		{Value: "Морозов", Weight: 0.8, Tags: []string{"descriptive"}},
		{Value: "Петров", Weight: 0.7, Tags: []string{"patronymic"}},
	},
}

var dsRu_7 = &DataSet{
	Language: "ru",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "текст"},
		{Value: "слово", Weight: 0.95},
		{Value: "предложение", Weight: 0.9},
		{Value: "параграф", Weight: 0.85},
		{Value: "статья", Weight: 0.8},
		{Value: "отчёт", Weight: 0.75},
		{Value: "исследование", Weight: 0.7},
		{Value: "анализ", Weight: 0.6499999999999999},
		{Value: "содержание", Weight: 0.6},
		{Value: "тема", Weight: 0.55},
		{Value: "понятие", Weight: 0.5},
		{Value: "значение", Weight: 0.44999999999999996},
		{Value: "язык", Weight: 0.3999999999999999},
		{Value: "литература", Weight: 0.35},
		{Value: "письмо", Weight: 0.29999999999999993},
	},
}
