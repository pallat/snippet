package payment

type FavoriteFunc func(id string) (string, error)

func Favorite(db dber) FavoriteFunc {
	return func(id string) (string, error) {
		return "KTB", nil
	}
}
