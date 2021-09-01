package internal

import (
	_ "github.com/golang/mock/mockgen/model"
)

//go:generate mockgen -destination=./mock/queue_mock.go -package=mock github.com/tarantool/go-tarantool/queue Queue
//go:generate mockgen -destination=./mock/cititarantoolclient.go -package=mock code.citik.ru/gobase/tarantool Client
//go:generate mockgen -destination=./mock/inserter.go -package=mock code.citik.ru/back/report-action/internal/distributed_work Inserter
//go:generate mockgen -destination=./mock/cmsfiles.go -package=mock code.citik.ru/back/report-action/internal/grpcclient/gen/citilink/cmsfiles/file/v1 FileAPIClient
