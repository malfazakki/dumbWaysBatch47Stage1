package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)
 
func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/add-project", addProject)
	e.GET("/testimonial", testimonial)
	e.GET("/project-detail", projectDetail)
	e.POST("/add-project", formAddProject)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func addProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/add-project.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id":      id,
		"Title":  "Dumbways Web App",
		"Content": "Lorem ipsum dolor sit amet consectetur adipisicing elit. Distinctio nemo repudiandae voluptas voluptatibus modi inventore totam quaerat itaque fugiat labore! Aliquid cumque nulla iusto eaque sequi impedit rerum harum magni minus vel? Officiis quod magnam minus asperiores repellendus, autem nemo quaerat aliquid, porro nesciunt ex mollitia. Veritatis architecto voluptatem earum amet dolor enim molestias, dicta qui magni similique vero! Quis obcaecati voluptas non eum amet, mollitia, ut commodi explicabo ad praesentium debitis nemo dicta voluptatum! Voluptatum odit a voluptas, quidem temporibus inventore! Iste repellat vitae autem! Ullam expedita atque odio dolorem laudantium tempora adipisci autem nulla iste at sequi eum eaque vero blanditiis, quis tempore molestias fugiat inventore exercitationem.",
	}
	var tmpl, err = template.ParseFiles("views/project-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return tmpl.Execute(c.Response(), data)
}

func formAddProject(c echo.Context) error {
	title := c.FormValue("input-project-name")
	startDate := c.FormValue("input-date-start")
	endDate := c.FormValue("input-date-end")
	content := c.FormValue("input-deskripsi")

	fmt.Println("Title :", title)
	fmt.Println("Start Date :", startDate)
	fmt.Println("End Date :", endDate)
	fmt.Println("Content :", content)

	

	return c.Redirect(http.StatusMovedPermanently, "/project-detail")
}