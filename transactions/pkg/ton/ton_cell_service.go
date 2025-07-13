package ton

import (
	"github.com/xssnick/tonutils-go/tvm/cell"
)

func CreateCell(message string) (string, error) {
	// Создаем новую пустую ячейку
	c := cell.BeginCell()

	if err := c.StoreUInt(0, 32); err != nil {
		return "", err
	}

	if err := c.StoreStringSnake(message); err != nil {
		return "", err
	}

	result := c.EndCell()

	return string(result.ToBOC()), nil
}
