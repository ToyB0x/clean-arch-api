package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (h *Handler) handleGetScheduleByMonth(w http.ResponseWriter, r *http.Request) {
	p, err := extractDateParams(r, []string{"year", "month"})
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}

	reservations, err := h.ScheduleUsecase.GetByMonth(r.Context(), p[0], p[1])
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, reservations)
}

func (h *Handler) handleGetScheduleByMonthMemStore(w http.ResponseWriter, r *http.Request) {
	p, err := extractDateParams(r, []string{"year", "month"})
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}

	reservations, err := h.ScheduleUsecase.GetByMonthMemStore(r.Context(), p[0], p[1])
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, reservations)
}

func (h *Handler) handleGetScheduleByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "schedule_id")
	reservations, err := h.ScheduleUsecase.GetByID(r.Context(), id)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, reservations)
}
