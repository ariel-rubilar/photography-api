package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

func GoogleAuthMiddleware(googleClientID string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		auth := ctx.GetHeader("X-ID-TOKEN")
		if auth == "" {
			ctx.Error(
				httperror.Wrap(
					errors.New("authorization header is required"),
					http.StatusUnauthorized, httperror.WithMessage("authorization"),
				),
			)
			ctx.Abort()
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.Error(
				httperror.Wrap(
					errors.New("authorization header format must be Bearer {token}"),
					http.StatusUnauthorized, httperror.WithMessage("authorization"),
				),
			)
			ctx.Abort()
			return
		}

		token := parts[1]

		payload, err := idtoken.Validate(
			context.Background(),
			token,
			googleClientID,
		)
		if err != nil {
			ctx.Error(
				httperror.Wrap(
					err,
					http.StatusUnauthorized, httperror.WithMessage("authorization"),
				),
			)
			ctx.Abort()
			return
		}

		_, ok := payload.Claims["email"].(string)
		if !ok {
			return
		}

		ctx.Next()
	}
}
