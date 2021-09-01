package internal

import (
	filev1 "code.citik.ru/back/report-action/internal/grpcclient/gen/citilink/cmsfiles/file/v1"
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
)

const (
	chunkMaxSize int = 1024 * 1024 * 2  // 2Mb
	fileMaxSize  int = 1024 * 1024 * 30 // 30Mb
)

// FileUploader отправитель файлов
type FileUploader interface {
	Upload(ctx context.Context, reader io.Reader, reportId ReportId) error
}

// fileUploader отправитель файлов
type fileUploader struct {
	client filev1.FileAPIClient
}

// NewFileUploader создает новый отправитель файлов
func NewFileUploader(client filev1.FileAPIClient) *fileUploader {
	return &fileUploader{
		client: client,
	}
}

// Upload выполняет отправку файла потоком
func (f *fileUploader) Upload(ctx context.Context, reader io.Reader, reportId ReportId) error {
	uploadId := uuid.UUID(reportId).String()

	buffer := make([]byte, chunkMaxSize)
	fileSize := 0
	client, err := f.client.Upload(ctx)
	if err != nil {
		return fmt.Errorf("can't start upload: %w", err)
	}

	for {
		n, errRead := reader.Read(buffer)
		if errRead != nil {
			if errRead == io.EOF {
				break
			}
			return fmt.Errorf("can't read chunk to buffer: %w", errRead)
		}

		fileSize += n
		if fileSize > fileMaxSize {
			return fmt.Errorf("file size exceeded")
		}

		err = client.Send(&filev1.UploadRequest{
			Id:    uploadId,
			Chunk: buffer[:n],
		})
		if err != nil {
			return fmt.Errorf("can't send chunk to server: %w", err)
		}
	}

	_, err = client.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("can't close and receive: %w", err)
	}

	return nil
}
