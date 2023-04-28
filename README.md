# Serialisation

Приложение для тестирования различных форматов сериализации. При этом тестирование каждого формата осуществляется в отдельном контейнере.

По запросу `get_result/{format}` сервис возвращает ответ вида:  `{Формат сериализации} – {Размер сериализованной структуры/объекта в байтах} – {Время сериализации}`

### Usage
Для запуска контейнеров можно использовать следующую команду:
```docker-compose -f ./docker-compose.yml up -d --build```

### Data
Сериализируемые данные представляют собой следующую структуру:

```Go
type Student struct {
	Name       string         `json:"name" yaml:"name" xml:"name"`
	Surname    string         `json:"surname" yaml:"surname" xml:"surname"`
	Age        int            `json:"age" yaml:"age" xml:"age"`
	Percentile float32        `json:"percentile" yaml:"percentile" xml:"percentile"`
	Direction  studyDirection `json:"direction" yaml:"direction" xml:"direction"`
	Courses    []string       `json:"courses" yaml:"courses" xml:"courses"`
	Marks      map[string]int `json:"marks" yaml:"marks" xml:"marks"`
}
```

Инициализируем следующим образом:
```Go
Student{
		Name:       "Maria",
		Surname:    "Mironova",
		Age:        20,
		Percentile: 50.78,
		Direction:  DirIndustrialDevelopment,
		Courses: []string{"machine learning", "database", "design of fault-tolerant systems",
			"service-oriented architectures", "optimization methods"},
		Marks: map[string]int{"machine learning": 5, "database": 7, "design of fault-tolerant systems": 8,
			"service-oriented architectures": 8, "optimization methods": 6},
	}
```