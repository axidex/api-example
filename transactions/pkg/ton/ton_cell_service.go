package ton

import (
	"encoding/base64"
	"fmt"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

func CreateCell(message string) (string, error) {
	c := cell.BeginCell()

	if err := c.StoreStringSnake(message); err != nil {
		return "", err
	}

	result := c.EndCell()

	return base64.StdEncoding.EncodeToString(result.ToBOC()), nil
}

func DecodeStringPayload(body *cell.Cell) (string, error) {
	if body == nil {
		return "", nil
	}

	slice := body.BeginParse()

	message, err := slice.LoadStringSnake()
	if err != nil {
		return "", fmt.Errorf("failed to read message: %w", err)
	}

	return message, nil
}
