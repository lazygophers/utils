package fake

func init() {
	registerDataSet("es", "addresses", "cities", dsEs_0)
	registerDataSet("es", "addresses", "streets", dsEs_1)
	registerDataSet("es", "companies", "names", dsEs_2)
	registerDataSet("es", "companies", "suffixes", dsEs_3)
	registerDataSet("es", "names", "first_female", dsEs_4)
	registerDataSet("es", "names", "first_male", dsEs_5)
	registerDataSet("es", "names", "last", dsEs_6)
	registerDataSet("es", "texts", "lorem", dsEs_7)
}

var dsEs_0 = &DataSet{
	Language: "es",
	Country:  "ES",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Madrid", Weight: 2.0, Tags: []string{"capital", "major"}},
		{Value: "Barcelona", Weight: 1.8, Tags: []string{"major", "catalonia"}},
		{Value: "Valencia", Weight: 1.6, Tags: []string{"major", "coastal"}},
		{Value: "Sevilla", Weight: 1.4, Tags: []string{"historical", "andalusia"}},
		{Value: "Zaragoza", Weight: 1.2, Tags: []string{"major", "aragon"}},
		{Value: "Málaga", Weight: 1.1, Tags: []string{"coastal", "tourism"}},
		{Value: "Murcia", Tags: []string{"southeastern"}},
		{Value: "Palma", Weight: 0.9, Tags: []string{"island", "tourism"}},
		{Value: "Las Palmas", Weight: 0.8, Tags: []string{"canary", "island"}},
		{Value: "Bilbao", Weight: 0.7, Tags: []string{"basque", "industrial"}},
	},
}

var dsEs_1 = &DataSet{
	Language: "es",
	Country:  "ES",
	Type:     "addresses",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Calle Mayor", Weight: 2.0, Tags: []string{"common", "main"}},
		{Value: "Gran Vía", Weight: 1.8, Tags: []string{"major", "commercial"}},
		{Value: "Calle Real", Weight: 1.6, Tags: []string{"traditional", "royal"}},
		{Value: "Plaza Mayor", Weight: 1.4, Tags: []string{"central", "square"}},
		{Value: "Avenida de la Constitución", Weight: 1.2, Tags: []string{"major", "constitutional"}},
		{Value: "Calle de Alcalá", Weight: 1.1, Tags: []string{"famous", "madrid"}},
		{Value: "Paseo de Gracia", Tags: []string{"elegant", "barcelona"}},
		{Value: "Calle Serrano", Weight: 0.9, Tags: []string{"upscale", "shopping"}},
		{Value: "Rambla", Weight: 0.8, Tags: []string{"pedestrian", "tourist"}},
		{Value: "Plaza de España", Weight: 0.7, Tags: []string{"central", "square"}},
	},
}

var dsEs_2 = &DataSet{
	Language: "es",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "Telefónica", Weight: 2.0, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Santander", Weight: 1.85, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Repsol", Weight: 1.7, Tags: []string{"major"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Iberdrola", Weight: 1.55, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "BBVA", Weight: 1.4, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Inditex", Weight: 1.25, Tags: []string{"medium"}, Meta: map[string]string{"industry": "various"}},
		{Value: "ACS", Weight: 1.1, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Ferrovial", Weight: 0.95, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Endesa", Weight: 0.8, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
		{Value: "Mapfre", Weight: 0.6500000000000001, Tags: []string{"small"}, Meta: map[string]string{"industry": "various"}},
	},
}

var dsEs_3 = &DataSet{
	Language: "es",
	Type:     "companies",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "S.A.", Weight: 3.0, Tags: []string{"common"}},
		{Value: "S.L.", Weight: 2.7, Tags: []string{"common"}},
		{Value: "S.C.", Weight: 2.4, Tags: []string{"common"}},
		{Value: "S.A.U.", Weight: 2.1, Tags: []string{"formal"}},
		{Value: "Sociedad Limitada", Weight: 1.8, Tags: []string{"formal"}},
		{Value: "Sociedad Anónima", Weight: 1.5, Tags: []string{"formal"}},
	},
}

