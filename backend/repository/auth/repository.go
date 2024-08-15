package auth

import (
	"errors"
	"sync"

	"github.com/drizzleent/wallet-tracker/backend/internal/model"
	"github.com/drizzleent/wallet-tracker/backend/repository/converter"
	datamodel "github.com/drizzleent/wallet-tracker/backend/repository/data_model"
	"github.com/google/uuid"
)

var id uuid.UUID

type repo struct {
	m     sync.RWMutex
	once  sync.Once
	users map[string]datamodel.User
}

func NewAuthRepository() *repo {
	r := repo{
		users: make(map[string]datamodel.User),
	}
	return &r
}

func (r *repo) Register(p *model.RegisterPayload) (*model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()
	nonce, err := r.Nonce()
	if err != nil {
		return nil, errors.New("failed to create user nonce " + err.Error())
	}

	u := converter.FromRegisterPayloadToDataUser(p)
	u.Nonce = nonce

	if _, exist := r.users[u.Nonce]; exist {
		return nil, errors.New("user already exist")
	}

	r.users[u.Address] = *u

	return converter.FromDataUserToModelUser(u), nil
}

func (r *repo) UserNonce(address string) (*model.User, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	u, exist := r.users[address]
	if !exist {
		return nil, errors.New("user not exist")
	}

	return converter.FromDataUserToModelUser(&u), nil
}

func (r *repo) Update(user *model.User) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.users[user.Address] = *converter.FromUserToDataUser(user)
	return nil
}

func (r *repo) Get(address string) (*model.User, error) {
	r.m.Lock()
	defer r.m.Unlock()

	u, exist := r.users[address]
	if !exist {
		return nil, errors.New("user not exist")
	}

	return converter.FromDataUserToModelUser(&u), nil
}

func (r *repo) Nonce() (string, error) {
	var err error
	r.once.Do(func() {
		id, err = uuid.NewRandom()
	})

	if err != nil {
		return "", err
	}

	return id.String(), nil
}
