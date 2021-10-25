package quiz

import "github.com/gofiber/fiber/v2"

func QuizzesSerializer(quizzes []Quiz) []fiber.Map {
	response := []fiber.Map{}
	for _, quiz := range quizzes {
		response = append(response, fiber.Map{
			"id":         quiz.ID,
			"start_time": quiz.StartTime,
		})
	}
	return response
}

func QuizSerializer(quiz Quiz) fiber.Map {
	return fiber.Map{
		"id":         quiz.ID,
		"start_time": quiz.StartTime,
		"questions":  QuestionsSerializer(quiz.Questions),
		"created_at": quiz.CreatedAt,
		"updated_at": quiz.UpdatedAt,
	}
}

func QuestionsSerializer(questions []Question) []fiber.Map {
	response := []fiber.Map{}
	for _, question := range questions {
		response = append(response, fiber.Map{
			"id":   question.ID,
			"text": question.Text,
		})
	}
	return response
}

func QuestionSerializer(question Question) fiber.Map {
	return fiber.Map{
		"id":         question.ID,
		"quiz_id":    question.QuizID,
		"text":       question.Text,
		"created_at": question.CreatedAt,
		"updated_at": question.UpdatedAt,
	}
}
