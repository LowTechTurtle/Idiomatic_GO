package stub

type User struct{}
type Pet struct {
	Name string
}
type Person struct{}

type Entities interface {
	GetUser(id string) (User, error)
	GetPets(userID string) ([]Pet, error)
	GetChildren(userID string) ([]Person, error)
	GetFriends(userID string) ([]Person, error)
	SaveUser(user User) error
}

// one way to test large interface is to embed it in a struct
// we must implement the method that will be called 
// or its a runtime error
type Logic struct {
	Entities Entities
}

func (l Logic) GetPetNames(userId string) ([]string, error) {
	pets, err := l.Entities.GetPets(userId)
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(pets))
	for _, p := range pets {
		out = append(out, p.Name)
	}
	return out, nil
}