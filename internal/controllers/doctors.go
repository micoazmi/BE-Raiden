package controllers

import (
	"beraiden/internal/models"

	"github.com/sev-2/raiden"
)

type DoctorController struct {
	raiden.ControllerBase
	Http  string `path:"/doctors" type:"rest"`
	Model models.Doctors
}
