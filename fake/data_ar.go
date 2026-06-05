package fake

func init() {
	registerDataSet("ar", "addresses", "cities", dsAr_0)
	registerDataSet("ar", "addresses", "streets", dsAr_1)
	registerDataSet("ar", "companies", "names", dsAr_2)
	registerDataSet("ar", "companies", "suffixes", dsAr_3)
	registerDataSet("ar", "names", "first_female", dsAr_4)
	registerDataSet("ar", "names", "first_male", dsAr_5)
	registerDataSet("ar", "names", "last", dsAr_6)
	registerDataSet("ar", "texts", "lorem", dsAr_7)
}

var dsAr_0 = &DataSet{
	Language: "ar",
	Country:  "SA",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "الرياض", Weight: 2.0, Tags: []string{"capital", "major"}, Meta: map[string]string{"region": "نجد"}},
		{Value: "جدة", Weight: 1.8, Tags: []string{"major", "coast"}, Meta: map[string]string{"region": "الحجاز"}},
		{Value: "مكة المكرمة", Weight: 1.6, Tags: []string{"holy", "major"}, Meta: map[string]string{"region": "الحجاز"}},
		{Value: "المدينة المنورة", Weight: 1.4, Tags: []string{"holy", "major"}, Meta: map[string]string{"region": "الحجاز"}},
		{Value: "الدمام", Weight: 1.2, Tags: []string{"major", "east"}, Meta: map[string]string{"region": "الشرقية"}},
		{Value: "تبوك", Tags: []string{"north"}, Meta: map[string]string{"region": "تبوك"}},
		{Value: "أبها", Weight: 0.9, Tags: []string{"south", "mountain"}, Meta: map[string]string{"region": "عسير"}},
		{Value: "الطائف", Weight: 0.8, Tags: []string{"summer", "mountain"}, Meta: map[string]string{"region": "الحجاز"}},
		{Value: "بريدة", Weight: 0.7, Tags: []string{"central"}, Meta: map[string]string{"region": "القصيم"}},
		{Value: "خميس مشيط", Weight: 0.6, Tags: []string{"south"}, Meta: map[string]string{"region": "عسير"}},
	},
}

var dsAr_1 = &DataSet{
	Language: "ar",
	Country:  "SA",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "شارع الملك فهد", Weight: 2.0, Tags: []string{"major", "royal"}},
		{Value: "شارع العليا", Weight: 1.8, Tags: []string{"commercial", "major"}},
		{Value: "طريق الملك عبدالعزيز", Weight: 1.6, Tags: []string{"major", "royal"}},
		{Value: "شارع التحلية", Weight: 1.4, Tags: []string{"commercial", "upscale"}},
		{Value: "شارع الأمير محمد بن عبدالعزيز", Weight: 1.2, Tags: []string{"major", "royal"}},
		{Value: "شارع الجامعة", Tags: []string{"education"}},
		{Value: "شارع المطار", Weight: 0.9, Tags: []string{"transport"}},
		{Value: "شارع الخالدية", Weight: 0.8, Tags: []string{"residential"}},
		{Value: "طريق الدائري", Weight: 0.7, Tags: []string{"highway"}},
		{Value: "شارع الحرمين", Weight: 0.6, Tags: []string{"religious"}},
	},
}

var dsAr_2 = &DataSet{
	Language: "ar",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "سابك", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "أرامكو", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "الراجحي", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "سامبا", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "الأهلي", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "موبايلي", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "زين", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "الكهرباء", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "مصرف الإنماء", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "صدق", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsAr_3 = &DataSet{
	Language: "ar",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "شركة", Weight: 3.0, Tags: []string{"common"}},
		{Value: "مؤسسة", Weight: 2.7, Tags: []string{"common"}},
		{Value: "مجموعة", Weight: 2.4, Tags: []string{"common"}},
		{Value: "القابضة", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "للتجارة", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "للخدمات", Weight: 1.5, Tags: []string{"formal"}},
		{Value: "المحدودة", Weight: 1.2000000000000002, Tags: []string{"formal"}},
	},
}

