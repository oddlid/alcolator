package srv

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/oddlid/alcolator/drink"
	log "github.com/sirupsen/logrus"
)

type AlcolatorServer struct {
	htmlTemplate *template.Template
}

type FormData struct {
	HasData               bool
	DrinkName             string
	DrinkVolML            int
	DrinkPCT              float64
	DrinkPrice            float64
	ResultWaterML         int
	ResultAlcML           int
	ResultPricePerAlcML   float64
	ResultPricePerDrinkML float64
}

func NewServer(tmpl *template.Template) *AlcolatorServer {
	return &AlcolatorServer{htmlTemplate: tmpl}
}

func (as *AlcolatorServer) CalcHandler(w http.ResponseWriter, r *http.Request) {
	hasData := r.FormValue("dsub") != ""

	if !hasData {
		as.htmlTemplate.Execute(w, nil)
		return
	}

	fd := FormData{}
	fd.HasData = hasData
	if !fd.HasData {
		log.Warnf("Should have form data, but it seems to be missing... From: %q", r.RemoteAddr)
	}

	fd.DrinkName = r.FormValue("dname")
	var err error
	fd.DrinkVolML, err = strconv.Atoi(r.FormValue("dvolml"))
	if err != nil {
		log.Errorf("Failed to convert dvolml. Input: %q from %q", r.FormValue("dvolml"), r.RemoteAddr)
	}
	dpct := strings.Replace(r.FormValue("dpct"), ",", ".", -1) // replace comma with dot in input
	fd.DrinkPCT, err = strconv.ParseFloat(dpct, 64)
	if err != nil {
		log.Errorf("Failed to convert dpct. Input: %q from %q", r.FormValue("dpct"), r.RemoteAddr)
	}
	fd.DrinkPrice, err = strconv.ParseFloat(r.FormValue("dprice"), 64)
	if err != nil {
		log.Errorf("Failed to convert dprice. Input: %q from %q", r.FormValue("dprice"), r.RemoteAddr)
	}

	d := &drink.Drink{
		Name:     fd.DrinkName,
		VolumeML: fd.DrinkVolML,
		AlcPCT:   fd.DrinkPCT,
		Price:    fd.DrinkPrice,
	}

	fd.ResultWaterML = d.WaterML()
	fd.ResultAlcML = d.AlcML()
	fd.ResultPricePerAlcML = d.PricePerAlcML()
	fd.ResultPricePerDrinkML = d.PricePerDrinkML()

	as.htmlTemplate.Execute(w, &fd)
}
