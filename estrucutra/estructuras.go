package estructuras

import "fmt"

type User struct {
	Id      int
	Name    string
	Age     int
	friends []User
}

func (u User) SayHello() {
	fmt.Println("hello", u.Name)
}

func (u *User) AddFriends(friend User) {
	u.friends = append(u.friends, friend)
}

func (u *User) ListFriends() {
	for i, friend := range u.friends {
		fmt.Printf("%d - %s\n", i+1, friend.Name)
	}
}

func (u *User) RemoveFriend(id int) {
	index := u.find(id)
	if index == -1 {
		return
	}
	u.friends = append(u.friends[:index], u.friends[index+1:]...)
}

func (u User) find(id int) int {
	for i, f := range u.friends {
		if f.Id == id {
			return i
		}
	}
	return -1
}
