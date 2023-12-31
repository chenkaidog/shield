package csrf

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"net/textproto"
	"shield/common/utils/random"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

// New validates CSRF token.
func New(opts ...Option) app.HandlerFunc {
	cfg := NewOptions(opts...)
	selectors := strings.Split(cfg.KeyLookup, ":")

	if len(selectors) != 2 {
		panic(errors.New("[CSRF] KeyLookup must in the form of <source>:<key>"))
	}

	if cfg.Extractor == nil {
		// By default, we extract from a header
		cfg.Extractor = CsrfFromHeader(textproto.CanonicalMIMEHeaderKey(selectors[1]))

		switch selectors[0] {
		case "form":
			cfg.Extractor = CsrfFromForm(selectors[1])
		case "query":
			cfg.Extractor = CsrfFromQuery(selectors[1])
		case "param":
			cfg.Extractor = CsrfFromParam(selectors[1])
		}
	}

	return func(ctx context.Context, c *app.RequestContext) {
		c.Set(csrfSecret, cfg.Secret)

		// Don't execute middleware if Next returns true
		if cfg.Next != nil && cfg.Next(ctx, c) {
			c.Next(ctx)
			return
		}

		if isIgnored(cfg.IgnoreMethods, string(c.Request.Method())) {
			c.Next(ctx)
			return
		}

		session := sessions.Default(c)
		salt, ok := session.Get(csrfSalt).(string)
		if !ok || len(salt) == 0 {
			c.Error(errMissingSalt)
			cfg.ErrorFunc(ctx, c)
			return
		}

		token, err := cfg.Extractor(ctx, c)
		if err != nil {
			c.Error(err)
			cfg.ErrorFunc(ctx, c)
			return
		}

		if tokenize(cfg.Secret, salt) != token {
			c.Error(errInvalidToken)
			cfg.ErrorFunc(ctx, c)
			return
		}

		c.Next(ctx)
	}
}

// GetToken returns a CSRF token.
func GetToken(c *app.RequestContext) string {
	if t, ok := c.Get(csrfToken); ok {
		return t.(string)
	}

	session := sessions.Default(c)
	salt, ok := session.Get(csrfSalt).(string)
	if !ok {
		salt = random.RandStr(16)
		session.Set(csrfSalt, salt)
		session.Save()
	}

	secret := c.MustGet(csrfSecret).(string)
	token := tokenize(secret, salt)
	c.Set(csrfToken, token)

	return token
}

// tokenize generates token through secret and salt.
func tokenize(secret, salt string) string {
	h := sha256.New()
	io.WriteString(h, salt+"-"+secret)
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return hash
}

// isIgnored determines whether the method is ignored.
func isIgnored(arr []string, value string) bool {
	ignore := false

	for _, v := range arr {
		if v == value {
			ignore = true
			break
		}
	}

	return ignore
}
