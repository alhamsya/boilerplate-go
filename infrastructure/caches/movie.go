package caches

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alhamsya/boilerplate-go/domain/constants"
	"github.com/alhamsya/boilerplate-go/infrastructure/externals/omdb"
	"github.com/alhamsya/boilerplate-go/lib/managers/custom_error"
)

const (
	KeyListMovie   = "movie:list_movie:search:%s:page:%d"
	KeyDetailMovie = "movie:detail:movie_id:%s"
)

func (cache *ServiceCache) SetListMovie(ctx context.Context, search string, page int64, req *omdb.OMDBList) (err error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return customError.Wrap(err, "json marshal")
	}

	err = cache.Redis.SetEX(ctx, fmt.Sprintf(KeyListMovie, search, page), jsonData, constCommon.RedisExpiry1Hour).Err()
	if err != nil {
		return customError.WrapFlag(err, "redis", "SetEX")
	}

	return nil
}

func (cache *ServiceCache) GetListMovie(ctx context.Context, search string, page int64) (resp *omdb.OMDBList, err error) {
	jsonData, err := cache.Redis.Get(ctx, fmt.Sprintf(KeyListMovie, search, page)).Result()
	if err != nil {
		return nil, customError.WrapFlag(err, "redis", "Get")
	}

	err = json.Unmarshal([]byte(jsonData), &resp)
	if err != nil {
		return nil, customError.Wrap(err, "json Unmarshal")
	}

	return resp, nil
}

func (cache *ServiceCache) SetDetailMovie(ctx context.Context, movieID string, req *omdb.OMDBDetail) (err error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return customError.Wrap(err, "json marshal")
	}

	err = cache.Redis.SetEX(ctx, fmt.Sprintf(KeyDetailMovie, movieID), jsonData, constCommon.RedisExpiry1Day).Err()
	if err != nil {
		return customError.WrapFlag(err, "redis", "SetEX")
	}

	return nil
}

func (cache *ServiceCache) GetDetailMovie(ctx context.Context, movieID string) (resp *omdb.OMDBDetail, err error) {
	jsonData, err := cache.Redis.Get(ctx, fmt.Sprintf(KeyDetailMovie, movieID)).Result()
	if err != nil {
		return nil, customError.WrapFlag(err, "redis", "Get")
	}

	err = json.Unmarshal([]byte(jsonData), &resp)
	if err != nil {
		return nil, customError.Wrap(err, "json Unmarshal")
	}

	return resp, nil
}
