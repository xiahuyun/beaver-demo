package version

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type VersionGenerator struct {
	redisClient *redis.Client
	db          *gorm.DB
}

func NewVersionGenerator(redisClient *redis.Client, db *gorm.DB) *VersionGenerator {
	return &VersionGenerator{
		redisClient: redisClient,
		db:          db,
	}
}

func (vg *VersionGenerator) GetNextVersion(table, field, value string) (int64, error) {
	var key string
	isConditional := field != "" && value != ""
	if isConditional {
		key = fmt.Sprintf("version:%s:%s:%s", table, field, value)
	} else {
		key = fmt.Sprintf("version:%s", table)
	}

	initVersion := int64(1)

	nextVersion, err := vg.redisClient.Incr(context.Background(), key).Result()
	if err != nil {
		logx.Errorf("Redis获取版本号失败: table=%s, field=%s, value=%s, error=%v", table, field, value, err)
		return 0, err
	}

	if nextVersion == initVersion {
		// redis 中 key 不存在或过期
		maxVersion, err := vg.GetMaxVersionFromMysql(table, field, value)
		if err != nil {
			return 0, err
		}
		// 将最大版本号设置到 Redis 中
		err = vg.redisClient.Set(context.Background(), key, maxVersion+1, 0).Err()
		if err != nil {
			logx.Errorf("Redis设置版本号失败: table=%s, field=%s, value=%s, error=%v", table, field, value, err)
			return 0, err
		}
		return maxVersion + 1, nil
	}

	return nextVersion, nil
}

func (vg *VersionGenerator) GetMaxVersionFromMysql(table, field, value string) (int64, error) {
	isConditional := field != "" && value != ""
	tableName := vg.getTableName(table)

	var maxVersion int64
	var err error
	query := fmt.Sprintf("SELECT COALESCE(MAX(version), 0) FROM %s", tableName)
	if isConditional {
		// 有条件查询
	} else {
		err = vg.db.Raw(query).Scan(&maxVersion).Error
	}
	if err != nil {
		logx.Errorf("数据库获取最大版本号失败: table=%s, field=%s, value=%s, error=%v", table, field, value, err)
		return 0, err
	}

	return maxVersion, nil
}

func (vg *VersionGenerator) getTableName(dataType string) string {
	switch dataType {
	case "friend_verify":
		return "friend_verify"
	// 可以在这里添加更多的数据类型和对应的表名
	default:
		return dataType // 默认返回传入的值
	}
}
