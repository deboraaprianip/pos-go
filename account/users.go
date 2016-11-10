package account

import(
    "database/sql"
)

type User struct {
    Id int
    Username string
    Password string
}

func GetAllUser( DB *sql.DB  ) ( *[]User ) {
    rows, _ := DB.Query("select id,username,password from users;")
    defer rows.Close()
    
    users := &[]User{}
    for rows.Next(){
        user := User{}
        rows.Scan(&user.Id, &user.Username, &user.Password)
        *users = append(*users, user)
    }
    return users
}
