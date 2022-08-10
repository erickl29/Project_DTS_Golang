package main

import (
	"net/http"
	"training/src/controllers/productcontroller"
)

func main() {

	http.HandleFunc("/", productcontroller.Index)
	http.HandleFunc("/product", productcontroller.Index)
	http.HandleFunc("/product/index", productcontroller.Index)
	http.HandleFunc("/product/add", productcontroller.Add)
	http.HandleFunc("/product/processadd", productcontroller.ProcessAdd)
	http.HandleFunc("/product/delete", productcontroller.Delete)

	fileServer := http.FileServer(http.Dir("css"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))

	http.ListenAndServe(":3000", nil)

}
