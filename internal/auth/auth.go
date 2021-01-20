package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/autom8ter/meshpaas/internal/logger"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jws"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
)

type ctxKey string

var (
	userCtxKey ctxKey = "user-ctx-key"
)

type Auth struct {
	jwksUri   string
	jwtIssuer string
	jwksSet   *jwk.Set
	mu        sync.RWMutex
	logger    *logger.Logger
}

func NewAuth(jwksUri string, jwtIssuer string, logger2 *logger.Logger) (*Auth, error) {
	a := &Auth{
		jwksUri:   jwksUri,
		jwksSet:   nil,
		mu:        sync.RWMutex{},
		logger:    logger2,
		jwtIssuer: jwtIssuer,
	}
	return a, a.RefreshJWKS()
}

func (a *Auth) RefreshJWKS() error {
	jwks, err := jwk.Fetch(a.jwksUri)
	if err != nil {
		return err
	}
	a.mu.Lock()
	a.jwksSet = jwks
	a.mu.Unlock()
	return nil
}

func (a *Auth) ParseAndVerify(token string) (map[string]interface{}, error) {
	message, err := jws.ParseString(token)
	if err != nil {
		return nil, err
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	if len(message.Signatures()) == 0 {
		return nil, fmt.Errorf("zero jws signatures")
	}
	kid, ok := message.Signatures()[0].ProtectedHeaders().Get("kid")
	if !ok {
		return nil, fmt.Errorf("jws kid not found")
	}
	algI, ok := message.Signatures()[0].ProtectedHeaders().Get("alg")
	if !ok {
		return nil, fmt.Errorf("jw alg not found")
	}
	alg, ok := algI.(jwa.SignatureAlgorithm)
	if !ok {
		return nil, fmt.Errorf("alg type cast error")
	}

	var payload []byte
	if a.jwksSet == nil {
		payload = message.Payload()
	} else {
		keys := a.jwksSet.LookupKeyID(kid.(string))
		if len(keys) == 0 {
			return nil, errors.Errorf("failed to lookup kid: %s - zero keys", kid.(string))
		}
		var key interface{}
		if err := keys[0].Raw(&key); err != nil {
			return nil, err
		}
		payload, err = jws.Verify([]byte(token), alg, key)
		if err != nil {
			return nil, err
		}
	}

	data := map[string]interface{}{}
	if err := json.Unmarshal(payload, &data); err != nil {
		return nil, err
	}
	if len(a.jwtIssuer) > 0 {
		if issuer := data["iss"]; issuer == nil {
			return nil, errors.Errorf("empty jwt.claims.iss issuer")
		} else {
			if a.jwtIssuer != cast.ToString(issuer) {
				return nil, errors.Errorf("unsupported jwt.claims.iss issuer: %s", cast.ToString(issuer))
			}
		}
	}
	return data, nil
}

func (a *Auth) Interceptor() grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		payload, err := a.ParseAndVerify(token)
		if err != nil {
			a.logger.Error(err.Error())
			return nil, status.Error(codes.Unauthenticated, "unverified")
		}
		ctx = context.WithValue(ctx, userCtxKey, payload)
		return ctx, nil
	}
}

func UserContext(ctx context.Context) (map[string]interface{}, bool) {
	if ctx.Value(userCtxKey) == nil {
		return nil, false
	}
	data, ok := ctx.Value(userCtxKey).(map[string]interface{})
	return data, ok
}