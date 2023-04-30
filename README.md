# Serialization

Приложение для тестирования различных форматов сериализации. При этом тестирование каждого формата осуществляется в отдельном контейнере.

Образы берем с DockerHub: https://hub.docker.com/repositories/morozhka.  
Используемые порты и пути для получения ответа следующие:

    ___________________ port _________ path _________
    json-service:       3030,       /get_result/json
    yaml-service:       3031,       /get_result/yaml
    gob-service:        3032,       /get_result/gob
    proto-service:      3033,       /get_result/proto
    avro-service:       3034,       /get_result/avro

### Usage
1. Для запуска контейнеров можно использовать следующую команду:
```docker-compose -f ./docker-compose.yml up -d --build```

2. По запросу `get_result/{format}` сервис возвращает ответ вида:  `{Формат сериализации} – {Размер сериализованной структуры/объекта в байтах} – {Время сериализации}`.    
_Например:_  
При переходе на ```http://localhost:3030/get_result/json``` получили ответ `"JSON - 24 - 186.886µs - 360.14µs"`



### Data
Сериализируемые данные представляют собой следующую структуру:

```Go
type Student struct {
	Name       string         `json:"name" yaml:"name" avro:"name" xml:"name"`
	Surname    string         `json:"surname" yaml:"surname" avro:"surname" xml:"surname"`
	Age        int            `json:"age" yaml:"age" avro:"age" xml:"age"`
	Percentile float32        `json:"percentile" yaml:"percentile" avro:"percentile" xml:"percentile"`
	Direction  studyDirection `json:"direction" yaml:"direction" avro:"direction" xml:"direction"`
	Courses    []string       `json:"courses" yaml:"courses" avro:"courses" xml:"courses"`
	Marks      map[string]int `json:"marks" yaml:"marks" avro:"marks" xml:"marks"`
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

### Results
Усредненные результаты получились следующими:

    GOB   - 24 -  387.892µs  - 1377.586µs
    JSON  - 24 -  175.305µs  -  360.642µs
    YAML  - 24 - 1279.606µs  - 1427.861µs
    PROTO - 24 -  184.803µs  -  136.391µs
    AVRO  - 24 -   95.176µs  -   71.244µs

Форматы по возрастанию времени преобразования:  

_Serialization:_    ```AVRO (95µs)``` - ```JSON (175µs)```  - ```PROTO (184µs)``` - ```GOB (387µs)``` - ```YAML (1279µs)```

_Deserialization:_  ```AVRO (71µs)``` - ```PROTO (136µs)``` - ```JSON (360µs)```  - ```GOB(1377µs)``` - ```YAML(1427µs)```