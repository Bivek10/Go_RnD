package middlewares

import (
	"github.com/bivek/fmt_backend/errors"
	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/responses"
	"github.com/bivek/fmt_backend/services"
	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleWare struct {
	jwtService services.JWTService 
	logger     infrastructure.Logger
	env        infrastructure.Env
	db         infrastructure.Database
}

func NewJWTAuthMiddleWare(
	jwtService services.JWTService,
	logger infrastructure.Logger,
	env infrastructure.Env,
	db infrastructure.Database,

) JWTAuthMiddleWare {
	return JWTAuthMiddleWare{
		jwtService: jwtService,
		logger:     logger,
		env:        env,
		db:         db,
	}
}

// Authenticate user with jwt using this middleware
func (m JWTAuthMiddleWare) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ok, err := m.jwtService.VerifyToken(c)
		if !ok || err != nil {
			m.logger.Zap.Error("Error verifying auth token: ", err.Error())
			err = errors.Unauthorized.Wrap(err, "Something went wrong")
			responses.HandleError(c, err)
			c.Abort()
			return
		}
		c.Next()
	}
}
