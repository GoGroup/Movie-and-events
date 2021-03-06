package permission

import (
	"fmt"
	"strings"
)

type permission struct {
	roles   []string
	methods []string
}

type authority map[string]permission

var authorities = authority{
	"/logout": permission{
		roles:   []string{"USER", "ADMIN"},
		methods: []string{"POST"},
	},
	"/admin": permission{
		roles:   []string{"ADMIN"},
		methods: []string{"GET", "POST"},
	},
}

// HasPermission checks if a given role has permission to access a given route for a given method
func HasPermission(path string, role string, method string) bool {
	if strings.HasPrefix(path, "/admin") {
		path = "/admin"
	}
	fmt.Println("authorities")
	perm := authorities[path]
	fmt.Println(perm)
	fmt.Println("path")
	fmt.Println(path)
	checkedRole := checkRole(role, perm.roles)
	fmt.Println("role")
	fmt.Println(checkedRole)
	checkedMethod := checkMethod(method, perm.methods)
	fmt.Println("method")
	fmt.Println(checkedMethod)
	if !checkedRole || !checkedMethod {
		return false
	}
	return true
}

func checkRole(role string, roles []string) bool {
	for _, r := range roles {
		if strings.ToUpper(r) == strings.ToUpper(role) {
			return true
		}
	}
	return false
}

func checkMethod(method string, methods []string) bool {
	for _, m := range methods {
		if strings.ToUpper(m) == strings.ToUpper(method) {
			return true
		}
	}
	return false
}
