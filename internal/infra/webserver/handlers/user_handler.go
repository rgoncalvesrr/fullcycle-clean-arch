package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/dto"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/entity"
	"github.com/rgoncalvesrr/fullcycle-clean-arch/internal/infra/database"
)

type UserHandler struct {
	UserGateway  database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(db database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		UserGateway:  db,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var userDto dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserGateway.FindByEmail(userDto.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(userDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&accessToken)
	// token := jwtauth.New("HS256", []byte("secret"), nil)

}

// Create user godoc
//
//	@Summay			Create User
//	@Description	Create User
//	@Tags			users
//	@Accept			json
//	@Produce		json
//
//	@Param			request	body	dto.CreateUserInput	true	"user request"
//	@Success		201
//	@Failure		500	{object}	dto.Error
//	@Router			/users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.UserGateway.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	o := &dto.CreateUserOutput{
		ID:    p.ID.String(),
		Name:  p.Name,
		Email: p.EMail,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&o)
}
