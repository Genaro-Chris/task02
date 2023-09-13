package main_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
	"task02/config"
	"task02/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var DB = config.OpenConnectionDb(&model.Person{}, "test.db")

func Test_Add(t *testing.T) {
	sum := 2 * 2
	if sum != 4 {
		t.Errorf("Expected %d, found %d", 4, sum)
	}
}

func Test_GetAll(t *testing.T) {
	app := fiber.New()
	app.Mount("/", NewRouter())
	req := httptest.NewRequest("GET", "/api", nil)
	resp, err := app.Test(req)
	if err != nil {
		panic(err)
	}
	expectedResults := []model.Person{}
	results := []model.Person{}
	byte_arr, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(byte_arr[:], &results)
	assert.Equalf(t, expectedResults, results, "Not equal")
}

func Test_Post(t *testing.T) {
	app := fiber.New()
	app.Mount("/", NewRouter())
	resp_body := model.Person{Name: "Post_Admin"}
	json_bytes, _ := json.Marshal(resp_body)
	json_string := string(json_bytes)
	reader_impl := strings.NewReader(json_string)
	req := httptest.NewRequest("POST", "/api", reader_impl)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	expectedResults := fmt.Sprintf("{name:%s}", resp_body.Name)
	result := model.Person{}
	byte_arr, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byte_arr, &result)
	if err != nil {
		panic(err)
	}
	finalresult := fmt.Sprintf("{name:%s}", result.Name)
	assert.Equalf(t, expectedResults, finalresult, "Not equal")
}

func Test_Get_By_ID(t *testing.T) {
	app := fiber.New()
	app.Mount("/", NewRouter())
	req := httptest.NewRequest("GET", "/api/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	expectedResults := model.Person{}
	var result model.Person
	byte_arr, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byte_arr, &result)
	if err != nil {
		panic(err)
	}
	assert.NotEqualf(t, expectedResults.Name, result.Name, "Equal")
}

func Test_Put_By_ID(t *testing.T) {
	app := fiber.New()
	app.Mount("/", NewRouter())
	resp_body := model.Person{Name: "PUTAdmin"}
	json_bytes, _ := json.Marshal(resp_body)
	json_string := string(json_bytes)
	reader_impl := strings.NewReader(json_string)
	req := httptest.NewRequest("PUT", "/api/1", reader_impl)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	expectedResults := resp_body
	expectedResults.ID = 1
	var result model.Person
	byte_arr, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byte_arr, &result)
	if err != nil {
		panic(err)
	}
	assert.Equalf(t, expectedResults.ID, result.ID, "Not Equal")
	assert.NotEqualf(t, expectedResults.Name, result.Name, "Equal")
}

func Test_Patch_By_ID(t *testing.T) {
	app := fiber.New()
	app.Mount("/", NewRouter())
	resp_body := model.Person{Name: "PATCHAdmin"}
	json_bytes, _ := json.Marshal(resp_body)
	json_string := string(json_bytes)
	reader_impl := strings.NewReader(json_string)
	req := httptest.NewRequest("PATCH", "/api/1", reader_impl)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	expectedResults := resp_body
	expectedResults.ID = 1
	result := model.Person{}
	byte_arr, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byte_arr, &result)
	if err != nil {
		panic(err)
	}

	assert.NotEqualf(t, expectedResults.Name, result.Name, "Equal")
	assert.Equalf(t, expectedResults.ID, result.ID, "Not Equal")
}

func Test_Delete_By_ID(t *testing.T) {
	app := fiber.New()
	app.Mount("/", NewRouter())
	req := httptest.NewRequest("DELETE", "/api/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	expectedResults := model.Person{}
	result := model.Person{}
	byte_arr, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byte_arr, &result)
	if err != nil {
		panic(err)
	}
	fmt.Println(expectedResults.Name, result.Name)
	assert.NotEqualf(t, expectedResults.Name, result.Name, "Equal")
}

func NewRouter() *fiber.App {
	router := fiber.New()
	router.Route("api/:id", func(router fiber.Router) {
		router.Delete("", func(c *fiber.Ctx) error {
			id, err := c.ParamsInt("id")
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Can't parse the ID from the query"
				c.JSON(&new_err)
				return nil
			}
			var person model.Person
			ptr := DB.Find(&person, id)
			if ptr == nil {
				new_err.Message = "No row with found with that id"
				c.JSON(&new_err)
				return nil
			}
			person_old := person
			DB.Unscoped().Delete(&person, id)
			c.JSON(person_old)
			return nil
		})

		router.Put("", func(c *fiber.Ctx) error {
			id, err := c.ParamsInt("id")
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Can't parse the ID from the query"
				c.JSON(&new_err)
				return nil
			}
			new_person := model.Person{}
			err = c.BodyParser(&new_person)
			if err != nil {
				new_err.Message = "Failed to parse body"
				c.JSON(&new_err)
				return nil
			}
			person := model.Person{}
			ptr := DB.Find(&person, id)
			if ptr == nil {
				new_err.Message = "No row with found with that id"
				c.JSON(&new_err)
				return nil
			}
			old_person := person
			person.Name = new_person.Name
			DB.Save(&person)
			c.JSON(old_person)
			return nil
		})

		router.Patch("", func(c *fiber.Ctx) error {
			id, err := c.ParamsInt("id")
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Can't parse the ID from the query"
				c.JSON(&new_err)
				return nil
			}
			new_person := model.Person{}
			err = c.BodyParser(&new_person)
			if err != nil {
				new_err.Message = "Failed to parse body"
				c.JSON(&new_err)
				return nil
			}
			person := model.Person{}
			ptr := DB.Find(&person, id)
			if ptr == nil {
				new_err.Message = "No row with found with that id"
				c.JSON(&new_err)
				return nil
			}
			old_person := person
			person.Name = new_person.Name
			DB.Save(&person)
			c.JSON(old_person)
			return nil
		})

		router.Get("", func(c *fiber.Ctx) error {
			id, err := c.ParamsInt("id")
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Can't parse the ID from the query"
				c.JSON(&new_err)
				return nil
			}
			var person = model.Person{}
			ptr := DB.Find(&person, id)
			if ptr == nil {
				new_err.Message = "No row with found with that id"
				c.JSON(&new_err)
				return nil
			}
			c.JSON(person)
			return nil
		})

	})

	router.Route("/api", func(router fiber.Router) {
		router.Get("", func(c *fiber.Ctx) error {
			var people = []model.Person{}
			DB.Find(&people)
			c.JSON(people)
			return nil
		})

		router.Post("", func(c *fiber.Ctx) error {
			person := model.Person{}
			err := c.BodyParser(&person)
			new_err := model.Error{}
			if err != nil {
				new_err.Message = "Failed to parse body"
				c.JSON(&new_err)
				return nil
			}
			DB.Create(&person)
			c.JSON(person)
			return nil
		})

	})

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "message",
			"message": "Alive",
		})
	})

	return router
}
