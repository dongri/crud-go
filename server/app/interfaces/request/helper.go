package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func readJSON(w http.ResponseWriter, r *http.Request, v interface{}) error {
	ct := r.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		msg := "Content-Type must be application/json: got %s"
		err := fmt.Errorf(msg, ct)
		return err
	}
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		switch {
		case errors.As(err, &syntaxError):
			err = fmt.Errorf("invalid json syntax: %w", err)
		case errors.As(err, &unmarshalTypeError):
			err = fmt.Errorf("invalid json field: %w", err)
		case errors.Is(err, io.EOF):
			err = fmt.Errorf("request body is empty: %w", err)
		case errors.Is(err, io.ErrUnexpectedEOF):
			err = fmt.Errorf("invalid json syntax: %w", err)
		default:
			err = fmt.Errorf("failed to decode json: %w", err)
			// エラー内容のログ出力は割愛
		}
		return err
	}
	return nil
}
