package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type postReservationJson struct {
	TicketID   string `json:"ticket_id"`
	ScheduleID string `json:"schedule_id"`
}

func (h *Handler) handlePostReservation(w http.ResponseWriter, r *http.Request) {

	var j postReservationJson
	err := render.DecodeJSON(r.Body, &j)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}

	err = h.ReservationUsecase.Create(r.Context(), j.TicketID, j.ScheduleID)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		return
	}
	render.Respond(w, r, "ok")
}
