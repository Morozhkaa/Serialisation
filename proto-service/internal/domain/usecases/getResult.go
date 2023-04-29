package usecases

import (
	"context"
	"fmt"
	"proto-service/internal/domain/models"
	"time"
	"unsafe"

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

	// run the serialization process numberIterations times, get the average execution time
	start := time.Now()
	for i := 0; i < numberIterations; i++ {
		var student models.Student
		models.NewStudent(&student)
		data, err = models.SerializeProto(&student)
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
		err = models.DeserializeProto(data, &student)
		if err != nil {
			log.Error("deserialization failed: " + err.Error())
			return "", err
		}
	}
	deserTime := time.Since(start) / 1000

	// print the result in the format: <SerFormat> - <SerStructureSizeInBytes> - <SerTime> - <DeserTime>
	res := fmt.Sprintf("PROTO - %d - %d.%dµs - %d.%dµs", unsafe.Sizeof(data), serTime/1000, serTime%1000, deserTime/1000, deserTime%1000)
	return res, err
}
