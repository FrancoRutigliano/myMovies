package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
)

type Store struct {
	Users *[]models.User
}

func NewStore(fileName string) (*Store, error) {
	// Verificar si el archivo existe
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		// El archivo no existe, inicializar con datos predeterminados
		err := initializeStoreWithDefaults(fileName)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err // Manejar otros errores al verificar la existencia del archivo
	}

	// Abrir el archivo JSON
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decodificar el archivo JSON
	var users []models.User
	if err := json.NewDecoder(file).Decode(&users); err != nil {
		return nil, err
	}

	return &Store{Users: &users}, nil
}

func initializeStoreWithDefaults(fileName string) error {
	// Crear una estructura de datos inicial con datos predeterminados
	defaultUsers := []models.User{
		{Name: "Usuario 1", Email: "usuario1@example.com", Password: "contraseña1"},
		{Name: "Usuario 2", Email: "usuario2@example.com", Password: "contraseña2"},
		// Agregar más usuarios predeterminados si es necesario
	}

	// Abrir o crear el archivo JSON
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Codificar y escribir los datos predeterminados en el archivo JSON
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(defaultUsers); err != nil {
		return err
	}

	return nil
}

func (s *Store) FindByEmail(email string) (*models.User, error) {
	for _, user := range *s.Users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (s *Store) EmailExist(email string) error {
	for _, user := range *s.Users {
		if user.Email == email {
			return fmt.Errorf("email user already exists")
		}
	}
	return nil
}

func (s *Store) CreateUser(user *models.User) error {

	idUser := len(*s.Users) + 1
	user.ID = int64(idUser)

	*s.Users = append(*s.Users, *user)

	// guardar los cambios en el json
	return helpers.StoreJson("./data/user.json", *s.Users)
}

func (s *Store) GetAll() []models.User {
	return *s.Users
}

func (s *Store) UpdateUserPassword(user *models.User) error {
	for i, u := range *s.Users {
		if u.ID == user.ID {
			(*s.Users)[i].Password = user.Password
			return helpers.StoreJson("./data/user.json", *s.Users)
		}
	}
	return errors.New("user not found")
}
