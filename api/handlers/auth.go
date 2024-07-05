package handlers

import (
	"auth-service/api/token"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/swaggo/swag"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, username, and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.RegisterReqSwag true "User registration request"
// @Success 201 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /register [post]
func (h *HTTPHandler) Register(c *gin.Context) {
	var req models.RegisterReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	exists, err := h.US.EmailExists(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "err": err.Error()})
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered", "email": req.Email})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "err": err.Error()})
		return
	}
	req.Password = string(hashedPassword)

	err = h.US.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error", "err": err.Error()})
		return
	}
	req.Role = "user"
	tokens := token.GenerateJWTToken(req.ID, req.Email, req.Username, req.Role)

	c.JSON(http.StatusCreated, tokens)
}

// Login godoc
// @Summary Login a user
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginReq true "User login credentials"
// @Success 200 {object} token.Tokens "JWT tokens"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 401 {object} string "Invalid email or password"
// @Router /login [post]
func (h *HTTPHandler) Login(c *gin.Context) {
	req := models.LoginReq{}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	user, err := h.US.GetProfile(&models.GetProfileReq{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Invalid email or password": err.Error()})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Invalid email or password...": err.Error()})
		return
	}

	tokens := token.GenerateJWTToken(user.ID, user.Email, user.Username, user.Role)

	c.JSON(http.StatusOK, tokens)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.GetProfileResp
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "User not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /profile [get]
func (h *HTTPHandler) Profile(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	email := claims.(jwt.MapClaims)["email"].(string)
	user, err := h.US.GetProfile(&models.GetProfileReq{Email: email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Server error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *HTTPHandler) GetByID(c *gin.Context) {
	id := &models.GetProfileByIdReq{ID: c.Param("id")}
	user, err := h.US.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Couldn't get the user": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// ChangeUserRole godoc
// @Summary Change a user's role
// @Description Changes the role of a user or admin. Only admins are allowed to use this function.
// @Tags admin-panel
// @Accept json
// @Produce json
// @Param id path string true "id or email of the user"
// @Param data query string true "Search with" Enums(id, email)
// @Param role query string true "New role of the user" Enums(admin, user)
// @Success 200 {object} string "User role updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /change-role/{id} [put]
func (h *HTTPHandler) ChangeUserRole(c *gin.Context) {
	id_or_email := c.Param("id")
	data := c.Query("data")
	role := c.Query("role")
	if data == "email" {
		err := h.US.ChangeRole(&models.ChangeRoleReq{Email: id_or_email, Role: role})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Couldn't change the role": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"User role updated": role})
	}
	if data == "id" {
		err := h.US.ChangeRole(&models.ChangeRoleReq{ID: id_or_email, Role: role})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Couldn't change the role": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"User role updated": role})
	}
}
