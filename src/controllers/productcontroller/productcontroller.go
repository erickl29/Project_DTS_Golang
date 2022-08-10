package productcontroller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"training/src/entities"
	"training/src/models"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var productModel models.ProductModel
	products, _ := productModel.FindAll()
	data := map[string]interface{}{
		"products": products,
	}

	tmp, _ := template.ParseFiles("views/product/index.html")
	tmp.Execute(response, data)

}

func Add(response http.ResponseWriter, request *http.Request) {

	tmp, _ := template.ParseFiles("views/product/add.html")
	tmp.Execute(response, nil)

}

func ProcessAdd(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	var pegawai entities.Product
	pegawai.NIP, _ = strconv.ParseInt(request.Form.Get("nip"), 10, 64)
	pegawai.Nama = request.Form.Get("nama")
	pegawai.Alamat = request.Form.Get("alamat")
	pegawai.Nohp, _ = strconv.ParseInt(request.Form.Get("nohp"), 10, 64)

	var productModel models.ProductModel
	productModel.Create(&pegawai)
	http.Redirect(response, request, "/product", http.StatusSeeOther)

}

func Delete(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	NIP, _ := strconv.ParseInt(query.Get("nip"), 10, 64)
	fmt.Println("NIP:", NIP)
	var productModel models.ProductModel
	productModel.Delete(NIP)
	http.Redirect(response, request, "/product", http.StatusSeeOther)

}
