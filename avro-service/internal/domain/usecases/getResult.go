package usecases

import (
	"avro-service/internal/domain/models"
	"context"
	"fmt"
	"time"
	"unsafe"

	"github.com/hamba/avro"
	"go.uber.org/zap"
)

type SerializationApp struct{}

func New() *SerializationApp {
	return &SerializationApp{}
}

func (s *SerializationApp) GetResult(ctx context.Context) (string, error) {
	log := zap.L()
	numberIterations := 10000
	var data []byte
	var err error

	schema, err := avro.Parse(`{
		"type": "record",
		"name": "student",
		"fields" : [
			{"name": "name", "type": "string"},
			{"name": "surname", "type": "string"},
			{"name": "age", "type": "int"},
			{"name": "percentile", "type": "float"},
			{"name": "direction", "type": "int"},
			{ 
				"name":"courses",  
				"type": { 
					"type": "array",
					"items": "string"
				} 
			},
			{ 
				"name":"marks",  
				"type": { 
					"type": "map",
					"values": "int"
				} 
			}
		]
	}`)
	if err != nil {
		return "", err
	}

	// run the serialization process numberIterations times, get the average execution time
	start := time.Now()
	for i := 0; i < numberIterations; i++ {
		student := models.NewStudent()
		data, err = student.SerializeAvro(schema)
		if err != nil {
			log.Error("serialization failed: " + err.Error())
			return "", err
		}
	}
	serTime := time.Since(start) / 1000

	// run the deserialization process numberIterations times, get the average execution time
	start = time.Now()
	for i := 0; i < numberIterations; i++ {
		var student models.Student
		err = models.DeserializeAvro(schema, data, &student)
		if err != nil {
			log.Error("deserialization failed: " + err.Error())
			return "", err
		}
	}
	deserTime := time.Since(start) / 1000

	// print the result in the format: <SerFormat> - <SerStructureSizeInBytes> - <SerTime> - <DeserTime>
	res := fmt.Sprintf("AVRO - %d - %d.%dµs - %d.%dµs", unsafe.Sizeof(data), serTime/1000, serTime%1000, deserTime/1000, deserTime%1000)
	return res, err
}
