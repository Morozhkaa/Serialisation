package models

import (
	"bytes"
	"encoding/gob"
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

func (s *Student) SerializeGOB() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(s)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DeserializeGOB(data []byte, s *Student) error {
	buf := bytes.NewBuffer(data)
	enc := gob.NewDecoder(buf)
	err := enc.Decode(s)
	if err != nil {
		return err
	}
	return nil
}
