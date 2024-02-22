package repo

import (
	"30New/entity"
	"strconv"
)

type Storage map[int]*entity.User

func (r Storage) Create(name string, age int) string {
	newUser := entity.NewUser()
	newUser.Name = name
	newUser.Age = age
	r[len(r)] = newUser
	return "Пользователь id: " + strconv.Itoa(len(r)) + " создан"
}

func (r Storage) Delete(id int) (string, bool) {
	if id > len(r) {
		return "Пользователя id: " + strconv.Itoa(id) + " не существует", false
	}
	name := r[id].Name
	delete(r, id)
	return "Пользователь id: " + name + " удален", true
}

func (r Storage) AddFriend(first, second int) (string, bool) {
	if first > len(r) {
		return "Пользователь " + strconv.Itoa(first) + " не существует", false
	}
	if second > len(r) {
		return "Пользователь " + strconv.Itoa(second) + " не существует", false
	}
	r[first].Friends = append(r[first].Friends, second)
	r[second].Friends = append(r[second].Friends, first)
	return r[first].Name + " и " + r[second].Name + " теперь друзья", true
}

func (r Storage) GetFriends(id int) (string, bool) {
	if id > len(r) {
		return "Пользователь " + strconv.Itoa(id) + " не существует", false
	}
	result := "Друзья пользователя" + strconv.Itoa(id) + ":\n"
	for _, iduser := range r[id].Friends {
		result += "Name is: " + r[iduser].Name + " and age: " + strconv.Itoa(r[iduser].Age)
	}
	return result, true
}

func (r Storage) UpdateAge(id, newade int) (string, bool) {
	if id > len(r) {
		return "Пользователь " + strconv.Itoa(id) + " не существует", false
	}
	r[id].Age = newade
	return "Возраст пользователя" + strconv.Itoa(id) + " успешно обновлен", true
}

func New() *Storage {
	return &Storage{}
}
