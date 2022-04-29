package repository

type UserRepository struct {
}

func (u *UserRepository) GetAllUsers() []map[string]interface{} {
	var users = []map[string]interface{}{
		{
			"name":       "Yogi Pristiawan",
			"gender":     "L",
			"birth_date": "2000-07-24",
			"height":     165,
			"wieght":     40,
		},
		{
			"name":       "Ica Ramdani",
			"gender":     "P",
			"birth_date": "2001-12-14",
			"height":     160,
			"wieght":     40,
		},
		{
			"name":       "Dimas Maulana",
			"gender":     "L",
			"birth_date": "2004-12-12",
			"height":     165,
			"weight":     60,
		},
		{
			"name":       "Alviana Linda",
			"gender":     "P",
			"birth_date": "2002-12-12",
			"height":     165,
			"weight":     50,
		},
	}

	return users
}
