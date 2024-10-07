package routes

import (
	"net/http"

	"github.com/blinkinglight/prj_ds/pkg/types"
	"github.com/blinkinglight/prj_ds/template"
	"github.com/delaneyj/datastar"
	"github.com/go-chi/chi/v5"
)

func SetupTest(router chi.Router) {
	var form types.Form
	resetForm := func() {
		form = types.Form{
			Roles: &types.FormRole{
				Name:     "1",
				Category: "cat0",
				Valid:    false,
			},
		}
	}
	resetForm()
	setForm := func() {
		form = types.Form{
			Roles: &types.FormRole{
				Name:     "1",
				Category: "cat0",
				Valid:    false,
				FormRole: &types.FormRole{
					Name:     "2",
					Category: "cat1",
					Valid:    true,
				},
			},
		}
	}
	setForm()

	addForm := func() {
		var lastRole *types.FormRole
		for role := form.Roles; role != nil; role = role.FormRole {
			if role != nil {
				lastRole = role
			}
		}
		lastRole.FormRole = &types.FormRole{
			Name:     "4",
			Category: "cat3",
			Valid:    false,
		}
	}
	addForm()

	router.Route("/", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			template.Form(form).Render(r.Context(), w)
		})

		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			sse := datastar.NewSSE(w, r)
			addForm()
			datastar.RenderFragmentTempl(sse, template.Forma(form))
			datastar.RenderFragmentTempl(sse, template.Debug(form))
			datastar.RenderFragmentTempl(sse, template.S(form))
		})

		r.Delete("/reset", func(w http.ResponseWriter, r *http.Request) {
			resetForm()
			sse := datastar.NewSSE(w, r)
			_ = sse
			datastar.PatchStore(sse, form)
			datastar.RenderFragmentTempl(sse, template.Forma(form))
			datastar.RenderFragmentTempl(sse, template.Debug(form))
			datastar.RenderFragmentTempl(sse, template.S(form))
		})

		r.Patch("/patch", func(w http.ResponseWriter, r *http.Request) {
			datastar.BodyUnmarshal(r, &form)
			sse := datastar.NewSSE(w, r)
			datastar.RenderFragmentTempl(sse, template.Debug(form))
			datastar.PatchStore(sse, form)
			datastar.RenderFragmentTempl(sse, template.S(form))
		})
	})
}
