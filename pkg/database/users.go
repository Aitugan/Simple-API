package database

import (
	"github.com/Aitugan/Techt/entity"
	"github.com/google/uuid"
)

var users []*entity.User = []*entity.User{
	{
		Id:       uuid.MustParse("4191aaba-14bd-4e99-a1ac-36006ea4da5b"),
		Email:    "James@mail.com",
		Password: "password",
		Name:     "James Clear",
	},
	{
		Id:       uuid.MustParse("fc44ea64-afc4-480d-a790-58c98d4dbb8d"),
		Email:    "Morgan@mail.com",
		Password: "password",
		Name:     "Morgan Housel",
	},
	{
		Id:       uuid.MustParse("ac324845-30cb-4664-bcbc-164a0b528b93"),
		Email:    "Francesc@mail.com",
		Password: "password",
		Name:     "Francesc Miralles",
	},
	{
		Id:       uuid.MustParse("01d89a8b-8f07-420a-a0d5-f42be831665f"),
		Email:    "Seep@mail.com",
		Password: "password",
		Name:     "Seep Pahuja",
	},
	{
		Id:       uuid.MustParse("9534815b-d15d-4ed3-b2d5-fe2fca9665d6"),
		Email:    "Gaur@mail.com",
		Password: "password",
		Name:     "Gaur Gopal Das",
	},
	{
		Id:       uuid.MustParse("3432859b-3e12-4b9c-82bd-34a671905d41"),
		Email:    "Sudha@mail.com",
		Password: "password",
		Name:     "Sudha Murty",
	},
	{
		Id:       uuid.MustParse("966e2778-0897-4de1-bb42-682b23773661"),
		Email:    "John@mail.com",
		Password: "password",
		Name:     "John Doe",
	},
}

