package store

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"time"

	"github.com/dewadg/go-playground-api/internal/adapter"
	"github.com/sirupsen/logrus"
)

type CreateItemRequest struct {
	Input []string
}

type UpdateItemRequest struct {
	Input []string
}

type CreateItemResponse struct {
	ShareID string
}

type UpdateItemResponse struct {
	ShareID string
}

// TODO: implement queue on this
func CreateItem(ctx context.Context, payload CreateItemRequest) (CreateItemResponse, error) {
	db, err := adapter.GetMysqlDB()
	if err != nil {
		return CreateItemResponse{}, err
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return CreateItemResponse{}, err
	}

	id, err := popID(ctx)
	if err != nil {
		return CreateItemResponse{}, err
	}

	query := `UPDATE items SET data = ?, updated_at = NOW(), assigned_at = NOW() WHERE id = ?`
	_, err = db.ExecContext(ctx, query, string(payloadBytes), id)
	if err != nil {
		return CreateItemResponse{}, err
	}

	return CreateItemResponse{
		ShareID: id,
	}, nil
}

type Item struct {
	ID    string
	Input []string
}

func FindItem(ctx context.Context, id string) (Item, error) {
	cachedItem, err := findItemFromRedis(ctx, id)
	if err == nil {
		return cachedItem, nil
	}

	item, err := findItemFromMysql(ctx, id)
	if err != nil {
		return Item{}, err
	}

	err = cacheItemOnRedis(ctx, item)
	if err != nil {
		logrus.WithError(err).Error("store.FindItem: failed to cache item")
	}

	return item, nil
}

func findItemFromMysql(ctx context.Context, id string) (Item, error) {
	db, err := adapter.GetMysqlDB()
	if err != nil {
		return Item{}, err
	}

	query := `SELECT data FROM items WHERE id = ? LIMIT 1`

	var dataBytes []byte
	if err = db.QueryRowContext(ctx, query, id).Scan(&dataBytes); err != nil {
		return Item{}, err
	}

	var item Item
	if err = json.Unmarshal(dataBytes, &item); err != nil {
		return Item{}, err
	}

	item.ID = id

	return item, nil
}

var cacheWriteBuff bytes.Buffer
var cacheWriteEnc = gob.NewEncoder(&cacheWriteBuff)

func cacheItemOnRedis(ctx context.Context, payload Item) error {
	client, err := adapter.GetRedisClient()
	if err != nil {
		return err
	}

	err = cacheWriteEnc.Encode(payload)
	if err != nil {
		return err
	}

	key := "items." + payload.ID
	err = client.Set(ctx, key, cacheWriteBuff.Bytes(), 1*time.Minute).Err()
	cacheWriteBuff.Reset()

	return err
}

var cacheReadBuff bytes.Buffer
var cacheReadDecoder = gob.NewDecoder(&cacheReadBuff)

func findItemFromRedis(ctx context.Context, id string) (Item, error) {
	client, err := adapter.GetRedisClient()
	if err != nil {
		return Item{}, err
	}

	key := "items." + id
	result, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return Item{}, err
	}

	_, err = cacheReadBuff.Write(result)
	if err != nil {
		return Item{}, err
	}
	cacheReadBuff.Reset()

	var item Item
	err = cacheReadDecoder.Decode(&item)

	return item, err
}

func removeItemFromRedis(ctx context.Context, id string) error {
	client, err := adapter.GetRedisClient()
	if err != nil {
		return err
	}

	key := "items." + id
	err = client.Del(ctx, key).Err()

	return err
}

func UpdateItem(ctx context.Context, shareID string, payload UpdateItemRequest) (Item, error) {
	db, err := adapter.GetMysqlDB()
	if err != nil {
		return Item{}, err
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return Item{}, err
	}

	query := `UPDATE items SET data = ?, updated_at = NOW() WHERE id = ?`
	_, err = db.ExecContext(ctx, query, string(payloadBytes), shareID)
	if err != nil {
		return Item{}, err
	}

	if err = removeItemFromRedis(ctx, shareID); err != nil {
		logrus.WithError(err).Error("error while removing item from redis after update")
	}

	return Item{
		ID:    shareID,
		Input: payload.Input,
	}, nil
}
