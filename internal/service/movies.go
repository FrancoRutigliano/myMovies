package service

import (
	"encoding/json"
	"os"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
)

var defaultMovies = []models.Movie{
	{ID: 1, Title: "Casa blanca", Year: 1999, Runtime: 102, Genres: []string{"drama", "comedia", "horror"}},
	{ID: 2, Title: "Casa negra", Year: 1998, Runtime: 102, Genres: []string{"drama", "comedia", "horror"}},
	{ID: 3, Title: "Casa verde", Year: 1997, Runtime: 102, Genres: []string{"drama", "comedia", "horror"}},
}

type MovieStore struct {
	Movies *[]models.Movie
}

func NewMovieStore(filename string) (*MovieStore, error) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// El archivo no existe, inicializar con datos predeterminados
		err := helpers.InitializeStoreWithDefaults(filename, defaultMovies)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		// Manejar otros errores al verificar la existencia del archivo
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var movies []models.Movie
	if err := json.NewDecoder(file).Decode(&movies); err != nil {
		return nil, err
	}
	return &MovieStore{Movies: &movies}, nil
}

func (s *MovieStore) CreateMovie(movie *models.Movie) error {
	idMovie := len(*s.Movies) + 1
	movie.ID = int64(idMovie)

	*s.Movies = append(*s.Movies, *movie)
	return helpers.StoreJson("./data/movies.json", *s.Movies)
}
