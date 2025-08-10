package utils

import (
	"sync"

	"github.com/bwmarrin/snowflake"
)

var (
	node     *snowflake.Node
	nodeOnce sync.Once
)

// 雪花id
func GenerateSnowflakeID(nodeId int64) (string, error) {
	var err error
	nodeOnce.Do(func() {
		node, err = snowflake.NewNode(nodeId)
	})
	if err != nil {
		return "", err
	}
	uuid := node.Generate().String()
	return uuid, nil
}