var dsEs_4 = &DataSet{
	Language: "es",
	Country:  "ES",
	Type:     "names",
	Version:  "2.0.0",
	Items: []DataItem{
		{Value: "María", Weight: 2.0, Tags: []string{"biblical", "traditional"}},
		{Value: "Carmen", Weight: 1.8, Tags: []string{"religious", "spanish"}},
		{Value: "Josefa", Weight: 1.6, Tags: []string{"biblical", "traditional"}},
		{Value: "Isabel", Weight: 1.4, Tags: []string{"royal", "biblical"}},
		{Value: "Ana", Weight: 1.2, Tags: []string{"biblical", "simple"}},
		{Value: "Francisca", Tags: []string{"religious", "saint"}},
		{Value: "Dolores", Tags: []string{"religious", "sorrows"}},
		{Value: "Antonia", Weight: 0.9, Tags: []string{"roman", "traditional"}},
		{Value: "María del Carmen", Weight: 0.9, Tags: []string{"compound", "religious"}},
		{Value: "Concepción", Weight: 0.8, Tags: []string{"religious", "immaculate"}},
		{Value: "Pilar", Weight: 0.8, Tags: []string{"religious", "spanish"}},
		{Value: "Teresa", Weight: 0.7, Tags: []string{"saint", "mystical"}},
		{Value: "María José", Weight: 0.7, Tags: []string{"compound", "popular"}},
		{Value: "Rosario", Weight: 0.6, Tags: []string{"religious", "rosary"}},
		{Value: "Encarnación", Weight: 0.6, Tags: []string{"religious", "incarnation"}},
		{Value: "Mercedes", Weight: 0.5, Tags: []string{"religious", "mercy"}},
		{Value: "Esperanza", Weight: 0.5, Tags: []string{"virtue", "hope"}},
		{Value: "Manuela", Weight: 0.5, Tags: []string{"biblical", "traditional"}},
		{Value: "Rosa", Weight: 0.4, Tags: []string{"flower", "simple"}},
		{Value: "Cristina", Weight: 0.4, Tags: []string{"christian", "modern"}},
		{Value: "Elena", Weight: 0.4, Tags: []string{"greek", "light"}},
		{Value: "Amparo", Weight: 0.4, Tags: []string{"religious", "protection"}},
		{Value: "Remedios", Weight: 0.3, Tags: []string{"religious", "remedy"}},
		{Value: "Montserrat", Weight: 0.3, Tags: []string{"catalan", "mountain"}},
		{Value: "Inmaculada", Weight: 0.3, Tags: []string{"religious", "pure"}},
		{Value: "Soledad", Weight: 0.3, Tags: []string{"religious", "solitude"}},
		{Value: "Rocío", Weight: 0.3, Tags: []string{"nature", "dew"}},
		{Value: "Patricia", Weight: 0.3, Tags: []string{"noble", "modern"}},
	},
}

var dsEs_5 = &DataSet{
	Language: "es",
	Country:  "ES",
	Type:     "names",
	Version:  "2.0.0",
	Items: []DataItem{
		{Value: "Antonio", Weight: 2.0, Tags: []string{"traditional", "roman"}},
		{Value: "Manuel", Weight: 1.8, Tags: []string{"biblical", "popular"}},
		{Value: "Francisco", Weight: 1.6, Tags: []string{"religious", "saint"}},
		{Value: "José", Weight: 1.4, Tags: []string{"biblical", "popular"}},
		{Value: "Juan", Weight: 1.2, Tags: []string{"biblical", "traditional"}},
		{Value: "David", Tags: []string{"biblical", "modern"}},
		{Value: "José Antonio", Tags: []string{"compound", "traditional"}},
		{Value: "José Luis", Weight: 0.9, Tags: []string{"compound", "popular"}},
		{Value: "Jesús", Weight: 0.9, Tags: []string{"religious", "biblical"}},
		{Value: "Javier", Weight: 0.8, Tags: []string{"basque", "saint"}},
		{Value: "Carlos", Weight: 0.8, Tags: []string{"royal", "germanic"}},
		{Value: "Miguel", Weight: 0.7, Tags: []string{"biblical", "archangel"}},
		{Value: "Alejandro", Weight: 0.7, Tags: []string{"greek", "defender"}},
		{Value: "Rafael", Weight: 0.6, Tags: []string{"biblical", "archangel"}},
		{Value: "Ángel", Weight: 0.6, Tags: []string{"religious", "angelic"}},
		{Value: "José María", Weight: 0.5, Tags: []string{"compound", "religious"}},
		{Value: "Fernando", Weight: 0.5, Tags: []string{"royal", "germanic"}},
		{Value: "Daniel", Weight: 0.5, Tags: []string{"biblical", "modern"}},
		{Value: "José Manuel", Weight: 0.4, Tags: []string{"compound", "popular"}},
		{Value: "Luis", Weight: 0.4, Tags: []string{"royal", "germanic"}},
		{Value: "Sergio", Weight: 0.4, Tags: []string{"roman", "modern"}},
		{Value: "Pablo", Weight: 0.4, Tags: []string{"apostle", "biblical"}},
		{Value: "Jorge", Weight: 0.3, Tags: []string{"saint", "greek"}},
		{Value: "Alberto", Weight: 0.3, Tags: []string{"germanic", "noble"}},
		{Value: "Pedro", Weight: 0.3, Tags: []string{"apostle", "biblical"}},
		{Value: "Adrián", Weight: 0.3, Tags: []string{"roman", "modern"}},
		{Value: "Raúl", Weight: 0.3, Tags: []string{"germanic", "modern"}},
		{Value: "Álvaro", Weight: 0.3, Tags: []string{"germanic", "noble"}},
	},
}

