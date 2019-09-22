package main

import "testing"

import "github.com/gin-gonic/gin"

func TestTestHeaderMiddleware(t *testing.T) {
	t.Run("checkHeader", func(t *testing.T) {
		type testContext struct {
			nextHandlerExecuted, stdouted bool
		}
		var context testContext
		middleware := newTestHeaderMiddleware()
		middleware.checkHeaderDeps.ctxHeaderGetter = aCtxHeaderGetter(func(*gin.Context, string) (_ string) {
			return
		})
		middleware.checkHeaderDeps.ctxNexter = aCtxNexter(func(*gin.Context) {
			context.nextHandlerExecuted = true
		})
		middleware.checkHeaderDeps.printlner = aPrintlner(func(a ...interface{}) (n int, _ error) {
			context.stdouted = true
			return
		})
		middleware.checkHeader(nil)
		if !(!context.stdouted && context.nextHandlerExecuted) {
			t.Fatal("❌")
		}
	})
}
