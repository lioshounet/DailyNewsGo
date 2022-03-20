package helper

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

/**
@TODO: 目前只用了真随机数，后续需要改为发号器模式进行生成
*/
func GenUid() (uint64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(100000))

	if err != nil {
		return 0, errors.New(fmt.Sprintf("gen uid fail err:%v", err))
	}

	return n.Uint64(), nil
}
