package main

import (
	"fmt"
	"net/http"
	"os"
)

type ErrorLevel int // ログのレベルのEnum

const LOG_INFO_ENV_NAME = "LOGGING_INFO"
const (
	Info    ErrorLevel = iota // Info（デフォルトでは出力しない）
	Debug                     // Debug
	Warning                   // Warning
	Fatal                     // Fatal（ログ出力後に終了）
)

// ErrorLevelを対応する文字列に変換する
func (el ErrorLevel) String() string {
	switch el {
	case Info:
		return "Info"
	case Debug:
		return "Debug"
	case Warning:
		return "Warning"
	case Fatal:
		return "Fatal"
	default:
		return ""
	}
}

// errorのinterfaceを満たす専用エラー型
type HTTPError struct {
	StatusCode    int
	Level         ErrorLevel
	internalError error
}

func (he HTTPError) Error() string {
	return fmt.Sprintf(
		"Level: %s, Status: %d %s, Message: %s",
		he.Level.String(),
		he.StatusCode,
		http.StatusText(he.StatusCode),
		he.internalError.Error(),
	)
}

func (he HTTPError) Unwrap() error {
	return he.internalError
}

// HTTPErrorのコンストラクタ
func NewHTTPError(statusCode int, err error) *HTTPError {
	e := new(HTTPError)
	e.internalError = err
	e.StatusCode = statusCode
	switch int(statusCode / 100) {
	case 2:
		e.Level = Info
	case 4:
		e.Level = Info
	case 5:
		e.Level = Warning
	default:
		e.Level = Info
	}
	return e
}

// ロガー
func LoggingHTTPError(statusCode int, err error) {
	// logging Info or not
	loggingInfo := os.Getenv(LOG_INFO_ENV_NAME)
	isInfo := loggingInfo != ""

	he := NewHTTPError(statusCode, err)

	switch he.Level {
	case Info:
		if isInfo {
			fmt.Println(he.Error())
		}
	case Debug:
		fmt.Println(he.Error())
	case Warning:
		fmt.Fprintln(os.Stderr, he.Error())
	case Fatal:
		fmt.Fprintln(os.Stderr, he.Error())
		// When Fatal, exit
		os.Exit(1)
	}
}
