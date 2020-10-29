package productCategoryController

import (
	"encoding/json"
	"github.com/Den1ske/GoMarket/src/app/entity"
	"github.com/Den1ske/GoMarket/src/app/store/productCategoryRepository"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type ProductCategoryControllerInterface interface {
	List(http.ResponseWriter, *http.Request)
	GetByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
}

type productCategoryController struct {
	productCategoryStore *productCategoryRepository.Store
}

func (pc *productCategoryController) List(w http.ResponseWriter, r *http.Request) {

	productCategorys, err := pc.productCategoryStore.Category().List()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	if productCategorys != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(productCategorys)
	}
}

func (pc *productCategoryController) GetByID(w http.ResponseWriter, r *http.Request) {

	var id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err !=nil {
		log.Print(err.Error())
	}

	productCategory, err := pc.productCategoryStore.Category().GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	if productCategory != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(productCategory)
	}
}

func (pc *productCategoryController) Create(w http.ResponseWriter, r *http.Request) {

	var err error
	var productCategory *entity.ProductCategory

	if err != nil {
		log.Fatal(err.Error())
	}

	productCategory = &entity.ProductCategory{
		Title:       r.FormValue("title"),
	}

	productCategory, err = pc.productCategoryStore.Category().Create(productCategory)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	if productCategory != nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(productCategory)
	}
}
