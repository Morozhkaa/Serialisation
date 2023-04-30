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
    msgpack-service:    3035,       /get_result/msgpack
    xml-service:        3036,       /get_result/xml

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
	Name       string         `json:"name" yaml:"name" avro:"name" msgpack:"name"`
	Surname    string         `json:"surname" yaml:"surname" avro:"surname" msgpack:"surname"`
	Age        int            `json:"age" yaml:"age" avro:"age" msgpack:"age"`
	Percentile float32        `json:"percentile" yaml:"percentile" avro:"percentile" msgpack:"percentile"`
	Direction  studyDirection `json:"direction" yaml:"direction" avro:"direction" msgpack:"direction"`
	Courses    []string       `json:"courses" yaml:"courses" avro:"courses" msgpack:"courses"`
	Marks      map[string]int `json:"marks" yaml:"marks" avro:"marks" msgpack:"marks"`
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


### Features when working with some formats

Сервисы имеют схожую архитектуру, основные изменения в коде происходили в директории `internal/domain`. Краткое описание, что было сделано:

    JSON: к структуре добавила теги вида `json:"name"`, использовала json.Marshal, json.Unmarshal функции.  

    YAML: к структуре добавила теги вида `yaml:"name"`, использовала yaml.Marshal, yaml.Unmarshal функции.

    GOB: создала bytes.Buffer, инициализировала gob.NewEncoder и gob.NewDecoder, использовала фунции Encode(), Decode().

    PROTO: создала student.proto файл, по которому сгенерировала структуру Student. Для этого из директории internal/domain/models/proto выполнила комманды:  
           > export PATH=$PATH:$GOPATH/bin
           > protoc --go_out=./gen student.proto
           Перенесла сгенерированный код в файл types.go, добавив функцию инициализации. Использовала proto.Marshal, proto.Unmarshal.

    AVRO: к структуре добавила теги вида `avro:"name"`, объявила схему, использовала avro.Marshal, avro.Unmarshal функции.  

    MSGPACK: к структуре добавила теги вида `msgpack:"name"`, использовала msgpack.Marshal, msgpack.Unmarshal функции.


```Go
XML: к структуре добавила теги вида `xml:"name"`, поскольку xml не поддерживает сериализацию map, немного изменила структуру и использовала xml.Marshal, xml.Unmarshal функции.

type Marks struct {
	Subject string `xml:"subject"`
	Mark    int    `xml:"mark"`
}

type Student struct {
	Name       string         `xml:"name"`
	Surname    string         `xml:"surname"`
	Age        int            `xml:"age"`
	Percentile float32        `xml:"percentile"`
	Direction  studyDirection `xml:"direction"`
	Courses    []string       `xml:"courses"`
	Marks      []Marks        `xml:"marks"`
}
```


### Results
Усредненные результаты получились следующими:

    GOB      - 24 -  387.892µs  - 1377.586µs
    JSON     - 24 -  175.305µs  -  360.642µs
    YAML     - 24 - 1279.606µs  - 1427.861µs
    PROTO    - 24 -  184.803µs  -  136.391µs
    AVRO     - 24 -   95.176µs  -   71.244µs
    MSGPACK  - 24 -  242.366µs  -  216.944µs
    XML      - 24 -  540.252µs  - 1967.406µs

Форматы по возрастанию времени преобразования:  

_Serialization:_    ```AVRO (95µs)``` - ```JSON (175µs)```  - ```PROTO (184µs)``` -  ```MSGPACK (242µs)``` - ```GOB (387µs)``` - ```XML(540µs)``` - ```YAML (1279µs)```

_Deserialization:_  ```AVRO (71µs)``` - ```PROTO (136µs)``` - ```MSGPACK (216µs)``` - ```JSON (360µs)```  - ```GOB(1377µs)``` - ```YAML(1427µs)``` - ```XML(1967µs)```