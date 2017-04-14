package main

import (
	"log"

	"github.com/kr/pretty"
)

type User struct {
	Id   int
	Name string
}

type RoleUser struct {
	Role Role
	User User
}

type Role struct {
	Id   int
	Name string
}

type RoleResource struct {
	Role     Role
	Resource Resource
	Value    int
}

type Resource struct {
	Id   int
	Name string
}

type ResourceOperation struct {
	Resource  Resource
	Operation Operation
}

type Operation struct {
	Id    int
	Name  string
	Value int
}

var (
	users              []User
	roles              []Role
	resources          []Resource
	operations         []Operation
	roleUsers          []RoleUser
	resourceOperations []ResourceOperation
	roleResources      []RoleResource
)

func init() {
	users = []User{
		User{1, "jimmy"},
		User{2, "qgm"},
	}

	roles = []Role{
		Role{1, "admin"},
		Role{2, "saler"},
	}

	resources = []Resource{
		Resource{1, "user"},
		Resource{2, "blog"},
	}

	operations = []Operation{
		Operation{1, "show", 1},
		Operation{2, "create", 1 << 2},
		Operation{3, "update", 1 << 3},
		Operation{4, "delete", 1 << 4},
		Operation{5, "import", 1 << 5},
		Operation{6, "download", 1 << 6},
	}

	resourceOperations = []ResourceOperation{}
	resourceOperations = append(resourceOperations, AddResourceOperation(resources[0], []Operation{operations[0], operations[2]})...)
	resourceOperations = append(resourceOperations, AddResourceOperation(resources[1], []Operation{operations[0], operations[2]})...)

	roleResources = []RoleResource{}
	roleResources = append(roleResources, AddRoleResource(roles[0], resources[0], []Operation{operations[0], operations[1]}))

	roleUsers = []RoleUser{}
	roleUsers = append(roleUsers, AssignRole(users[0], roles[0]))
}

func main() {
	yes := HasPermission(users[0], resources[0], operations[0])
	log.Println(yes)
}

func HasPermission(user User, res Resource, op Operation) bool {
	role := findRole(user)
	log.Printf("role: %# v\n", pretty.Formatter(role))
	roleResource := findRoleResource(role, res)
	if op.Value&roleResource.Value == 0 {
		return false
	}
	return true
}
func findRole(user User) Role {
	for _, ru := range roleUsers {
		if ru.User.Id == user.Id {
			return ru.Role
		}
	}
	return Role{}
}

func findRoleResource(role Role, res Resource) RoleResource {
	for _, rr := range roleResources {
		if rr.Role.Id == role.Id && rr.Resource.Id == res.Id {
			return rr
		}
	}
	return RoleResource{}
}

func AssignRole(user User, role Role) RoleUser {
	return RoleUser{
		User: user,
		Role: role,
	}
}

func AddRoleResource(role Role, res Resource, ops []Operation) RoleResource {
	num := 0
	for _, op := range ops {
		num += op.Value
	}
	return RoleResource{
		Role:     role,
		Resource: res,
		Value:    num,
	}
}

func AddResourceOperation(res Resource, ops []Operation) []ResourceOperation {
	ro := []ResourceOperation{}
	for _, op := range ops {
		ro = append(ro, ResourceOperation{res, op})
	}
	return ro
}
