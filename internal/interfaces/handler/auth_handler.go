package handler

import (
	"auth_service/internal/domain"
	"auth_service/internal/usecase"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthUseCase *usecase.AuthUseCase
}

func NewAuthHandler(useCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		AuthUseCase: useCase,
	}
}

// Register godoc
// @Summary Registra un nuevo usuario
// @Description Registra un nuevo usuario con un nombre de usuario y contraseña
// @Tags auth
// @Accept json
// @Produce json
// @Param user body domain.User true "Información del usuario"
// @Success 201 {string} string "User registered"
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Internal server error"
// @Router /register [post]
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.AuthUseCase.Register(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("User registered")
}

// Login godoc
// @Summary Inicia sesión de un usuario
// @Description Inicia sesión con un nombre de usuario y contraseña
// @Tags auth
// @Accept json
// @Produce json
// @Param user body domain.User true "Información del usuario"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {string} string "Invalid request"
// @Failure 401 {string} string "Invalid credentials"
// @Failure 500 {string} string "Internal server error"
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.AuthUseCase.Login(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token, "message": "Login successful"})
}
