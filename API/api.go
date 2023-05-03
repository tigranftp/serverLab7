package API

import (
	"db_lab7/config"
	"db_lab7/db"
	"db_lab7/types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type API struct {
	config *config.Config
	router *mux.Router
	store  *db.Store
}

func NewAPI() (*API, error) {
	res := new(API)
	var err error
	res.config, err = config.GetConfig()
	if err != nil {
		return nil, err
	}
	res.router = mux.NewRouter()
	return res, nil
}

func (a *API) Start() error {
	a.configureRouter()
	a.configureDB()
	fmt.Println(a.store.Open())
	return http.ListenAndServe(a.config.Port, a.router)
}

func (a *API) Stop() {
	fmt.Println("Stopping API...")
	a.store.Close()
	fmt.Println("API stopped...")
}

func (a *API) configureRouter() {
	a.router.HandleFunc("/kek", a.handleKEK())
	a.router.HandleFunc("/add_country", a.handleAddCountry())
	a.router.HandleFunc("/delete_university", a.handleDeleteUniversity())
	a.router.HandleFunc("/add_university", a.handleAddUniversity())
	a.router.HandleFunc("/delete_ranking_criteria", a.handleDeleteRankingCriteria())
	a.router.HandleFunc("/change_university_year_staff_ratio", a.handleChangeUniversityYearStaffRatio())
	a.router.HandleFunc("/add_university_ranking_year", a.handleAddUniversityRankingYear())
}

func (a *API) configureDB() {
	a.store = db.New(a.config)
}

func (a *API) handleAddUniversityRankingYear() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "can't read body", http.StatusBadRequest)
			return
		}
		err = request.Body.Close()
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		var aury types.AddUniversityRankingYear
		err = json.Unmarshal(body, &aury)
		if err != nil {
			http.Error(writer, "error during unmarshal", http.StatusBadRequest)
			return
		}
		_, err = a.store.Exec(db.AddUniversityRankingYear, aury.UniversityName, aury.CriteriaName, aury.Year, aury.Score)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func (a *API) handleChangeUniversityYearStaffRatio() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "can't read body", http.StatusBadRequest)
			return
		}
		err = request.Body.Close()
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		var cssr types.ChangeStudentStaffRatio
		err = json.Unmarshal(body, &cssr)
		if err != nil {
			http.Error(writer, "error during unmarshal", http.StatusBadRequest)
			return
		}
		_, err = a.store.Exec(db.ChangeUniversityYearStaffRatio, cssr.NewStaffRatio, cssr.UniversityName, cssr.Year)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func (a *API) handleAddCountry() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "can't read body", http.StatusBadRequest)
			return
		}
		err = request.Body.Close()
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		var cnt types.Country
		err = json.Unmarshal(body, &cnt)
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		if cnt.CountryName == "" {
			http.Error(writer, "can't add country empty with empty countryName", http.StatusInternalServerError)
			return
		}
		_, err = a.store.Exec(db.AddCountryQuery, cnt.CountryName)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func (a *API) handleDeleteUniversity() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "can't read body", http.StatusBadRequest)
			return
		}
		err = request.Body.Close()
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		var univ types.University
		err = json.Unmarshal(body, &univ)
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		if univ.UniversityName == "" {
			http.Error(writer, "can't add country empty with empty countryName", http.StatusInternalServerError)
			return
		}
		_, err = a.store.Exec(db.DeleteUniversityQuery, univ.UniversityName, univ.UniversityName, univ.UniversityName)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func (a *API) handleDeleteRankingCriteria() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "can't read body", http.StatusBadRequest)
			return
		}
		err = request.Body.Close()
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		var rc types.RankingCriteria
		err = json.Unmarshal(body, &rc)
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		if rc.CriteriaName == "" {
			http.Error(writer, "can't add country empty with empty countryName", http.StatusInternalServerError)
			return
		}
		_, err = a.store.Exec(db.DeleteRankingCriteriaQuery, rc.SystemID, rc.CriteriaName, rc.SystemID, rc.CriteriaName)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func (a *API) handleAddUniversity() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "can't read body", http.StatusBadRequest)
			return
		}
		err = request.Body.Close()
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		var univ types.University
		err = json.Unmarshal(body, &univ)
		if err != nil {
			http.Error(writer, "can't close body", http.StatusInternalServerError)
			return
		}
		if univ.UniversityName == "" {
			http.Error(writer, "can't add country empty with empty countryName", http.StatusInternalServerError)
			return
		}
		_, err = a.store.Exec(db.AddUniversityQuery, univ.CountryName, univ.UniversityName)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
	}
}

func (a *API) handleKEK() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		a.GetAllCountries()
	}
}

func (a *API) GetAllCountries() error {
	rows, err := a.store.Query(db.SelectAllCountries)
	if err != nil {
		return err
	}
	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(id, name)
	}
	return nil
}
