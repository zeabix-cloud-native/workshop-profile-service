package httpv2

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
	DOB       string `json:"dob"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	OID       string `json:"oid"`
}

type ProfileResponse struct {
	CreateProfileRequest
	ID string `json:"id"`
}

type Handler struct {
	s ports.ProfileService
	g fiber.Router
}

func NewHttpHandlerV2(s ports.ProfileService, g fiber.Router) *Handler {
	return &Handler{
		s: s,
		g: g,
	}
}

func (h *Handler) Initialize() error {
	h.g.Post("/profiles", h.createProfileHandler)
	h.g.Get("/profiles/:oid", h.getProfileHandler)
	return nil
}

func (h *Handler) getProfileHandler(c *fiber.Ctx) error {
	oid := c.Params("oid")
	p, err := h.s.GetProfileByOID(oid)
	if err == ports.ErrProfileNotFound {
		return fiber.ErrNotFound
	}

	if err != nil {
		return fiber.ErrInternalServerError
	}

	// Serialized Response
	res := new(ProfileResponse)
	res.ID = p.ID
	res.Username = p.Username
	res.Firstname = p.Firstname
	res.Lastname = p.Lastname
	res.Avatar = p.Avatar
	res.DOB = p.DOB
	res.Mobile = p.Mobile
	res.Address = p.Address
	res.OID = p.OID

	resStr, err := json.Marshal(res)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	c.Status(fiber.StatusOK).SendString(string(resStr))
	c.Set("content-type", "application/json")
	c.Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Cache-Control")
	return nil
}

func (h *Handler) createProfileHandler(c *fiber.Ctx) error {
	req := new(CreateProfileRequest)
	if err := c.BodyParser(req); err != nil {
		return fiber.ErrBadRequest
	}

	// Create Domain object
	p := new(domain.UserProfile)
	p.Username = req.Username
	p.Firstname = req.Firstname
	p.Lastname = req.Lastname
	p.Avatar = req.Avatar
	p.DOB = req.DOB
	p.Mobile = req.Mobile
	p.Address = req.Address
	p.OID = req.OID

	if err := h.s.CreateProfile(p); err != nil {
		return fiber.ErrInternalServerError
	}

	// Serialize Response
	res := new(ProfileResponse)
	res.ID = p.ID
	res.Firstname = p.Firstname
	res.Lastname = p.Lastname
	res.Username = p.Username
	res.Avatar = p.Avatar
	res.DOB = p.DOB
	res.Mobile = p.Mobile
	res.Address = p.Address
	res.OID = p.OID

	// resStr, err := json.Marshal(res)
	// if err != nil {
	// 	return fiber.ErrInternalServerError
	// }

	c.Status(fiber.StatusCreated).JSON(res)
	c.Set("content-type", "application/json; charset=utf-8")
	c.Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Cache-Control")
	return nil
}
