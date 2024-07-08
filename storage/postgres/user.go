package postgres

import (
	"database/sql"

	pb "github.com/Azizbek-Qodirov/hr_platform_auth_service/genprotos"
	"github.com/google/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) Register(user *pb.CreateUser) (*pb.UserResponse, error) {
	userid := uuid.NewString()
	query := `
		INSERT INTO users (id, name, email, password, rating, role)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.db.Exec(query, userid, user.Name, user.Email, user.Password, 0, user.Role)
	if err != nil {
		return nil, err
	}
	query = `
    INSERT INTO refresh_tokens (user_id, token)
    VALUES ($1, $2)
`
	_, err = p.db.Exec(query, userid, user.Token)
	if err != nil {
		return &pb.UserResponse{Success: false}, err
	}
	return &pb.UserResponse{Success: true, Message: "Muvaffaqiyatli!!!"}, err
}

func (p *UserStorage) Login(user *pb.LoginRequest) (*pb.LoginResponse, error) {
	query := `
    SELECT id FROM users 
    WHERE email = $1 and password=$2
`
	var id string
	err := p.db.QueryRow(query, user.Name, user.Password).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.LoginResponse{Success: false, Message: "User topilmadi!!!"}, nil
		}
		return nil, err
	}
	query = `
    SELECT token FROM refresh_tokens 
    WHERE user_id = $1
`
	var rtoken string
	err = p.db.QueryRow(query, id).Scan(&rtoken)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.LoginResponse{Success: false, Message: "Token topilmadi!!!"}, nil
		}
		return nil, err
	}	

	return &pb.LoginResponse{Success: true, Message: "Muvaffaqiyatli!!!", RefreshToken: rtoken}, nil
}
func (p *UserStorage) RefreshToken(token *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	query := `
    SELECT token FROM refresh_tokens 
	where token = $1
`
	var gettoken string
	err := p.db.QueryRow(query, token.Token).Scan(&gettoken)
	if err != nil {
		return &pb.RefreshTokenResponse{Success: false, Message: "Token topilmadi"}, err
	}

	return &pb.RefreshTokenResponse{Success: true, Token: gettoken}, nil
}

// func (p *UserStorage) Logout(token *pb.LogoutRequest) (*pb.Void, error) {
// 	query := `
//     delete FROM refresh_tokens
// 	where token = $1
// `
// 	result, err := p.db.Exec(query, token.Token)
// 	if err != nil {
// 		return nil, err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if rowsAffected == 0 {
// 		return &pb.LogoutResponse{Success: false, Message: "Token topilmadi"}, nil
// 	}

// 	return &pb.LogoutResponse{Success: true, Message: "Ochrildi"}, nil
// }

func (p *UserStorage) Update(user *pb.UpdateUser) (*pb.Void, error) {
	query := `
        update users set name=$1, email=$2, password=$3
	    where id = $4
    `
	res, err := p.db.Exec(query, user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	if rowsAffected == 0 {
		return &pb.Void{Success: false, Message: "Userlar topilmadi!!!"}, nil
	}

	return &pb.Void{Success: true, Message: "Muvaffaqiyatli!!!"}, nil
}


func (p *UserStorage) UpdateForAdmin(user *pb.UpdateUserForAdmin) (*pb.Void, error) {
	query := `
    update users set role=$1, position_id=$2
	where id = $3
`
	res, err := p.db.Exec(query, user.Role, user.PositionId, user.Id)
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}

	if rowsAffected == 0 {
		return &pb.Void{Success: false, Message: "Userlar topilmadi!!!"}, nil
	}
	query=`update refresh_tokens
		   set token = $1	
		   where user_id=$2`
	
	_, err = p.db.Exec(query, user.Token, user.Id)
	if err != nil {
		return &pb.Void{Success: false, Message: err.Error()}, err
	}
	   
	return &pb.Void{Success: true, Message: "Muvaffaqiyatli!!!"}, nil
}

func (u *UserStorage) GEtInformations(user *pb.Byid) (*pb.CreateUser, error){
	row, err:=u.db.Query("select name, email, role from users where id =$1", user.Id)
	if err!=nil{
		return nil, err
	}
	res:=pb.CreateUser{}
	row.Scan(&res.Name, &res.Email, &res.Role)

	return &res, nil
}