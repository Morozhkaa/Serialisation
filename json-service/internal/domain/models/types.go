package models

import (
	"encoding/json"
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
	Name       string         `json:"name" yaml:"name" xml:"name"`
	Surname    string         `json:"surname" yaml:"surname" xml:"surname"`
	Age        int            `json:"age" yaml:"age" xml:"age"`
	Percentile float32        `json:"percentile" yaml:"percentile" xml:"percentile"`
	Direction  studyDirection `json:"direction" yaml:"direction" xml:"direction"`
	Courses    []string       `json:"courses" yaml:"courses" xml:"courses"`
	Marks      map[string]int `json:"marks" yaml:"marks" xml:"marks"`
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

func (s *Student) SerializeJSON() ([]byte, error) {
	return json.Marshal(s)
}

func DeserializeJSON(data []byte, s *Student) error {
	return json.Unmarshal(data, s)
}
