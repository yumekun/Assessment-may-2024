package request_processor

import (
	"context"

	"mutasi-service/utils/errs"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func (requestProcessor *RequestProcessor) getStreams(ctx context.Context, stream string, count int) ([]redis.XStream, error) {
	const op errs.Op = "request_processor/getStreams"

	// call store layer
	streams, err := requestProcessor.store.redis.GetStreams(ctx, stream, count)
	if err != nil {
		requestProcessor.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "GetStreams",
			"err":   err.Error(),
		}).Error("error!")

		return nil, err
	}

	return streams, nil
}
