package ton

import (
	"fmt"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

func CreateCell(message string) (string, error) {
	c := cell.BeginCell()

	if err := c.StoreStringSnake(message); err != nil {
		return "", err
	}

	result := c.EndCell()

	return string(result.ToBOC()), nil
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

func AnalyzePayload(body *cell.Cell) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	if body == nil {
		result["type"] = "empty"
		return result, nil
	}

	slice := body.BeginParse()

	// Читаем op code
	opCode, err := slice.LoadUInt(32)
	if err != nil {
		return nil, fmt.Errorf("failed to read op code: %w", err)
	}

	result["op_code"] = opCode

	switch opCode {
	case 0:
		// Текстовый комментарий
		message, err := slice.LoadStringSnake()
		if err == nil {
			result["type"] = "text_comment"
			result["message"] = message
		}
	default:
		// Попытка прочитать как числовое значение
		if slice.BitsLeft() >= 32 {
			value, err := slice.LoadUInt(32)
			if err == nil {
				result["type"] = "numeric"
				result["value"] = value
			}
		}
	}

	return result, nil
}
