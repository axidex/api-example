package ton

import (
	"encoding/base64"
	"fmt"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

func CreateCell(message string) (string, error) {
	c := cell.BeginCell()

	if err := c.StoreUInt(0, 32); err != nil {
		return "", err
	}

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

	_, err := slice.LoadUInt(32)
	if err != nil {
		return "", fmt.Errorf("failed to read op code: %w", err)
	}

	message, err := slice.LoadStringSnake()
	if err != nil {
		return "", fmt.Errorf("failed to read message: %w", err)
	}

	return message, nil
}
