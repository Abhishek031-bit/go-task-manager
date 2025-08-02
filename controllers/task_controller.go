package controllers

import (
	"task-manager/database"
	"task-manager/models"
	"task-manager/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"err": "Invalid input"})
	}
	userID := c.Locals("user_id").(uint)
	task.UserID = userID
	if err := database.DB.Create(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"err": "Failed to create task"})
	}
	return c.JSON(task)
}

func GetTasks(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var tasks []models.Task
	if err := database.DB.Scopes(utils.Paginate(c)).Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"err": "Failed to get tasks"})
	}
	return c.JSON(tasks)
}

func GetTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	userID := c.Locals("user_id").(uint)
	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"err": "Task not found"})
	}
	return c.JSON(task)
}

func UpdateTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	userID := c.Locals("user_id").(uint)
	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"err": "Task not found"})
	}
	var input models.Task
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{"err": "Invalid input"})
	}
	task.Title = input.Title
	task.Description = input.Description
	task.Status = input.Status
	if err := database.DB.Save(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"err": "Failed to update task"})
	}
	return c.JSON(task)
}

func DeleteTask(c *fiber.Ctx) error {
	taskID := c.Params("id")
	userID := c.Locals("user_id").(uint)
	var task models.Task
	if err := database.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{"err": "Task not found"})
	}
	if err := database.DB.Delete(&task).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{"err": "Failed to delete task"})
	}
	return c.JSON(&fiber.Map{"message": "Task deleted successfully"})
}