var dsEs_6 = &DataSet{
	Language: "es",
	Country:  "ES",
	Type:     "names",
	Version:  "2.0.0",
	Items: []DataItem{
		{Value: "García", Weight: 3.0, Tags: []string{"common", "patronymic"}},
		{Value: "Rodríguez", Weight: 2.8, Tags: []string{"common", "patronymic"}},
		{Value: "González", Weight: 2.6, Tags: []string{"common", "patronymic"}},
		{Value: "Fernández", Weight: 2.4, Tags: []string{"common", "patronymic"}},
		{Value: "López", Weight: 2.2, Tags: []string{"common", "patronymic"}},
		{Value: "Martínez", Weight: 2.0, Tags: []string{"common", "patronymic"}},
		{Value: "Sánchez", Weight: 1.8, Tags: []string{"common", "patronymic"}},
		{Value: "Pérez", Weight: 1.6, Tags: []string{"common", "patronymic"}},
		{Value: "Gómez", Weight: 1.4, Tags: []string{"common", "patronymic"}},
		{Value: "Martín", Weight: 1.2, Tags: []string{"common", "roman"}},
		{Value: "Jiménez", Tags: []string{"common", "patronymic"}},
		{Value: "Ruiz", Tags: []string{"common", "patronymic"}},
		{Value: "Hernández", Weight: 0.9, Tags: []string{"common", "patronymic"}},
		{Value: "Moreno", Weight: 0.9, Tags: []string{"descriptive", "color"}},
		{Value: "Muñoz", Weight: 0.8, Tags: []string{"common", "patronymic"}},
		{Value: "Álvarez", Weight: 0.8, Tags: []string{"common", "patronymic"}},
		{Value: "Castillo", Weight: 0.7, Tags: []string{"topographic", "castle"}},
		{Value: "Romero", Weight: 0.7, Tags: []string{"religious", "pilgrim"}},
		{Value: "Gutiérrez", Weight: 0.6, Tags: []string{"common", "patronymic"}},
		{Value: "Ortega", Weight: 0.6, Tags: []string{"topographic", "nettle"}},
		{Value: "Delgado", Weight: 0.5, Tags: []string{"descriptive", "thin"}},
		{Value: "Castro", Weight: 0.5, Tags: []string{"topographic", "fort"}},
		{Value: "Ortiz", Weight: 0.5, Tags: []string{"common", "patronymic"}},
		{Value: "Rubio", Weight: 0.4, Tags: []string{"descriptive", "blonde"}},
		{Value: "Ramos", Weight: 0.4, Tags: []string{"topographic", "branches"}},
		{Value: "Vázquez", Weight: 0.4, Tags: []string{"common", "patronymic"}},
		{Value: "Herrera", Weight: 0.4, Tags: []string{"occupational", "blacksmith"}},
		{Value: "Flores", Weight: 0.3, Tags: []string{"topographic", "flowers"}},
		{Value: "Navarro", Weight: 0.3, Tags: []string{"geographic", "navarre"}},
		{Value: "Rivera", Weight: 0.3, Tags: []string{"topographic", "river"}},
	},
}

var dsEs_7 = &DataSet{
	Language: "es",
	Type:     "texts",
	Version:  "1.0.0",
	Items: []DataItem{
		{Value: "texto"},
		{Value: "palabra", Weight: 0.95},
		{Value: "oración", Weight: 0.9},
		{Value: "párrafo", Weight: 0.85},
		{Value: "artículo", Weight: 0.8},
		{Value: "informe", Weight: 0.75},
		{Value: "estudio", Weight: 0.7},
		{Value: "análisis", Weight: 0.6499999999999999},
		{Value: "contenido", Weight: 0.6},
		{Value: "tema", Weight: 0.55},
		{Value: "concepto", Weight: 0.5},
		{Value: "significado", Weight: 0.44999999999999996},
		{Value: "idioma", Weight: 0.3999999999999999},
		{Value: "literatura", Weight: 0.35},
		{Value: "escritura", Weight: 0.29999999999999993},
	},
}
