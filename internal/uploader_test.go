package internal

//go:generate mockgen -destination=./mock/cmsfiles_upload_client.go -package=mock code.citik.ru/back/report-action/internal/grpcclient/gen/citilink/cmsfiles/file/v1 FileAPI_UploadClient

import (
	"bytes"
	"code.citik.ru/back/report-action/internal/mock"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"testing"
)

type ErrorReader struct {
}

func (ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("can't read")
}

type EOFErrorReader struct {
}

func (EOFErrorReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

func TestUploader_Upload(t *testing.T) {
	t.Run("can't start upload", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.Background()

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockCmsFilesClient := mock.NewMockFileAPIClient(ctrl)
		uploader := NewFileUploader(mockCmsFilesClient)

		mockCmsFilesClient.EXPECT().Upload(ctx).Return(nil, errors.New("some error"))

		err := uploader.Upload(ctx, bytes.NewReader([]byte("1")), ReportId(id))
		assert.EqualError(t, err, "can't start upload: some error")
	})
	t.Run("can't read chunk to buffer", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.Background()

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockCmsFilesClient := mock.NewMockFileAPIClient(ctrl)
		uploader := NewFileUploader(mockCmsFilesClient)

		mockCmsFilesClient.EXPECT().Upload(ctx).Return(nil, nil)

		err := uploader.Upload(ctx, ErrorReader{}, ReportId(id))
		assert.EqualError(t, err, "can't read chunk to buffer: can't read")
	})
	t.Run("eof receive error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.Background()

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockCmsFilesClient := mock.NewMockFileAPIClient(ctrl)
		mockCmsUploadClient := mock.NewMockFileAPI_UploadClient(ctrl)
		uploader := NewFileUploader(mockCmsFilesClient)

		mockCmsFilesClient.EXPECT().Upload(ctx).Return(mockCmsUploadClient, nil)
		mockCmsUploadClient.EXPECT().CloseAndRecv().Return(nil, errors.New("some error"))

		err := uploader.Upload(ctx, EOFErrorReader{}, ReportId(id))
		assert.EqualError(t, err, "can't close and receive: some error")
	})
	t.Run("eof success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.Background()

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockCmsFilesClient := mock.NewMockFileAPIClient(ctrl)
		mockCmsUploadClient := mock.NewMockFileAPI_UploadClient(ctrl)
		uploader := NewFileUploader(mockCmsFilesClient)

		mockCmsFilesClient.EXPECT().Upload(ctx).Return(mockCmsUploadClient, nil)
		mockCmsUploadClient.EXPECT().CloseAndRecv().Return(nil, nil)

		err := uploader.Upload(ctx, EOFErrorReader{}, ReportId(id))
		assert.NoError(t, err)
	})
	t.Run("max filesize", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.Background()

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockCmsFilesClient := mock.NewMockFileAPIClient(ctrl)
		mockCmsUploadClient := mock.NewMockFileAPI_UploadClient(ctrl)
		uploader := NewFileUploader(mockCmsFilesClient)

		mockCmsFilesClient.EXPECT().Upload(ctx).Return(mockCmsUploadClient, nil)
		mockCmsUploadClient.EXPECT().Send(gomock.Any()).Return(nil).Times(15)

		buf := make([]byte, fileMaxSize+1)
		rand.Read(buf)

		err := uploader.Upload(ctx, bytes.NewReader(buf), ReportId(id))
		assert.EqualError(t, err, "file size exceeded")
	})
	t.Run("can't send chunk to server", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		ctx := context.Background()

		id, _ := uuid.Parse("40d206a2-06d1-46cb-9c05-e3540c1eb0cf")
		mockCmsFilesClient := mock.NewMockFileAPIClient(ctrl)
		mockCmsUploadClient := mock.NewMockFileAPI_UploadClient(ctrl)
		uploader := NewFileUploader(mockCmsFilesClient)

		mockCmsFilesClient.EXPECT().Upload(ctx).Return(mockCmsUploadClient, nil)
		mockCmsUploadClient.EXPECT().Send(gomock.Any()).Return(errors.New("some error"))

		err := uploader.Upload(ctx, bytes.NewReader([]byte("1")), ReportId(id))
		assert.EqualError(t, err, "can't send chunk to server: some error")
	})
}
