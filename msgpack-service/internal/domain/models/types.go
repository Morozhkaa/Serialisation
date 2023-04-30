package models

import (
	"gopkg.in/vmihailenco/msgpack.v2"
)

type studyDirection int

const (
	DirMachineLearning studyDirection = iota
	DirDistributedSystems
	DirIndustrialDevelopment
	DirTheoreticalInformatics
	DirDataAnalysis
)

type Student struct {
	Name       string         `msgpack:"name"`
	Surname    string         `msgpack:"surname"`
	Age        int            `msgpack:"age"`
	Percentile float32        `msgpack:"percentile"`
	Direction  studyDirection `msgpack:"direction"`
	Courses    []string       `avro:"courses"`
	Marks      map[string]int `avro:"marks"`
}

func NewStudent() *Student {
	return &Student{
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
}

func (s *Student) SerializeMsgpack() ([]byte, error) {
	return msgpack.Marshal(s)
}

func DeserializeMsgpack(data []byte, s *Student) error {
	return msgpack.Unmarshal(data, s)
}
