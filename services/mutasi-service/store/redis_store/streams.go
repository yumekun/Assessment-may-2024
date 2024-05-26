package redis_store

import (
	"context"

	"mutasi-service/utils/errs"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func (store *RedisStore) AddToStream(ctx context.Context, stream string, values interface{}) error {
	const op errs.Op = "redis_store/AddToStream"

	// XAdd command
	cmd := store.client.XAdd(ctx, &redis.XAddArgs{
		Stream: stream,
		Values: values,
	})

	// execute the command
	result, err := cmd.Result()
	if err != nil {
		store.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "XAdd",
			"err":   err.Error(),
		}).Error("error!")

		return err
	}

	// log the result for data tracing purpose
	store.logger.WithFields(logrus.Fields{
		"op":     op,
		"result": result,
	}).Debug("result!")

	return nil
}

func (store *RedisStore) DeleteFromStream(ctx context.Context, stream string, messageID string) error {
	const op errs.Op = "redis_store/DeleteFromStream"

	// XDel command
	cmd := store.client.XDel(ctx, stream, messageID)

	// execute the command
	result, err := cmd.Result()
	if err != nil {
		store.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "XDel",
			"err":   err.Error(),
		}).Error("error!")

		return err
	}

	// log the result for data tracing purpose
	store.logger.WithFields(logrus.Fields{
		"op":     op,
		"result": result,
	}).Debug("result!")

	return nil
}

func (store *RedisStore) GetFromStream(ctx context.Context, stream string, count int) ([]redis.XMessage, error) {
	const op errs.Op = "redis_store/GetFromStream"

	// XRead command
	cmd := store.client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, "0"},
		Count:   int64(count),
		Block:   0,
	})

	// execute the command
	streams, err := cmd.Result()
	if err != nil {
		store.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "XRead",
			"err":   err.Error(),
		}).Error("error!")

		return nil, err
	}

	// Check if any message was received
	if len(streams) > 0 && len(streams[0].Messages) > 0 {
		messages := streams[0].Messages

		// log the stream Messages for data tracing purpose
		store.logger.WithFields(logrus.Fields{
			"op":       op,
			"messages": messages,
		}).Debug("messages!")

		return messages, nil
	}

	return nil, nil
}

func (store *RedisStore) GetStreams(ctx context.Context, stream string, count int) ([]redis.XStream, error) {
	const op errs.Op = "redis_store/GetFromStream"

	// XRead command
	cmd := store.client.XRead(ctx, &redis.XReadArgs{
		Streams: []string{stream, "0"},
		Count:   int64(count),
		Block:   0,
	})

	// execute the command
	streams, err := cmd.Result()
	if err != nil {
		store.logger.WithFields(logrus.Fields{
			"op":    op,
			"scope": "XRead",
			"err":   err.Error(),
		}).Error("error!")

		return nil, err
	}

	// Check if any streams was received
	if len(streams) > 0 {
		return streams, nil
	}

	return nil, nil
}
