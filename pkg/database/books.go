package database

import (
	"github.com/Aitugan/Techt/entity"
	"github.com/google/uuid"
)

var books []*entity.Book = []*entity.Book{
	{
		Id:              uuid.MustParse("da011294-51c3-4c01-8624-c2f1ac1671b3"),
		Title:           "Atomic Habits",
		Author:          "James Clear",
		PublicationDate: "01.10.2018",
	},
	{
		Id:              uuid.MustParse("96da145a-456e-46d6-9b4e-3a8fd5960287"),
		Title:           "The Psychology of Money",
		Author:          "Morgan Housel",
		PublicationDate: "01.09.2020",
	},
	{
		Id:              uuid.MustParse("3c28de98-e09a-4f3e-884b-624694b13c96"),
		Title:           "Ikigai: The Japanese secret to a long and happy life",
		Author:          "Francesc Miralles",
		PublicationDate: "27.09.2017",
	},
	{
		Id:              uuid.MustParse("e9da499e-1f02-4b43-aee5-06836bf72b66"),
		Title:           "Educart Abhyaas 1-20 Final Bundle of NEET Mock Test Papers",
		Author:          "Seep Pahuja",
		PublicationDate: "19.01.2023",
	},
	{
		Id:              uuid.MustParse("d5bded79-fe13-4370-9a20-1ca47b53cba4"),
		Title:           "Energize Your Mind: Learn the Art of Mastering Your Thoughts",
		Author:          "Gaur Gopal Das",
		PublicationDate: "01.01.2023",
	},
	{
		Id:              uuid.MustParse("b5bd3fb7-b497-4d3b-a9ea-d1f997c33a88"),
		Title:           "Grandma's Bag of Stories",
		Author:          "Sudha Murty",
		PublicationDate: "01.01.2015",
	},
}

