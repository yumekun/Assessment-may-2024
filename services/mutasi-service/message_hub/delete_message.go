package request_processor

import (
	"context"

	"mutasi-service/utils/errs"

	"github.com/sirupsen/logrus"
)

func (requestProcessor *RequestProcessor) deleteMessage(ctx context.Context, stream string, messageID string) error {
	const op errs.Op = "request_processor/deleteMessage"

	// call store layer
	err := requestProcessor.store.redis.DeleteFromStream(ctx, stream, messageID)
	if err != nil {
		requestProcessor.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "DeleteFromStream",
			"err":   err.Error(),
		}).Error("error!")

		return err
	}

	return nil
}
