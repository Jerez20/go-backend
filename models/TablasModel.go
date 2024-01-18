package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccionesRH struct {
	gorm.Model
	Id               string    `gorm:"type:varchar(12);primaryKey;not null"`
	Fecha            time.Time `gorm:"null"`
	Empleado         string    `gorm:"type:varchar(10);not null"`
	TipoNovedad      string    `gorm:"type:varchar(3);not null"`
	FechaEfectividad time.Time `gorm:"null"`
	SueldoAnterior   float64   `gorm:"null"`
	Sueldo           float64   `gorm:"null"`
	Observaciones    string    `gorm:"type:text;not null"`
	Fecha_salida     time.Time `gorm:"null"`
	Fecha_entrada    time.Time `gorm:"null"`
	Dia              int       `gorm:"null"`
	Id_regreso       string    `gorm:"type:varchar(12);not null"`
}

type Usuario struct {
	gorm.Model
	Usuario string `gorm:"type:varchar(30) not null"`  //varchar(30)) NOT NULL,
	Clave   string `gorm:"type:varchar(200) not null"` //varchar(200)) NOT NULL,
	Dire    string `gorm:"type:varchar(40) null"`      //varchar(40)) NULL,,
	Tele    string `gorm:"type:varchar(20) null"`      //varchar(20)) NULL,,
	Correo  string `gorm:"type:varchar(80) null"`      //varchar(80)) NULL,,
}

// Antes de guardar un nuevo usuario o actualizar uno existente,
// puedes llamar a CreatePassword para encriptar la contraseña.

func (a *AccionesRH) BeforeCreate(tx *gorm.DB) (err error) {
	a.Id = uuid.New().String()
	return
}
