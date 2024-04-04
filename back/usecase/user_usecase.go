package usecase

import (
	"time"

	"github.com/akiradomi/workspace/go-practice/back/model"
	"github.com/akiradomi/workspace/go-practice/back/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseInterface interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.UserRepositoryInterFace
}

// コンストラクタ,main.goから呼び出される
func NewUserUsecase(ur repository.UserRepositoryInterFace) UserUsecaseInterface {
	return &userUsecase{ur}
}

func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	//userが入力したpasswordをハッシュ化する
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	//入力されたemailのユーザーがいるかチェック
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	//入力されたpasswordとDB内のpasswordを比較
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	//JWTの発行
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 6).Unix(),
	})
	//トークンに署名を付与
	tokenString, err := token.SignedString([]byte("myfirstgoproject"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
