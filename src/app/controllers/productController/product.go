package productController

import (
	"encoding/json"
	"github.com/Den1ske/GoMarket/src/app/entity"
	"github.com/Den1ske/GoMarket/src/app/store/productRepository"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
)

type ProductControllerInterface interface {
	List(http.ResponseWriter, *http.Request)
	GetByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
}

type productController struct {
	productStore *productRepository.Store
}

func (pc *productController) List(w http.ResponseWriter, r *http.Request) {

	products, err := pc.productStore.Product().List()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	if products != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func (pc *productController) GetByID(w http.ResponseWriter, r *http.Request) {

	var id, err = strconv.Atoi(chi.URLParam(r, "id"))
	if err !=nil {
		log.Print(err.Error())
	}

	product, err := pc.productStore.Product().GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	if product != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func (pc *productController) Create(w http.ResponseWriter, r *http.Request) {

	var article = r.FormValue("article")
	var categoryID *int64
	var price int
	var err error
	var product *entity.Product

	if(r.FormValue("category_id") == "") {
		categoryID = nil
	} else {
		var cID int64
		cID, err = strconv.ParseInt(r.FormValue("category_id"), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		categoryID = &cID
	}


	if err != nil {
		log.Fatal(err.Error())
	}
	price, err = strconv.Atoi(r.FormValue("price"))
	if err != nil {
		log.Fatal(err.Error())
	}

	product, _ = pc.productStore.Product().GetByArticle(article)

	if product != nil {
		w.WriteHeader(http.StatusConflict)
		var httpError = map[string]interface{}{
			"error_message": "Ошибка, товар с таким article уже существует",
		}
		json.NewEncoder(w).Encode(httpError)
		return
	}

	product = &entity.Product{
		Title:       r.FormValue("title"),
		Article:     &article,
		Price:       price,
		Category_ID: categoryID,
	}

	product, err = pc.productStore.Product().Create(product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}

	if product != nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}
