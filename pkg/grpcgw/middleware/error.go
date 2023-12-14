package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func ErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	if md, ok := runtime.ServerMetadataFromContext(ctx); ok {
		if code := md.HeaderMD.Get("x-http-code"); len(code) > 0 {
			if status, _ := strconv.Atoi(code[0]); status >= 400 {
				err = &runtime.HTTPStatusError{
					HTTPStatus: status,
					Err:        err,
				}
			}
			delete(w.Header(), "Grpc-Metadata-X-Http-Code")
		}
	}
	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}
