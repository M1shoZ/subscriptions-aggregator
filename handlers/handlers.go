package handlers

import (
	"log"
	"net/url"

	"github.com/M1shoZ/subscriptions-aggregator/database"
	"github.com/M1shoZ/subscriptions-aggregator/models"
	"github.com/gofiber/fiber/v2"

	_ "github.com/M1shoZ/subscriptions-aggregator/docs"
)

// Home godoc
// @Summary      Health check
// @Description  Проверка доступности API
// @Tags         system
// @Success      200  {string}  string  "hiii M1shoZ!!!"
// @Router       /home [get]
func Home(c *fiber.Ctx) error {
	return c.SendString("hiii M1shoZ!!! ")
}

// GetAllUsers godoc
// @Summary      Получить всех пользователей
// @Description  Возвращает список всех пользователей
// @Tags         users
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /get_users [get]
func GetAllUsers(c *fiber.Ctx) error {
	log.Println("[GetAllUsers] Request")

	users := &[]models.User{}

	if err := database.DB.Db.Find(users).Error; err != nil {
		log.Printf("[GetAllUsers] Failed to find: %v", err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "couldnot get users from DB^^^",
		})
	}
	log.Println("[GetAllUsers] Users got succesfully")
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Users got succesfully",
		"users":   users,
	})
	return nil
}

// CreateSub godoc
// @Summary      Создать подписку
// @Description  Добавляет новую подписку для пользователя
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription  body  models.Subscriptions  true  "Данные подписки"
// @Success      200  {object}  models.Subscriptions
// @Failure      400  {object}  map[string]interface{}
// @Router       /create_sub [post]
func CreateSub(c *fiber.Ctx) error {
	log.Println("[CreateSub] Request")

	sub := new(models.Subscriptions)

	if err := c.BodyParser(sub); err != nil {
		log.Printf("[CreateSub] Failed to parse body: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Printf("[CreateSub] Parsed subscription: %+v", sub)

	// database.DB.Db.Create(&sub)
	if err := database.DB.Db.Create(&sub).Error; err != nil {
		log.Printf("[CreateSub] Failed to create subscription: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create subscription",
			"error":   err.Error(),
		})
	}

	log.Println("[CreateSub] Subscription created successfully")
	return c.Status(fiber.StatusOK).JSON(sub)
}

// GetAllSubs godoc
// @Summary      Получить все подписки
// @Description  Возвращает список всех подписок
// @Tags         subscriptions
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /get_subs [get]
func GetAllSubs(c *fiber.Ctx) error {
	log.Println("[GetAllSubs] Request")

	subs := &[]models.Subscriptions{}
	if err := database.DB.Db.Find(subs).Error; err != nil {
		log.Printf("[GetAllSubs] Failed to find: %v", err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "couldnot get users from DB^^^",
		})
	}
	log.Println("[GetAllSubs] subs got succesfully")
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "subs got succesfully",
		"subs":    subs,
	})
	return nil
}

// GetSubById godoc
// @Summary      Получить подписку по ID
// @Description  Возвращает подписку по ID
// @Tags         subscriptions
// @Produce      json
// @Param        id  path  int  true  "ID подписки"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /get_subs/{id} [get]
func GetSubById(c *fiber.Ctx) error {

	log.Println("[GetSubById] Request")

	id := c.Params("id")
	SubscriptionModel := &models.Subscriptions{}
	if id == "" {
		c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	err := database.DB.Db.Where("id = ?", id).First(SubscriptionModel).Error
	log.Printf("[GetSubById] id = %s", id)
	if err != nil {
		log.Printf("[GetSubById] could not get the Subscription: %v", err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get the Subscription",
		})
		return nil
	}
	log.Println("[GetSubById] Subscription got succesfully")
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Subscription got succesfully",
		"Subscription": SubscriptionModel,
	})
	return nil
}

