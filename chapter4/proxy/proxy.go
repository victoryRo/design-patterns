package proxy

import (
	"fmt"
)

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

// ---------------------------------------------------

// UserList RealObject
type UserList []User

// FindUser will iterate over the list to try to find a user with the same
// name that the param or return an error if it can't find it
func (l *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*l); i++ {
		if (*l)[i].ID == id {
			return (*l)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be found\n", id)
}

// addUser adds a new user to the end of the Users slice
func (l *UserList) addUser(newUser User) {
	*l = append(*l, newUser)
}

// ---------------------------------------------------

// UserListProxy Proxy
type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackSize {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}

// FindUser will search for the specified name in the parameter in the cache
// list. If it finds it, it will return it. If not, it will search in the heavy
// list. Finally, if it's not in the heavy list, it will return an error
// (generated from the heavy list)
func (u *UserListProxy) FindUser(id int32) (User, error) {
	//Search for the object in the cache list first
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.LastSearchUsedCache = true
		return user, nil
	}

	//Object is not in the cache list. Search in the heavy list
	user, err = u.MockedDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	//Adds the new user to the stack, removing the last if necessary
	u.addUserToStack(user)

	fmt.Println("Returning user from database")
	u.LastSearchUsedCache = false
	return user, nil
}
