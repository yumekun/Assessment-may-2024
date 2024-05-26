package request_processor

import (
	"context"

	"mutasi-service/utils/errs"

	"github.com/sirupsen/logrus"
)

// Consume consumes stream data sent to redis stream
func (requestProcessor *RequestProcessor) consume(ctx context.Context) error {
	const op errs.Op = "request_processor/consume"

	streams, err := requestProcessor.getStreams(ctx, requestProcessor.config.RedisMutasiRequestStream, 1)
	if err != nil {
		requestProcessor.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "getStreams",
			"err":   err.Error(),
		}).Error("error!")

		return err
	}

	for _, stream := range streams {
		for _, message := range stream.Messages {
			// create new context everytime we received message from stream
			ctx := context.TODO()

			// get values from stream message
			values := message.Values

			// delete the message from the stream
			err = requestProcessor.deleteMessage(ctx, requestProcessor.config.RedisMutasiRequestStream, message.ID)
			if err != nil {
				requestProcessor.logger.WithFields(logrus.Fields{
					"op":    op,
					"scope": "deleteMessage",
					"err":   err.Error(),
				}).Error("error!")

				return err
			}

			// log the request stream data value for data tracing purpose
			requestProcessor.logger.WithFields(logrus.Fields{
				"op":             op,
				"message_values": values,
			}).Debug("params!")

			params, err := requestProcessor.service.NewCreateMutasiParamsFromMap(values)
			if err != nil {
				requestProcessor.logger.WithFields(logrus.Fields{
					"op":    op,
					"scope": "NewCreateMutasiParamsFromMap",
					"err":   err.Error(),
				}).Error("error!")

				return err
			}

			_, err = requestProcessor.service.CreateMutasi(ctx, params)
			if err != nil {
				requestProcessor.logger.WithFields(logrus.Fields{
					"op":    op,
					"scope": "CreateMutasi",
					"err":   err.Error(),
				}).Error("error!")

				return err
			}
		}
	}

	return nil
}
