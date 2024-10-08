package routes

import (
	"net/http"
	"strconv"

	"github.com/blinkinglight/prj_ds/pkg/types"
	"github.com/blinkinglight/prj_ds/template"
	"github.com/delaneyj/datastar"
	"github.com/delaneyj/toolbelt"
	"github.com/go-chi/chi/v5"
	"github.com/ituoga/toolbox"
)

func SetupTest(router chi.Router) {
	var form types.Form

	resetForm := func() {
		form = types.Form{
			Count: 1,
			Roles: &toolbox.LinkedList[types.FormRole]{},
		}
		form.Roles = &toolbox.LinkedList[types.FormRole]{}
		form.Roles.Add(types.FormRole{
			Id: toolbelt.NextID(),
		})
	}

	resetForm()

	router.Route("/", func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			template.Form(form).Render(r.Context(), w)

		})

		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			for i := 0; i < form.Count; i++ {
				form.Roles.Add(types.FormRole{
					Id: toolbelt.NextID(),
				})
			}
			sse := datastar.NewSSE(w, r)
			datastar.RenderFragmentTempl(sse, template.Forma(form), datastar.WithoutViewTransitions())
		})

		r.Delete("/reset", func(w http.ResponseWriter, r *http.Request) {
			resetForm()
			sse := datastar.NewSSE(w, r)
			datastar.RenderFragmentTempl(sse, template.Forma(form), datastar.WithoutViewTransitions())
		})

		r.Delete("/remove/{id}", func(w http.ResponseWriter, r *http.Request) {
			tmpID := chi.URLParam(r, "id")
			id, _ := strconv.ParseInt(tmpID, 10, 64)

			form.Roles.RemoveById(func(role types.FormRole) bool {
				return role.Id == id
			})

			sse := datastar.NewSSE(w, r)
			datastar.RenderFragmentTempl(sse, template.Forma(form), datastar.WithoutViewTransitions())
		})

		r.Patch("/patch", func(w http.ResponseWriter, r *http.Request) {
			datastar.BodyUnmarshal(r, &form)
			sse := datastar.NewSSE(w, r)
			datastar.RenderFragmentTempl(sse, template.Forma(form))
		})
	})
}
