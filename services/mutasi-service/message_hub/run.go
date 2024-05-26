package request_processor

import (
	"context"

	"mutasi-service/utils/errs"

	"github.com/sirupsen/logrus"
)

func (requestProcessor *RequestProcessor) Run(ctx context.Context) {
	const op errs.Op = "request_processor/Run"

	// start a goroutine to listen for the message
	for {
		select {
		case <-ctx.Done():
			{
				requestProcessor.logger.WithFields(logrus.Fields{
					"op": op,
				}).Debug("context was done!")

				return
			}
		default:
			{
				// create new context everytime core logic executed
				ctx := context.Background()

				// execute core logic
				err := requestProcessor.consume(ctx)
				if err != nil {
					requestProcessor.logger.WithFields(logrus.Fields{
						"op":    op,
						"scope": "consume",
						"err":   err.Error(),
					}).Error("error!")
				}
			}
		}
	}
}
