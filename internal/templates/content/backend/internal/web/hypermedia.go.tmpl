package web

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

const initialBufferSize = 4096

func RenderComponent(ctx context.Context, w http.ResponseWriter, component templ.Component) error {
	buf := bytes.NewBuffer(make([]byte, 0, initialBufferSize))
	err := component.Render(ctx, buf)
	if err != nil {
		return fmt.Errorf("an error happened while rendering a component: %w", err)
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("an error happened while writing a view to http.ResponseWriter: %w", err)
	}
	return nil
}
