package common

import (
	"fmt"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/gavv/httpexpect"
	"net/http"
)

func OperationLogTest(e *httpexpect.Expect, sesId *http.Cookie) {

	fmt.Println()
	printlnWithColor("Operation Log", "blue")
	fmt.Println("============================")

	// show

	printlnWithColor("show", "green")
	e.GET(config.Get().Url("/info/op")).
		WithCookie(sesId.Name, sesId.Value).
		Expect().
		Status(200).
		Body().Contains(language.Get("operation log"))
}
