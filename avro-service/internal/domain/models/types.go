package models

import (
	"github.com/hamba/avro"
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
	Name       string         `avro:"name"`
	Surname    string         `avro:"surname"`
	Age        int            `avro:"age"`
	Percentile float32        `avro:"percentile"`
	Direction  studyDirection `avro:"direction"`
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

func (s *Student) SerializeAvro(schema avro.Schema) ([]byte, error) {
	return avro.Marshal(schema, s)
}

func DeserializeAvro(schema avro.Schema, data []byte, s *Student) error {
	return avro.Unmarshal(schema, data, s)
}
