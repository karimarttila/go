package userdb

import (
	"errors"
	"github.com/karimarttila/go/simpleserver/app/util"
	"hash/fnv"
	"strconv"
)

type User struct {
	UserId         int
	email          string
	firstName      string
	lastName       string
	hashedPassword string
}

type AddUserResponse struct {
	Ret   string
	Email string
}

type UsersDb struct {
	usersMap map[int]User
}

func hashString(myStr string) string {
	algorithm := fnv.New32a()
	algorithm.Write([]byte(myStr))
	hashedInt := algorithm.Sum32()
	ret := strconv.Itoa(int(hashedInt))
	return ret
}

// UsersDB singleton.
var myUsersDB = initUsersDb()

var nextId = createCounter()

// NOTE: Counter and Go closure example (compare to equivalent Python closure, practically the same).
func createCounter() func() int {
	i := 3 // Initialize to test user count.
	return func() int {
		i++
		return i
	}
}

func initUsersDb() UsersDb {
	util.LogEnter()
	testUser1 := User{1, "kari.karttinen@foo.com", "Kari", "Karttinen", "2842551024"}
	testUser2 := User{2, "timo.tillinen@foo.com", "Timo", "Tillinen", "3655654034"}
	testUser3 := User{3, "erkka.erkkila@foo.com", "Erkka", "Erkkila", "2077629983"}
	userMap := make(map[int]User)
	userMap[1] = testUser1
	userMap[2] = testUser2
	userMap[3] = testUser3
	ret := UsersDb{userMap}
	util.LogExit()
	return ret
}

func EmailAlreadyExists(givenEmail string) bool {
	util.LogEnter()
	ret := false
	usersMap := myUsersDB.usersMap
	for _, user := range usersMap {
		if user.email == givenEmail {
			ret = true
			break
		}
	}
	util.LogExit()
	return ret
}

func AddUser(email string, firstName string, lastName string, password string) (ret AddUserResponse, err error) {
	util.LogEnter()
	if EmailAlreadyExists(email) {
		buf := "Email already exists: " + email
		util.LogWarn(buf)
		err = errors.New(buf)
	} else {
		id := nextId()
		newUser := User{id, email, firstName, lastName, hashString(password)}
		myUsersDB.usersMap[id] = newUser
		ret = AddUserResponse{"ok", email}
	}
	util.LogExit()
	return ret, err
}

func CheckCredentials(userEmail string, userPassword string) bool {
	util.LogEnter()
	ret := false
	usersMap := myUsersDB.usersMap
	for _, user := range usersMap {
		if user.email == userEmail && user.hashedPassword == hashString(userPassword) {
			ret = true
			break
		}
	}
	util.LogExit()
	return ret
}
