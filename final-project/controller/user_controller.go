package controller

import (
	"errors"
	"net/http"

	"github.com/firouzdimas/Hacktiv8-Project-/helper"
	"github.com/firouzdimas/Hacktiv8-Project-/model"
	"github.com/firouzdimas/Hacktiv8-Project-/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
}

type UserControllerImpl struct {
	userService service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		userService: service,
	}
}

// Register godoc
//
// @Summary		register user
// @Description	filled some form for registration
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		request	body		model.UserRegisterReq	true	"request is required"
// @Success		200		{object}	model.SuccessResponse{data=model.UserRegisterRes}
// @Failure		400		{object}	model.ErrorResponse{errors=interface{}}
// @Failure		500		{object}	model.ErrorResponse{errors=interface{}}
// @route		/register [post]
func (c *UserControllerImpl) Register(ctx *gin.Context) {
	request := model.UserRegisterReq{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.UserRegisterValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := c.userService.Register(request)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
				Errors: errors.New("This email or username already registered").Error(),
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User registered successfully",
		Data:    response,
	})
}

// Login godoc
//
// @Summary		login user
// @Description	login user using username and password
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		request	body		model.UserLoginReq	true	"request is required"
// @Success		200		{object}	model.SuccessResponse{data=model.UserLoginRes}
// @Failure		400		{object}	model.ErrorResponse{errors=interface{}}
// @Failure		500		{object}	model.ErrorResponse{errors=interface{}}
// @route		/login [post]
func (c *UserControllerImpl) Login(ctx *gin.Context) {
	request := model.UserLoginReq{}
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: err.Error(),
		})
		return
	}

	validateErrs := []error{}
	validateErrs = helper.UserLoginValidator(request)
	if validateErrs != nil {
		errResponseStr := make([]string, len(validateErrs))
		for i, err := range validateErrs {
			errResponseStr[i] = err.Error()
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Errors: errResponseStr,
		})
		return
	}

	response, err := c.userService.Login(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.SuccessResponse{
		Code:    http.StatusOK,
		Message: "User login successfully",
		Data: model.UserLoginRes{
			Token: *response,
		},
	})
}

func (c *UserControllerImpl) GetProfile(ctx *gin.Context) {
	userID, userIDIsExist := ctx.Get("userID")
	if !userIDIsExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: "UserID doesn't exist",
		})
		return
	}

	user, err := c.userService.GetProfile(userID.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ErrorResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.UserResponse{
		UserID:      user.UserID,
		Username:    user.Username,
		Email:       user.Email,
		Age:         user.Age,
		Photos:      user.Photos,
		SocialMedia: user.SocialMedia,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	})
}
