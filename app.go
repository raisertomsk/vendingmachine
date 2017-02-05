package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	Products map[string]*Product
	Config   *Config
	Handler  *Handler
}

func NewApp(config *Config) *App {
	app := new(App)
	app.Config = config
	app.Products = map[string]*Product{}
	app.Handler = NewHandler(app)
	return app
}

func (a *App) Run() {
	log.Fatalln(http.ListenAndServe(a.Config.Listen, a.getRouter()))
}

func (app *App) getRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/list", app.Handler.List)
	router.GET("/get/:name", app.Handler.Get)
	router.PUT("/add/:name", app.Handler.Add)
	router.POST("/update/:name", app.Handler.Update)
	router.DELETE("/delete/:name", app.Handler.Delete)
	return router
}

func (app *App) List() map[string]*Product {
	products := map[string]*Product{}
	for index, product := range app.Products {
		if product.Quantity > 0 {
			products[index] = product
		}
	}
	return products
}

func (app *App) Get(name string, count int, cash float64) (*Product, *Error) {
	return nil, nil
}

func (app *App) Add(name string, data []byte) *Error {
	product := new(Product)
	err := json.Unmarshal(data, product)
	if err != nil {
		log.Println("Error parse product: " + err.Error())
		return NewError(E_PARSE_PRODUCT_JSON, 0)
	}
	if app.Products[name] != nil {
		return NewError(E_PRODUCT_ALREADY_EXISTS, 0)
	}
	if len(name) == 0 {
		return NewError(E_PRODUCT_NAME_LENGTH, 0)
	}
	if len(product.Title) == 0 {
		return NewError(E_PRODUCT_TITLE_LENGTH, 0)
	}
	app.Products[name] = product
	return NewError(E_NO_ERROR, 0)
}

func (app *App) Update(name string, data []byte) *Error {
	return NewError(E_NO_ERROR, 0)
}

func (app *App) Delete(name string) *Error {
	if app.Products[name] == nil {
		return NewError(E_PRODUCT_NOT_EXIST, 0)
	}
	delete(app.Products, name)
	return NewError(E_NO_ERROR, 0)
}