// GetSubByUserId godoc
// @Summary      Получить подписки по ID пользователя
// @Description  Возвращает список подписок для заданного пользователя
// @Tags         subscriptions
// @Produce      json
// @Param        user_id  path  string  true  "UUID пользователя"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /get_subs_by_user/{user_id} [get]
func GetSubByUserId(c *fiber.Ctx) error {
	log.Println("[GetSubByUserId] Request")

	userId := c.Params("user_id")
	SubscriptionModel := &[]models.Subscriptions{}
	if userId == "" {
		c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	err := database.DB.Db.Where("user_id = ?", userId).Find(SubscriptionModel).Error
	log.Printf("[GetSubByUserId] user_id = %s", userId)
	if err != nil {
		log.Printf("[GetSubByUserId] could not get the Subscription: %v", err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get the Subscription",
		})
		return nil
	}

	log.Println("[GetSubByUserId] Subscriptions got succesfully")
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Subscriptions got succesfully",
		"Subscription": SubscriptionModel,
	})
	return nil
}

// DeleteSub godoc
// @Summary      Удалить подписку
// @Description  Удаляет подписку по ID
// @Tags         subscriptions
// @Produce      json
// @Param        id  path  int  true  "ID подписки"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /delete/{id} [delete]
func DeleteSub(c *fiber.Ctx) error {
	log.Println("[DeleteSub] Request")

	subscription := &models.Subscriptions{}
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}
	err := database.DB.Db.Delete(subscription, id)
	if err.Error != nil {
		log.Printf("[DeleteSub] could not delete the subscription: %v", err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not delete the subscription",
		})
		return err.Error
	}

	log.Println("[DeleteSub] subscription deleted succesfully")
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "subscription deleted succesfully",
	})
	return nil
}

// UpdateSub godoc
// @Summary      Обновить подписку
// @Description  Обновляет данные подписки по ID
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id            path  int                  true  "ID подписки"
// @Param        subscription  body  models.Subscriptions  true  "Новые данные"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}   map[string]interface{}
// @Router       /update/{id} [patch]
func UpdateSub(c *fiber.Ctx) error {
	log.Println("[UpdateSub] Request")

	subscription := new(models.Subscriptions)
	id := c.Params("id")

	if id == "" {
		c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	if err := c.BodyParser(subscription); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	err := database.DB.Db.Model(&subscription).Where("id = ?", id).Updates(subscription)
	if err.Error != nil {
		log.Printf("[UpdateSub] could not update the subscription: %v", err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not update the subscription",
		})
		return err.Error
	}

	log.Println("[UpdateSub] Subscription updated succesfully!")
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Subscription updated succesfully!",
		"date":    subscription,
	})
	return nil
}

// GetSum godoc
// @Summary      Получить сумму подписок
// @Description  Возвращает сумму цен подписок по пользователю, сервису и периоду
// @Tags         subscriptions
// @Produce      json
// @Param        user_id       path  string  true  "UUID пользователя"
// @Param        service_name  path  string  true  "Название сервиса"
// @Param        start_date    path  string  true  "Дата начала (YYYY-MM-DD)"
// @Param        end_date      path  string  true  "Дата окончания (YYYY-MM-DD)"
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {object}  map[string]interface{}
// @Router       /get_sum/{user_id}/{service_name}/{start_date}/{end_date} [get]
func GetSum(c *fiber.Ctx) error {
	log.Println("[GetSum] Request")

	UserID := c.Params("user_id")
	// ServiceName := c.Params("service_name")
	StartDate := c.Params("start_date")
	EndDate := c.Params("end_date")

	serviceNameRaw := c.Params("service_name")
	ServiceName, err := url.QueryUnescape(serviceNameRaw)
	if err != nil {
		log.Printf("Error decoding service name: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid service name",
		})
	}

	SubscriptionModel := &models.Subscriptions{}
	var totalPrice int

	if UserID == "" || ServiceName == "" || StartDate == "" || EndDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required parameters",
		})
	}

	// fmt.Println("GOT:", UserID, ServiceName, StartDate, EndDate)

	err = database.DB.Db.Model(SubscriptionModel).
		Select("COALESCE(SUM(price), 0)").
		Where("user_id = ? AND service_name = ? AND start_date BETWEEN ? AND ?", UserID, ServiceName, StartDate, EndDate).
		Scan(&totalPrice).Error

	if err != nil {
		log.Printf("[UpdateSub] could not get the Sum: %v", err)
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "could not get the Sum",
			"error:":  err,
		})
		return err
	}

	log.Println("[UpdateSub] Sum got succesfully!")
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Sum got succesfully",
		"Subscription": totalPrice,
	})
	return nil
}
