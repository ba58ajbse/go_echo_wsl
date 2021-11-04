package handler

import (
	"database/sql"
	"echo_app/app/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

var db *sql.DB

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 全件取得
func GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		db = database.Connect()
		defer db.Close()

		users := []User{}
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			user := User{}
			err := rows.Scan(&user.Id, &user.Name, &user.Email)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			users = append(users, user)
		}

		return c.JSON(http.StatusOK, users)
	}
}

// 1件取得
func Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		db = database.Connect()
		defer db.Close()

		user := User{}
		id := c.Param("id")

		err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.Name, &user.Email)

		switch {
		case err == sql.ErrNoRows:
			return c.JSON(http.StatusBadRequest, "invalid user id")
		case err != nil:
			return c.JSON(http.StatusBadRequest, err.Error())
		default:
			return c.JSON(http.StatusOK, user)
		}
	}
}

// 新規作成
func Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		db = database.Connect()
		defer db.Close()

		u := new(User)

		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer tx.Rollback()

		stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer stmt.Close()

		res, err := stmt.Exec(u.Name, u.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// 追加したレコードのid
		LastID, err := res.LastInsertId()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, LastID)
	}
}

// 更新
func Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		db = database.Connect()
		defer db.Close()

		u := new(User)
		id := c.Param("id")

		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer tx.Rollback()

		stmt, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer stmt.Close()

		res, err := stmt.Exec(u.Name, u.Email, id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// 更新したレコード数
		rowCnt, err := res.RowsAffected()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if rowCnt == 0 {
			return c.JSON(http.StatusBadRequest, "invalid user id")
		}

		return c.JSON(http.StatusOK, rowCnt)
	}
}

// 削除
func Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		db = database.Connect()
		defer db.Close()

		id := c.Param("id")

		tx, err := db.Begin()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer tx.Rollback()

		stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer stmt.Close()

		res, err := stmt.Exec(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		rowCnt, err := res.RowsAffected()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if rowCnt == 0 {
			return c.JSON(http.StatusBadRequest, "invalid user id")
		}

		return c.JSON(http.StatusOK, id)
	}
}
