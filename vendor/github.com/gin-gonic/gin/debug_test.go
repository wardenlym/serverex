// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO
// func debugRoute(httpMethod, absolutePath string, handlers HandlersChain) {
// func debugPrint(format string, values ...interface{}) {

func TestIsDebugging(t *testing.T) {
	SetMode(DebugMode)
	assert.True(t, IsDebugging())
	SetMode(ReleaseMode)
	assert.False(t, IsDebugging())
	SetMode(TestMode)
	assert.False(t, IsDebugging())
}

func TestDebugPrint(t *testing.T) {
	var w bytes.Buffer
	setup(&w)
	defer teardown()

	SetMode(ReleaseMode)
	debugPrint("DEBUG this!")
	SetMode(TestMode)
	debugPrint("DEBUG this!")
	assert.Empty(t, w.String())

	SetMode(DebugMode)
	debugPrint("these are %d %s\n", 2, "error messages")
	assert.Equal(t, "[GIN-debug] these are 2 error messages\n", w.String())
}

func TestDebugPrintError(t *testing.T) {
	var w bytes.Buffer
	setup(&w)
	defer teardown()

	SetMode(DebugMode)
	debugPrintError(nil)
	assert.Empty(t, w.String())

	debugPrintError(errors.New("this is an error"))
	assert.Equal(t, "[GIN-debug] [ERROR] this is an error\n", w.String())
}

func TestDebugPrintRoutes(t *testing.T) {
	var w bytes.Buffer
	setup(&w)
	defer teardown()

	debugPrintRoute("GET", "/path/to/route/:param", HandlersChain{func(c *Context) {}, handlerNameTest})
	assert.Regexp(t, `^\[GIN-debug\] GET    /path/to/route/:param     --> (.*/vendor/)?github.com/gin-gonic/gin.handlerNameTest \(2 handlers\)\n$`, w.String())
}

func TestDebugPrintLoadTemplate(t *testing.T) {
	var w bytes.Buffer
	setup(&w)
	defer teardown()

	templ := template.Must(template.New("").Delims("{[{", "}]}").ParseGlob("./testdata/template/hello.tmpl"))
	debugPrintLoadTemplate(templ)
	assert.Regexp(t, `^\[GIN-debug\] Loaded HTML Templates \(2\): \n(\t- \n|\t- hello\.tmpl\n){2}\n`, w.String())
}

func TestDebugPrintWARNINGSetHTMLTemplate(t *testing.T) {
	var w bytes.Buffer
	setup(&w)
	defer teardown()

	debugPrintWARNINGSetHTMLTemplate()
	assert.Equal(t, "[GIN-debug] [WARNING] Since SetHTMLTemplate() is NOT thread-safe. It should only be called\nat initialization. ie. before any route is registered or the router is listening in a socket:\n\n\trouter := gin.Default()\n\trouter.SetHTMLTemplate(template) // << good place\n\n", w.String())
}

func TestDebugPrintWARNINGDefault(t *testing.T) {
	var w bytes.Buffer
	setup(&w)
	defer teardown()

	debugPrintWARNINGDefault()
	assert.Equal(t, "[GIN-debug] [WARNING] Now Gin requires Go 1.6 or later and Go 1.7 will be required soon.\n\n[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.\n\n", w.String())
}

func TestDebugPrintWARNINGNew(t *testing.T) {
	var w bytes.Buffer
	setup(&w)
	defer teardown()

	debugPrintWARNINGNew()
	assert.Equal(t, "[GIN-debug] [WARNING] Running in \"debug\" mode. Switch to \"release\" mode in production.\n - using env:\texport GIN_MODE=release\n - using code:\tgin.SetMode(gin.ReleaseMode)\n\n", w.String())
}

func setup(w io.Writer) {
	SetMode(DebugMode)
	log.SetOutput(w)
}

func teardown() {
	SetMode(TestMode)
	log.SetOutput(os.Stdout)
}
