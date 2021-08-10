package main

import (
	"database/sql"
	"net/http"

	"github.com/emahiro/il/rightingSoftware/arch/engine"
	"github.com/emahiro/il/rightingSoftware/arch/handler"
	"github.com/emahiro/il/rightingSoftware/arch/manager"
	"github.com/emahiro/il/rightingSoftware/arch/resource"
	"github.com/emahiro/il/rightingSoftware/arch/util"
)

func main() {
	db := &sql.DB{}
	// db, err := sql.Open("$DatabaseName", "$SourceName")
	// if err != nil {
	// 	panic(err)
	// }

	r := http.NewServeMux()
	user := handler.NewUserHandler(
		manager.NewUserManager(
			engine.NewRegisterEngine(resource.NewUserAccess(db)),
			engine.NewUserSearchEngine(resource.NewUserAccess(db)),
			engine.NewMailEngine(),
		),
	)
	r.HandleFunc("/user/register", user.Create)

	if err := http.ListenAndServe(":8080", util.Logger(r)); err != nil {
		panic(err)
	}
}
