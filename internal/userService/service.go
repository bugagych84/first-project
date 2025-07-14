package userService

// Интерфейс сервиса задач
type UserService interface {
	CreateUser(user User) ([]User, error)
	GetAllUsers() ([]User, error)
	GetUserById(userId string) (User, error)
	UpdateUser(userId string, newUser User) ([]User, error)
	DeleteUserById(userId string) ([]User, error)
}

// Структура сервиса задач
type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{repo: r}
}

func (s userService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s userService) CreateUser(user User) ([]User, error) {
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return s.GetAllUsers()
}

func (s userService) GetUserById(userId string) (User, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s userService) UpdateUser(userId string, newUser User) ([]User, error) {
	user, err := s.repo.GetUserById(userId)
	if err != nil {
		return []User{}, err
	}

	user.Email = newUser.Email
	user.Password = newUser.Password

	if err := s.repo.UpdateUser(user); err != nil {
		return []User{}, err
	}

	users, err := s.GetAllUsers()
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (s userService) DeleteUserById(userId string) ([]User, error) {
	_, err := s.repo.GetUserById(userId)
	if err != nil {
		return []User{}, err
	}

	err = s.repo.DeleteUserById(userId)
	if err != nil {
		return []User{}, err
	}

	users, err := s.GetAllUsers()
	if err != nil {
		return []User{}, err
	}

	return users, nil
}
