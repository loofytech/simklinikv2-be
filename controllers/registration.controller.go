package controllers

import (
	"fmt"
	"sim-klinikv2/config"
	"sim-klinikv2/models"
	"strconv"

	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRegistrationHandler(c *fiber.Ctx) error {
	var dumpPatientId int64

	var patient models.Patient
	var ppx *models.CreatePatientSchema
	var mmx *models.UpdatePatientSchema
	now := time.Now()

	payload := struct {
		MedicalRecord       string `json:"medical_record"`
		ServiceId           int64  `json:"service_id"`
		ResponsibleName     string `json:"responsible_name"`
		ResponsiblePhone    string `json:"responsible_phone"`
		ResponsibleAddress  string `json:"responsible_address"`
		ResponsibleRelation string `json:"responsible_relation"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if payload.MedicalRecord != "" {
		chkDB := config.DB.Where(&models.Patient{MedicalRecord: payload.MedicalRecord}).First(&patient)
		if chkDB.RowsAffected == 0 {
			return c.Status(400).JSON(fiber.Map{"status": "bad request", "message": "Data medical record tidak ditemukan"})
		}

		if err := c.BodyParser(&mmx); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		result := chkDB
		if err := result.Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Patient with that Id exists"})
			}
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		updates := make(map[string]interface{})
		updates["patient_name"] = mmx.PatientName
		updates["patient_address"] = mmx.PatientAddress
		updates["patient_phone"] = mmx.PatientPhone
		updates["patient_nik"] = mmx.PatientNik
		updates["patient_gender"] = mmx.PatientGender
		updates["patient_blood_type"] = mmx.PatientBloodType
		updates["birth_place"] = mmx.BirthPlace
		updates["birth_date"] = mmx.BirthPlace
		updates["province"] = mmx.Province
		updates["regency"] = mmx.Regency
		updates["district"] = mmx.District
		updates["sub_district"] = mmx.SubDistrict
		updates["updated_at"] = time.Now()

		config.DB.Model(&patient).Updates(updates)
		if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Patient already exist"})
		} else if result.Error != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
		}

		dumpPatientIds := &dumpPatientId
		*dumpPatientIds = int64(patient.ID)
	} else {
		if err := c.BodyParser(&ppx); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
		}

		errors := models.ValidateStruct(ppx)
		if errors != nil {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}

		var counting int64
		config.DB.Find(&patient).Count(&counting)

		mr := strings.Join([]string{"A", fmt.Sprintf("%06d", counting+1)}, "")

		newPatient := models.Patient{
			MedicalRecord:    mr,
			PatientName:      ppx.PatientName,
			PatientAddress:   ppx.PatientAddress,
			PatientPhone:     ppx.PatientPhone,
			PatientNik:       ppx.PatientNik,
			BirthPlace:       ppx.BirthPlace,
			BirthDate:        ppx.BirthDate,
			Province:         ppx.Province,
			Regency:          ppx.Regency,
			District:         ppx.District,
			SubDistrict:      ppx.SubDistrict,
			PatientGender:    ppx.PatientGender,
			PatientBloodType: ppx.PatientBloodType,
			JobId:            ppx.JobId,
			EthnicId:         ppx.EthnicId,
			ReligionId:       ppx.ReligionId,
			EducationId:      ppx.EducationId,
			MaritalStatusId:  ppx.MaritalStatusId,
			CreatedAt:        now,
			UpdatedAt:        now,
		}

		result := config.DB.Create(&newPatient)

		if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Patient already exist"})
		} else if result.Error != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
		}

		dumpPatientIds := &dumpPatientId
		*dumpPatientIds = int64(newPatient.ID)
	}

	newRegistration := models.Registration{
		ResponsibleName:     payload.ResponsibleName,
		ResponsiblePhone:    payload.ResponsiblePhone,
		ResponsibleAddress:  payload.ResponsibleAddress,
		ResponsibleRelation: payload.ResponsibleRelation,
		ServiceId:           payload.ServiceId,
		PatientId:           dumpPatientId,
		CreatedAt:           now,
		UpdatedAt:           now,
	}

	newReg := config.DB.Create(&newRegistration)

	if newReg.Error != nil && strings.Contains(newReg.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Patient already exist"})
	} else if newReg.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": newReg.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   fiber.Map{"Registration": "registration OK"},
	})

}

func FindRegistration(c *fiber.Ctx) error {
	var page = c.Query("page", "1")
	var limit = c.Query("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var registration []models.Registration
	results := config.DB.Limit(intLimit).Offset(offset).Find(&registration)
	if results.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": results.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "results": len(registration), "user": registration})
}

func UpdateRegistration(c *fiber.Ctx) error {
	// registrationId := c.Params("registrationId")

	// var payload *models.UpdateRegistrationSchema

	// if err := c.BodyParser(&payload); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	// }

	// var registration models.Registration
	// result := config.DB.First(&registration, "id = ?", registrationId)
	// if err := result.Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Registration with that Id exists"})
	// 	}
	// 	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	// }

	// updates := make(map[string]interface{})
	// updates["updated_at"] = time.Now()

	// config.DB.Model(&registration).Updates(updates)

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": registration}})
	return c.Status(200).JSON(fiber.Map{})
}

func FindRegistrationById(c *fiber.Ctx) error {
	registrationId := c.Params("registrationId")

	var registration models.Registration
	result := config.DB.First(&registration, "id = ?", registrationId)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Registration with that Id exists"})
		}
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": registration}})
}

func RegistrationDelete(c *fiber.Ctx) error {
	registrationId := c.Params("registrationId")

	result := config.DB.Delete(&models.Registration{}, "id = ?", registrationId)

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "fail", "message": "No Registration with that Id exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
	}

	return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error})
}
