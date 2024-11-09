package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"workWithCache/cache"
	"workWithCache/db"
	converting "workWithCache/server/elements/converting"

	"github.com/go-chi/chi"
)

func responseValue(parameter string, _ http.ResponseWriter, r *http.Request) converting.GetInfoRequestData {
	return converting.GetInfoRequestData(r.URL.Query().Get(parameter))
}

type Server struct {
	db    *db.Database
	cache *cache.Cache
}

func New(db *db.Database, cache *cache.Cache) *Server {
	return &Server{
		db:    db,
		cache: cache,
	}
}

func Run() {
	router := chi.NewRouter()

	s := New(db.New(), cache.New())

	router.Get("/info", s.InfoHandler)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func (s *Server) InfoHandler(w http.ResponseWriter, r *http.Request) {
	parameter := "numbers"
	parameterValue := responseValue(parameter, w, r)
	if !checkParameterValue(parameterValue, parameter, w) {
		return
	}

	inputNumbers, err := parameterValue.ToIntSlice(parameter)
	if !checkInputNumbers(err, w) {
		return
	}

	factor := 4

	cache := s.cache
	userResponse, err := s.db.FindNumbers(w, inputNumbers, factor, cache)
	if !checkDatabase(err, w) {
		return
	}

	fmt.Fprintf(w, "User Response: %s\n", strings.Join(intSliceToStringSlice(userResponse), ","))

	cache.Data.Output()

	// пытаемся найти в кэше, потому что это быстрее
	// cache.Find(inputNumbers)
	// идём в бд и "ищем" эти числа, если надо
	// собираем результат
	// кладём в кэш результаты
	// отдаёт пользователю
}

func checkParameterValue(parameterValue converting.GetInfoRequestData, parameter string, w http.ResponseWriter) bool {
	if parameterValue == "" {
		http.Error(w, fmt.Sprintf("This URL doesn't contain requested parameter '%s'", parameter), http.StatusBadRequest)
		return false
	}
	return true
}

func checkInputNumbers(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("%v", err)
		fmt.Fprintln(w, err)
		return false
	}
	return true
}

func checkDatabase(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Printf("%v", err)
		fmt.Fprintln(w, err)
		return false
	}
	return true
}

func intSliceToStringSlice(ints []int) []string {
	strs := make([]string, len(ints)) // Создаем срез строк нужной длины
	for i, v := range ints {
		strs[i] = strconv.Itoa(v) // Преобразуем int в string
	}
	return strs
}
