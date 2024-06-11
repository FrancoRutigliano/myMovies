package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
)

var defaultUsers = []models.User{
	{Name: "Usuario 1", Email: "usuario1@example.com", Password: "contraseña1"},
	{Name: "Usuario 2", Email: "usuario2@example.com", Password: "contraseña2"},
}

type UserStore struct {
	Users *[]models.User
}

func NewUserStore(filename string) (*UserStore, error) {
	// Verificar si el archivo existe
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// El archivo no existe, inicializar con datos predeterminados
		err := helpers.InitializeStoreWithDefaults(filename, defaultUsers)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err // Manejar otros errores al verificar la existencia del archivo
	}

	// Abrir el archivo JSON
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decodificar el archivo JSON
	var users []models.User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		return nil, err
	}

	return &UserStore{Users: &users}, nil
}

func (s *UserStore) FindById(id int64) (*models.User, error) {
	for _, user := range *s.Users {
		if id == user.ID {
			userProfile := &models.User{
				Name:      user.Name,
				Review:    user.Review,
				CreatedAt: user.CreatedAt,
			}
			return userProfile, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserStore) FindByEmail(email string) (*models.User, error) {
	user, ok := s.EmailExist(email)

	if ok {
		userProfile := &models.User{
			Name:   user.Name,
			Review: user.Review,
		}
		return userProfile, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (s *UserStore) EmailExist(email string) (*models.User, bool) {
	for _, user := range *s.Users {
		if user.Email == email {
			return &user, true
		}
	}
	return nil, false
}

func (s *UserStore) CreateUser(user *models.User) error {

	idUser := len(*s.Users) + 1
	user.ID = int64(idUser)

	*s.Users = append(*s.Users, *user)

	// guardar los cambios en el json
	return helpers.StoreJson("./data/user.json", *s.Users)
}

func (s *UserStore) GetAll() []models.User {
	return *s.Users
}

func (s *UserStore) UpdateUserPassword(user *models.User) error {
	for i, u := range *s.Users {
		if u.ID == user.ID {
			(*s.Users)[i].Password = user.Password
			return helpers.StoreJson("./data/user.json", *s.Users)
		}
	}
	return errors.New("user not found")
}
