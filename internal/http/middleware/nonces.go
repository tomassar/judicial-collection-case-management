package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type key string

var NonceKey key = "nonces"

type Nonces struct {
	Htmx            string
	ResponseTargets string
	Tw              string
	HtmxCSSHash     string
}

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func CSPMiddleware() gin.HandlerFunc {
	// To use the same nonces in all responses, move the Nonces
	// struct creation to here, outside the handler.

	return func(c *gin.Context) {
		// Create a new Nonces struct for every request when here.
		// move to outside the handler to use the same nonces in all responses
		// TODO: (@tomassar) should set better nonces
		nonceSet := Nonces{
			//Htmx:            generateRandomString(16),
			ResponseTargets: generateRandomString(16),
			Tw:              generateRandomString(16),
			//HtmxCSSHash:     "sha256-pgn1TCGZX6O77zDvy0oTODMOxemn0oj0LeCnQTRj7Kg=",
		}

		// set nonces in context
		//ctx := context.WithValue(r.Context(), NonceKey, nonceSet)
		ctx := context.WithValue(c.Request.Context(), NonceKey, nonceSet)
		// insert the nonces into the content security policy header
		/* cspHeader := fmt.Sprintf("default-src 'self'; script-src 'nonce-%s' 'nonce-%s' ; style-src 'nonce-%s' '%s';",
		nonceSet.Htmx,
		nonceSet.ResponseTargets,
		nonceSet.Tw,
		nonceSet.HtmxCSSHash) */
		cspHeader := fmt.Sprintf("default-src 'self'; style-src 'nonce-%s';",
			nonceSet.Tw)
		//w.Header().Set("Content-Security-Policy", cspHeader)
		c.Writer.Header().Set("Content-Security-Policy", cspHeader)

		//next.ServeHTTP(w, r.WithContext(ctx))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetNonces(ctx context.Context) Nonces {
	nonceSet := ctx.Value(NonceKey)
	if nonceSet == nil {
		log.Fatal("error getting nonce set - is nil")
	}

	nonces, ok := nonceSet.(Nonces)

	if !ok {
		log.Fatal("error getting nonce set - not ok")
	}

	return nonces
}

func GetTwNonce(ctx context.Context) string {
	nonceSet := GetNonces(ctx)
	return nonceSet.Tw
}
