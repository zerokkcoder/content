package main

import (
	"fmt"
	"hash/fnv"
	"math/big"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

const contentNumTables = 4

func main() {
	u := uuid.New().String()
	getContentDetailTable(u)
}

func getContentDetailTable(contentID string) string {
	tableIndex := getContentTableIndex(contentID)
	table := fmt.Sprintf("cms_content.t_content_details_%d", tableIndex)
	log.Infof("content_id = %s, table = %s", contentID, table)
	return table
}

func getContentTableIndex(uuid string) int {
	hash := fnv.New64()
	hash.Write([]byte(uuid))
	hashValue := hash.Sum64()
	fmt.Println("hashValue = ", hashValue)

	bigNum := big.NewInt(int64(hashValue))
	mod := big.NewInt(contentNumTables)
	tableIndex := bigNum.Mod(bigNum, mod).Int64()
	return int(tableIndex)
}
