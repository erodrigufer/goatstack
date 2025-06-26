package web

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/google/uuid"
)

// ErrorMessageBody the standard data type used to deliver error messages.
type ErrorMessageBody struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Details string `json:"details"`
	} `json:"error"`
}

type stackFrame struct {
	Function string `json:"function"`
	File     string `json:"file"`
	Line     int    `json:"line"`
}

// NewErrorMessageBody returns a body for an error message that can be sent with
// SendJSONResponse.
func NewErrorMessageBody(code, message, details string) ErrorMessageBody {
	body := ErrorMessageBody{
		Error: struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Details string `json:"details"`
		}{
			Code:    code,
			Message: message,
			Details: details,
		},
	}

	return body
}

// getErrorStackFrame returns the stackframe with all the functions
// thar were called until an error happened.
func getErrorStackFrame(directCall bool) []stackFrame {
	pcs := make([]uintptr, 10)
	var skippedFrames int
	if directCall {
		skippedFrames = 3
	} else {
		skippedFrames = 4
	}
	runtime.Callers(skippedFrames, pcs)

	var stackFrames []stackFrame

	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		var sf stackFrame
		sf.Line = frame.Line
		sf.Function = frame.Function
		sf.File = filepath.Base(frame.File)
		stackFrames = append(stackFrames, sf)
		if !more {
			break
		}
	}
	return stackFrames
}

// LogErrorsWithStack logs an error and displays the stack straces of all the functions
// called before the error took place.
// LogErrorsWithStack does not log context.Canceled errors.
// `msg` is the message that should be logged in the record (see log/slog practices) and `directCall`
// should be true if LogErrorsWithStack was directly called to log the errors (if LogErrorsWithStack is
// being called within another function, this value should be false).
func LogErrorsWithStack(ctx context.Context, msg string, directCall bool, callerError error, errorLogger *slog.Logger) {
	// Do not log context.Canceled errors.
	if errors.Is(callerError, context.Canceled) {
		return
	}
	stackFrames := getErrorStackFrame(directCall)

	var errorFile slog.Attr
	var errorFileValue string
	var errorLine slog.Attr
	var errorLineValue int
	var errorUUID slog.Attr
	var errorUUIDValue string
	var slogFrames slog.Attr

	if stackFrames != nil {
		slogFrames = slog.Any("traces", stackFrames)
		// Only initialize this logs if stackFrames != nil,
		// otherwise reading the first value of stackFrames slice
		// could panic.
		errorFileValue = stackFrames[0].File
		errorLineValue = stackFrames[0].Line
		errorFile = slog.String("error_file", errorFileValue)
		errorLine = slog.Int("error_line", errorLineValue)
	} else {
		// If for some reason no stack frames could be gathered.
		slogFrames = slog.String("traces", "no information available")
	}

	uuid, err := uuid.NewV7()
	// If an error happened, default values will be used.
	if err == nil {
		errorUUIDValue = uuid.String()
		errorUUID = slog.String("error_uuid", errorUUIDValue)
	}

	errorLogger.Error(msg, errorUUID, slog.String("error_message", callerError.Error()), errorFile, errorLine, slogFrames)
}

// HandleServerError sends an error message and stack trace to the error logger and
// then sends a generic 500 Internal Server Error response to the client.
func HandleServerError(w http.ResponseWriter, r *http.Request, err error, errorLogger *slog.Logger) {
	LogErrorsWithStack(r.Context(), "an internal server error happened", false, err, errorLogger)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func SendHTMXErrorMessage(w http.ResponseWriter, r *http.Request, statusCode int, errorLogger *slog.Logger, errorMessage string) {
	h := w.Header()
	h.Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := fmt.Fprintln(w, errorMessage)
	if err != nil {
		HandleServerError(w, r, fmt.Errorf("could not write to http.ResponseWriter: %w", err), errorLogger)
	}
}
