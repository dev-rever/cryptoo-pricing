package router

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dev-rever/cryptoo-pricing/internal/controllers"
	"github.com/dev-rever/cryptoo-pricing/internal/middleware/mredis"
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
)

type Engine struct {
	gin        *gin.Engine
	mredis     *mredis.Wrap
	jwt        func(*gin.Context)
	userCtrl   *controllers.User
	cryptoCtrl *controllers.Crypto
}

const (
	root     = "/"
	readmeEN = "/readme.en"
	readmeZH = "/readme.zh"
	docs     = "/docs"
	swagger  = "/swagger.yaml"

	// user
	userRegister = "/user/register"
	userLogin    = "/user/login"
	userProfile  = "/user/profile"

	// crypto
	cryptoQueries       = "/crypto/queries"
	cryptoPrice         = "/crypto/price"
	cryptoSupCurrencies = "/crypto/supcurrencies"
)

func ProvideRouter(
	userCtrl *controllers.User,
	cryptoCtrl *controllers.Crypto,
	mredis *mredis.Wrap,
	jwt gin.HandlerFunc,
) *Engine {
	return &Engine{
		gin:        gin.Default(),
		mredis:     mredis,
		jwt:        jwt,
		userCtrl:   userCtrl,
		cryptoCtrl: cryptoCtrl,
	}
}

func (e *Engine) Init() {
	jwt := e.gin.Group("", e.jwt)

	// GET
	e.gin.GET(root, home)
	e.gin.GET(readmeEN, readmeLangEN)
	e.gin.GET(readmeZH, readmeLangZH)
	e.gin.GET(docs, apiDocs)
	e.gin.GET(swagger, swaggerYaml)

	jwt.GET(userProfile, e.userCtrl.Profile)
	jwt.GET(cryptoQueries, e.cryptoCtrl.SearchCoin)
	jwt.GET(cryptoSupCurrencies, e.cryptoCtrl.SupportCurrencies)
	jwt.GET(cryptoPrice, e.cryptoCtrl.GetCryptoPriceByIDs)

	// POST
	e.gin.POST(userRegister, e.userCtrl.Register)
	e.gin.POST(userLogin, e.mredis.RateLimitMiddleware(5, time.Minute), e.userCtrl.Login)

	if port := os.Getenv("PORT"); port == "" {
		log.Fatalf("failed to run server, invalid PORT")
	} else {
		if err := e.gin.Run(fmt.Sprintf(":%s", port)); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
	}
}

func home(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, readmeEN)
}

func readmeLangEN(ctx *gin.Context) {
	readmeMD("README.md", ctx)
}

func readmeLangZH(ctx *gin.Context) {
	readmeMD("README.zh.md", ctx)
}

func apiDocs(ctx *gin.Context) {
	tmplPath := filepath.Join("templates", "swagger.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to read template file")
		return
	}

	var out bytes.Buffer
	err = tmpl.Execute(&out, map[string]interface{}{}) // 可塞參數
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to render template")
		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", out.Bytes())
}

func swaggerYaml(ctx *gin.Context) {
	ctx.File("docs/swagger.yaml")
}

func readmeMD(path string, ctx *gin.Context) {
	mdPath := filepath.Join(path)
	mdBytes, err := os.ReadFile(mdPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to read markdown file")
		return
	}

	var mdBuf bytes.Buffer
	if err := goldmark.Convert(mdBytes, &mdBuf); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to convert markdown")
		return
	}

	tmplPath := filepath.Join("templates", "markdown.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to read template file")
		return
	}

	var out bytes.Buffer
	err = tmpl.Execute(&out, map[string]interface{}{
		"Content": template.HTML(mdBuf.String()),
	})
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to render template")
		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", out.Bytes())
}
