package models

import "gopkg.in/yaml.v2"

type studyDirection int

const (
	DirMachineLearning studyDirection = iota
	DirDistributedSystems
	DirIndustrialDevelopment
	DirTheoreticalInformatics
	DirDataAnalysis
)

type Student struct {
	Name       string         `yaml:"name"`
	Surname    string         `yaml:"surname"`
	Age        int            `yaml:"age"`
	Percentile float32        `yaml:"percentile"`
	Direction  studyDirection `yaml:"direction"`
	Courses    []string       `yaml:"courses"`
	Marks      map[string]int `yaml:"marks"`
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

func (s *Student) SerializeYAML() ([]byte, error) {
	return yaml.Marshal(s)
}

func DeserializeYAML(data []byte, s *Student) error {
	return yaml.Unmarshal(data, s)
}
