package routes

import (
	"net/http"
	"strconv"

	"github.com/blinkinglight/prj_ds/pkg/types"
	"github.com/blinkinglight/prj_ds/template"
	"github.com/delaneyj/datastar"
	"github.com/go-chi/chi/v5"
)

func SetupTest(router chi.Router) {
	var form = types.Form{
		Roles: make([]types.FormRole, 3),
	}

	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			template.Form(form).Render(r.Context(), w)
		})
		r.Patch("/patch", func(w http.ResponseWriter, r *http.Request) {
			var ds types.DataStore
			datastar.BodyUnmarshal(r, &ds)
			switch ds.Elf {
			case "name":
				form.Name = ds.Elv
			case "roles":
				i, _ := strconv.Atoi(ds.Elr)
				form.Roles[i].Name = ds.Elv
			}
			sse := datastar.NewSSE(w, r)
			_ = sse
			datastar.RenderFragmentTempl(sse, template.Debug(form))
		})
	})
}
