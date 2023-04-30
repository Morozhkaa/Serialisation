package models

import (
	"encoding/xml"
)

type studyDirection int

const (
	DirMachineLearning studyDirection = iota
	DirDistributedSystems
	DirIndustrialDevelopment
	DirTheoreticalInformatics
	DirDataAnalysis
)

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

func NewStudent() *Student {
	return &Student{
		Name:       "Maria",
		Surname:    "Mironova",
		Age:        20,
		Percentile: 50.78,
		Direction:  DirIndustrialDevelopment,
		Courses: []string{"machine learning", "database", "design of fault-tolerant systems",
			"service-oriented architectures", "optimization methods"},
		// Marks: map[string]int{"machine learning": 5, "database": 7, "design of fault-tolerant systems": 8,
		// 	"service-oriented architectures": 8, "optimization methods": 6},
		Marks: []Marks{{Subject: "machine learning", Mark: 5}, {Subject: "database", Mark: 7},
			{Subject: "design of fault-tolerant systems", Mark: 8}, {Subject: "service-oriented architectures", Mark: 8},
			{Subject: "optimization methods", Mark: 6}},
	}
}

func (s *Student) SerializeXML() ([]byte, error) {
	return xml.Marshal(s)
}

func DeserializeXML(data []byte, s *Student) error {
	return xml.Unmarshal(data, s)
}
