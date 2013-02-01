package main

import (
	"github.com/nvcnvn/gorms"
	"github.com/openvn/gocms"
	"github.com/openvn/toys/locale"
	"github.com/openvn/toys/view"
	"labix.org/v2/mgo"
	"net/http"
)

func main() {
	//database session
	dbsess, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer dbsess.Close()

	//multi language support
	cmslang := locale.NewLang("language/cmslang")
	cmslang.Parse("vi")
	cmslang.Parse("en")

	rmslang := locale.NewLang("language/rmslang")
	rmslang.Parse("vi")
	rmslang.Parse("en")

	//template for cms
	cmstmpl := view.NewView("template/cmstmpl")
	cmstmpl.Resource = "//localhost:8080/statics"
	cmstmpl.SetLang(cmslang)
	cmstmpl.Parse("default")

	//template for rms
	rmstmpl := view.NewView("template/rmstmpl")
	rmstmpl.Resource = "//localhost:8080/statics"
	rmstmpl.SetLang(rmslang)
	rmstmpl.Parse("default")

	//routing
	http.Handle("/", gocms.NewHandler(gocms.Router, dbsess, cmstmpl))
	http.Handle("/data/", gorms.NewHandler(gorms.Data, dbsess, rmstmpl))
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
	http.ListenAndServe("localhost:8080", nil)
}