var dsAr_4 = &DataSet{
	Language: "ar",
	Country:  "SA",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "فاطمة", Weight: 2.0, Tags: []string{"islamic", "traditional"}},
		{Value: "عائشة", Weight: 1.8, Tags: []string{"islamic", "traditional"}},
		{Value: "خديجة", Weight: 1.6, Tags: []string{"islamic", "traditional"}},
		{Value: "زينب", Weight: 1.4, Tags: []string{"islamic", "beautiful"}},
		{Value: "مريم", Weight: 1.4, Tags: []string{"biblical", "traditional"}},
		{Value: "أسماء", Weight: 1.2, Tags: []string{"traditional", "noble"}},
		{Value: "نورا", Weight: 1.2, Tags: []string{"modern", "light"}},
		{Value: "سارة", Weight: 1.1, Tags: []string{"biblical", "traditional"}},
		{Value: "رقية", Tags: []string{"islamic", "traditional"}},
		{Value: "هدى", Weight: 0.9, Tags: []string{"islamic", "guidance"}},
	},
}

var dsAr_5 = &DataSet{
	Language: "ar",
	Country:  "SA",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "محمد", Weight: 2.0, Tags: []string{"islamic", "traditional"}},
		{Value: "أحمد", Weight: 1.8, Tags: []string{"islamic", "traditional"}},
		{Value: "علي", Weight: 1.6, Tags: []string{"islamic", "noble"}},
		{Value: "حسن", Weight: 1.4, Tags: []string{"islamic", "virtue"}},
		{Value: "عمر", Weight: 1.4, Tags: []string{"islamic", "traditional"}},
		{Value: "خالد", Weight: 1.2, Tags: []string{"traditional", "strong"}},
		{Value: "يوسف", Weight: 1.2, Tags: []string{"biblical", "traditional"}},
		{Value: "إبراهيم", Weight: 1.1, Tags: []string{"biblical", "traditional"}},
		{Value: "عبدالله", Tags: []string{"islamic", "religious"}},
		{Value: "عبدالرحمن", Weight: 0.9, Tags: []string{"islamic", "religious"}},
	},
}

var dsAr_6 = &DataSet{
	Language: "ar",
	Country:  "SA",
	Type:     "names",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "الأحمد", Weight: 2.0, Tags: []string{"patronymic"}},
		{Value: "العلي", Weight: 1.8, Tags: []string{"patronymic"}},
		{Value: "المحمد", Weight: 1.6, Tags: []string{"patronymic"}},
		{Value: "الحسن", Weight: 1.4, Tags: []string{"patronymic"}},
		{Value: "العمر", Weight: 1.2, Tags: []string{"patronymic"}},
		{Value: "الخالد", Weight: 1.1, Tags: []string{"patronymic"}},
		{Value: "اليوسف", Tags: []string{"patronymic"}},
		{Value: "الإبراهيم", Weight: 0.9, Tags: []string{"patronymic"}},
		{Value: "العبدالله", Weight: 0.8, Tags: []string{"patronymic"}},
		{Value: "الصالح", Weight: 0.7, Tags: []string{"descriptive"}},
	},
}

var dsAr_7 = &DataSet{
	Language: "ar",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "النص"},
		{Value: "الكلام", Weight: 0.95},
		{Value: "الجملة", Weight: 0.9},
		{Value: "الفقرة", Weight: 0.85},
		{Value: "المقطع", Weight: 0.8},
		{Value: "الموضوع", Weight: 0.75},
		{Value: "البحث", Weight: 0.7},
		{Value: "الدراسة", Weight: 0.6499999999999999},
		{Value: "التقرير", Weight: 0.6},
		{Value: "الكتابة", Weight: 0.55},
		{Value: "القراءة", Weight: 0.5},
		{Value: "التعلم", Weight: 0.44999999999999996},
		{Value: "المعرفة", Weight: 0.3999999999999999},
		{Value: "العلم", Weight: 0.35},
		{Value: "الثقافة", Weight: 0.29999999999999993},
	},
}
