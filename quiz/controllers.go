package quiz

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dennybiasiolli/go-quiz/common"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func QuizList(c *fiber.Ctx) error {
	db := common.GetDB()
	var quizzes []Quiz
	var count int64
	err := db.Find(&quizzes).Count(&count).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{
		"results": QuizzesSerializer(quizzes),
		"count":   count,
	})
}

func QuizDetail(c *fiber.Ctx) error {
	db := common.GetDB()
	var quiz Quiz
	id := c.Params("id")
	err := db.Preload("Questions").First(&quiz, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(QuizSerializer(quiz))
}

func QuizCreate(c *fiber.Ctx) error {
	input := new(Quiz)
	c.BodyParser(input)
	if err := validator.New().Struct(*input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	db := common.GetDB()
	if err := db.Omit(clause.Associations).Create(&input).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(QuizSerializer(*input))
}

func QuizUpdate(c *fiber.Ctx) error {
	input := new(Quiz)
	c.BodyParser(input)
	if err := validator.New().Struct(*input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id := c.Params("id")
	if input.ID == 0 || fmt.Sprint(input.ID) != id {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "URL ID and Body ID are not the same",
		})
	}

	db := common.GetDB()
	res := db.Omit(clause.Associations).Updates(&input)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	} else if res.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to find selected record",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func QuizDelete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db := common.GetDB()
	res := db.Delete(&Quiz{}, id)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	} else if res.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to find selected record",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func QuestionCreate(c *fiber.Ctx) error {
	input := new(Question)
	c.BodyParser(input)
	if err := validator.New().Struct(*input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	db := common.GetDB()
	if err := db.Omit(clause.Associations).Create(&input).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(QuestionSerializer(*input))
}

func QuestionUpdate(c *fiber.Ctx) error {
	input := new(Question)
	c.BodyParser(input)
	if err := validator.New().Struct(*input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id := c.Params("id")
	if input.ID == 0 || fmt.Sprint(input.ID) != id {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "URL ID and Body ID are not the same",
		})
	}

	db := common.GetDB()
	res := db.Omit(clause.Associations).Updates(&input)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	} else if res.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to find selected record",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}

func QuestionDelete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db := common.GetDB()
	res := db.Delete(&Question{}, id)
	if res.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": res.Error.Error(),
		})
	} else if res.RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unable to find selected record",
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
