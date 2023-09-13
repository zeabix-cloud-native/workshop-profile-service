package httpv1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/domain"
	"github.com/zeabix-cloud-native/workshop-profile-service/internal/core/ports"

	"encoding/json"
)

type CreateProfileRequest struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Avatar    string `json:"avatar"`
}

type ProfileResponse struct {
	CreateProfileRequest
	ID string `json:"id"`
}

type Handler struct {
	s   ports.ProfileService
	app *fiber.App
}

func NewHttpHandlerV1(s ports.ProfileService) *Handler {
	return &Handler{
		s:   s,
		app: fiber.New(),
	}
}

func (h *Handler) Serve(port string) error {
	v1 := h.app.Group("v1")
	v1.Post("/profiles", h.createProfileHandler)
	v1.Get("/profiles/:id", h.getProfileHandler)

	return h.app.Listen(port)
}

func (h *Handler) getProfileHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return fiber.ErrBadRequest
	}

	p, err := h.s.GetProfile(id)

	if err == ports.ErrProfileNotFound {
		return fiber.ErrNotFound
	}

	res := new(ProfileResponse)
	res.ID = p.ID
	res.Firstname = p.Firstname
	res.Lastname = p.Lastname
	res.Username = p.Username
	res.Avatar = p.Avatar

	resStr, err := json.Marshal(res)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusCreated).SendString(string(resStr))
	return nil

}

func (h *Handler) createProfileHandler(c *fiber.Ctx) error {
	req := new(CreateProfileRequest)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	p := new(domain.UserProfile)
	p.Username = req.Username
	p.Firstname = req.Firstname
	p.Lastname = req.Lastname
	p.Avatar = req.Avatar

	if err := h.s.CreateProfile(p); err != nil {
		return fiber.ErrInternalServerError
	}

	res := new(ProfileResponse)
	res.ID = p.ID
	res.Firstname = p.Firstname
	res.Lastname = p.Lastname
	res.Username = p.Username
	res.Avatar = p.Avatar

	resStr, err := json.Marshal(res)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusCreated).SendString(string(resStr))
	return nil
}
