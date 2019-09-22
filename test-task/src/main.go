package main

import (
	"fmt"
	"net/http"
)

import "github.com/gin-gonic/gin"

type (
	testHeaderMiddleware struct {
		checkHeaderDeps struct {
			printlner
			ctxHeaderGetter
			ctxNexter
		}
	}

	printlner interface {
		println(...interface{}) (n int, _ error)
	}

	aPrintlner func(a ...interface{}) (n int, _ error)

	ctxHeaderGetter interface {
		getHeaderFromCtx(_ *gin.Context, key string) string
	}

	aCtxHeaderGetter func(_ *gin.Context, key string) string

	ctxNexter interface {
		nextFrom(*gin.Context)
	}

	aCtxNexter func(*gin.Context)
)

func main() {
	router := gin.New()
	router.GET("/", newTestHeaderMiddleware().checkHeader, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%s\n", "hello-world")
	})
	router.Run(":3001")
}

func newTestHeaderMiddleware() (middleware testHeaderMiddleware) {
	middleware.checkHeaderDeps.printlner = aPrintlner(fmt.Println)
	middleware.checkHeaderDeps.ctxHeaderGetter = aCtxHeaderGetter(func(ctx *gin.Context, key string) string {
		return ctx.GetHeader(key)
	})
	middleware.checkHeaderDeps.ctxNexter = aCtxNexter(func(ctx *gin.Context) {
		ctx.Next()
	})
	return
}

func (m testHeaderMiddleware) checkHeader(ctx *gin.Context) {
	d := &m.checkHeaderDeps

	defer d.nextFrom(ctx)
	if d.getHeaderFromCtx(ctx, "test") == "yes" {
		d.println("header is fine")
	}
}

func (p aPrintlner) println(a ...interface{}) (n int, _ error) {
	return p(a...)
}

func (g aCtxHeaderGetter) getHeaderFromCtx(ctx *gin.Context, key string) string {
	return g(ctx, key)
}

func (n aCtxNexter) nextFrom(ctx *gin.Context) {
	n(ctx)
}
