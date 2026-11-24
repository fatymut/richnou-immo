
package routes

import (
    "net/http"
    "richnou-immo/agence/handler"
    "richnou-immo/agence/middleware"
)

func RegisterAdminRoutes(mux *http.ServeMux) {

    mux.Handle(
        "GET /admin/export",
        middleware.AuthMiddleware(
            middleware.RequireRole("admin")(
                http.HandlerFunc(handler.ExportAnnonces),
            ),
        ),
    )

    mux.Handle(
        "POST /admin/import",
        middleware.AuthMiddleware(
            middleware.RequireRole("admin")(
                http.HandlerFunc(handler.ImportAnnonces),
            ),
        ),
    )
}
