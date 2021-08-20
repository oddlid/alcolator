package srv

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/oddlid/alcolator/drink"
	"github.com/rs/zerolog/log"
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
		log.Warn().
			Str("remote_addr", r.RemoteAddr).
			Msg("Should have form data, but it seems to be missing")
	}

	fd.DrinkName = r.FormValue("dname")
	var err error
	fd.DrinkVolML, err = strconv.Atoi(r.FormValue("dvolml"))
	if err != nil {
		log.Error().
			Str("remote_addr", r.RemoteAddr).
			Str("dvolml", r.FormValue("dvolml")).
			Msg("Failed to convert input")
	}
	dpct := strings.Replace(r.FormValue("dpct"), ",", ".", -1) // replace comma with dot in input
	fd.DrinkPCT, err = strconv.ParseFloat(dpct, 64)
	if err != nil {
		log.Error().
			Str("remote_addr", r.RemoteAddr).
			Str("dpct", r.FormValue("dpct")).
			Msg("Failed to convert input")
	}
	fd.DrinkPrice, err = strconv.ParseFloat(r.FormValue("dprice"), 64)
	if err != nil {
		log.Error().
			Str("remote_addr", r.RemoteAddr).
			Str("dprice", r.FormValue("dprice")).
			Msg("Failed to convert input")
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
