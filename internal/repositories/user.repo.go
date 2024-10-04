package repositories

import (
	"jeevva/cms-be-service/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser{
	return &RepoUser{db}
}

func(r *RepoUser) CreateUser(data *models.User)(string, error){
	q := `INSERT INTO public.users(
				first_name,
				last_name,
				email, 
				"password",
				"role",
				image,
				phone_number)
				VALUES(:first_name, :last_name, :email, :password, :role, :image, :phone_number)`
	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}
	return "1 user created", nil
}