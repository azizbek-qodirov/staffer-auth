package handler

import (
	"log"

	pb "github.com/Azizbek-Qodirov/hr_platform_auth_service/genprotos"

	"github.com/Azizbek-Qodirov/hr_platform_auth_service/api/token"
	"github.com/gin-gonic/gin"
)

// Registr 			handles the creation of a new Registr
// @Summary 		Create Registr
// @Description 	Create page
// @Tags 			Users
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param   		Create  body    pb.CreateUser    true   "Create"
// @Success 		200   {string}  string    "Create Successful"
// @Failure 		401   {string}  string    "Error while Created"
// @Router 			/user/register [post]
func (h *Handler) Register(ctx *gin.Context) {
	arr := &pb.CreateUser{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	arr.Role = "user"
	t := token.GenereteJWTToken(arr)
	arr.Token = t.RefreshToken
	res, err := h.Auth.Register(ctx, arr)
	if err != nil {
		panic(err)
	}
	res.RefreshToken = t.RefreshToken
	res.AccessToken = t.AccessToken
	ctx.JSON(200, res)
}

// Logout 			handles the creation of a new Logout
// @Summary 		Delete Logout
// @Description 	Delete page
// @Tags 			Users
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param   		Logout  body    pb.LogoutRequest    true   "Logout"
// @Success 		200 {string}  string  "Logout Successful"
// @Failure 		401 {string}  string  "Error while Deleted"
// @Router 			/user/logout [delete]
func (h *Handler) Logout(ctx *gin.Context) {
	arr := pb.LogoutRequest{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		log.Fatal(err)
	}
	res, err := h.Auth.Logout(ctx, &arr)
	if err != nil {
		panic(err)
	}
	ctx.JSON(200, res)
}

// Login 		    handles the creation of a new Login
// @Summary 		Login
// @Description 	Create page
// @Tags 			Users
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param   		Create  body    pb.LoginRequest    true   "Create"
// @Success 		200   {string}  string    "Create Successful"
// @Failure 		401   {string}  string    "Error while Created"
// @Router 			/user/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	arr := pb.LoginRequest{}
	err := ctx.BindJSON(&arr)
	if err != nil {
		panic(err)
	}
	res, err := h.Auth.Login(ctx, &arr)
	if err != nil {
		panic(err)
	}
	if !res.Success {
		ctx.JSON(200, res)
		return
	}

	ctx.JSON(200, res)
}

// Refreshtoken		handles the creation of a new Login
// @Summary 		Login
// @Description 	Create page
// @Tags 			Users
// @Accept  		json
// @Produce  		json
// @Security  		BearerAuth
// @Param   		query query   pb.RefreshTokenRequest    true   "Create"
// @Success 		200   {string}  pb.RefreshTokenRequest   "Successful"
// @Failure 		401   {string}  string    "Error while Created"
// @Router 			/user/refreshTokenRequest [get]
func (h *Handler) RefreshToken(ctx *gin.Context) {
	arr := pb.RefreshTokenRequest{}
	arr.Token = ctx.Query("token")
	res, err := h.Auth.RefreshToken(ctx, &arr)
	if err != nil {
		panic(err)
	}

	claims, err:=token.ExtractClaim(arr.Token)
	if err!=nil{
		newtoken:=token.GenereteJWTToken(&pb.CreateUser{Name: claims["username"].(string), 
		Email: claims["email"].(string), Role: claims["role"].(string)})
		ctx.JSON(200, newtoken.RefreshToken)
		return
	}
	ctx.JSON(200, res)
}

// UpdateForAdmin updates user information for admin
// @Summary Update user for admin
// @Description Update user details for admin
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body pb.UpdateUserForAdmin true "Update User For Admin"
// @Success 200 {object} pb.Void "Update Successful"
// @Failure 400 {string} string "Error while updating"
// @Router /user/admin/update [put]
func (u *Handler) UpdateForAdmin(ctx *gin.Context) {
	user := pb.UpdateUserForAdmin{}
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	arr, err:=u.Auth.GEtInformations(ctx, &pb.Byid{Id: user.Id})
	if err != nil {
		log.Fatal(err)
	}
	t:=token.GenereteJWTToken(&pb.CreateUser{Name: arr.Name, Role: arr.Role, Email: arr.Email})
	user.Token=t.RefreshToken
	res, err := u.Auth.UpdateForAdmin(ctx, &user)
	if err != nil {
		panic(err)
	}
	if !res.Success {
		ctx.JSON(400, res)
	}
	ctx.JSON(200, res)
}

// Update updates user information
// @Summary Update user
// @Description Update user details
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body pb.UpdateUser true "Update User"
// @Success 200 {object} pb.Void "Update Successful"
// @Failure 400 {string} string "Error while updating"
// @Router /user/update [put]
func (u *Handler) Update(ctx *gin.Context) {
	user := pb.UpdateUser{}
	err := ctx.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}
	res, err := u.Auth.Update(ctx, &user)
	if err != nil {
		panic(err)
	}
	if !res.Success {
		ctx.JSON(400, res)
	}
	ctx.JSON(200, res)
}
