package controller

import (
	"Assignment2/Model/user"
	"Assignment2/Web/dbhelper"
	"errors"
)

func AllUser() ([]user.User, error) {
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return nil, err
	}
	var User []user.User
	res := connection.Where("status = 0").Find(&User)
	if res.Error != nil {
		return nil, errors.New("error retrieving User: " + res.Error.Error())
	}
	if len(User) == 0 {
		return nil, errors.New("no User found")
	}
	return User, nil
}
func AddUser(item *user.User) (string, error) {
	if item.Name == "" {
		return "name", errors.New("name is required")
	}
	if item.Surname == "" {
		return "surname", errors.New("surname is required")
	}
	if item.Email == "" {
		return "email", errors.New("email is required")
	}
	if item.Password == "" {
		return "password", errors.New("password is required")
	}
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}
	if item.ID <= 0 {
		res := connection.Create(&user.User{
			Name:     item.Name,
			Surname:  item.Surname,
			Age:      item.Age,
			Birthday: item.Birthday,
			Job:      item.Job,
			Email:    item.Email,
			Password: item.Password,
			Status:   item.Status,
		})
		if res.Error != nil {
			return "error", res.Error
		}

	}
	return "success", nil
}

func UpdateUser(item *user.User) (string, error) {

	if item.ID <= 0 {
		return "id", errors.New("id is required")
	}
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}

	if item.ID >= 0 {
		var user user.User
		res := connection.Where("status = 0 AND id = ?", item.ID).First(&user)

		if res.Error != nil {
			return "error", errors.New("user not found")
		} else {
			user.Name = item.Name
			user.Surname = item.Surname
			user.Age = item.Age
			user.Birthday = item.Birthday
			user.Job = item.Job
			user.Email = item.Email
			user.Password = item.Password
			user.Status = item.Status
			res = connection.Save(&user)
			if res.Error != nil {
				return "error", res.Error
			}
		}
	}
	return "success", nil
}

func DeleteUser(item *user.User) (string, error) {
	if item.ID <= 0 {
		return "id", errors.New("id is required")
	}
	connection, err := dbhelper.GetOpenConnection()
	if err != nil {
		return "error", err
	}

	if item.ID >= 0 {
		var user user.User
		res := connection.Where("status = 0 AND id =? ", item.ID).First(&user)

		if res.Error != nil {
			return "error", errors.New("user not found")
		} else {
			user.Status = 1
			res = connection.Save(&user)
			if res.Error != nil {
				return "error", res.Error
			}
		}
	}
	return "success", nil
}
