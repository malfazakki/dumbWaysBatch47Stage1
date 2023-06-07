package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
)

type Project struct {
	Title string
	Content string
	StartDate string
	EndDate string
	Duration string
}

var dataProject = []Project {
	{
		Title: "Dumbways Web Apps 2023",
		Content: "Content",
		StartDate: "2023/07/01",
		EndDate: "2023/07/06",  
		Duration: "3 months",
	},
	{
		Title: "Dumbways Web Apps 2023",
		Content: "Content 2",
		StartDate: "2023/07/01",
		EndDate: "2023/07/06",
		Duration: "3 months",
	},
}
 
func main() {
	e := echo.New()

	e.Static("/public", "public")

	// routing
	//get
	e.GET("/", home)
	e.GET("/contact", contact)
	e.GET("/add-project", addProject)
	e.GET("/testimonial", testimonial)
	e.GET("/project-detail/:id", projectDetail)
	//post
	e.POST("/add-project", formAddProject)
	e.POST("/project-delete/:id", deleteProject)
	e.Logger.Fatal(e.Start("localhost:5000"))
}

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	projects := map[string]interface{} {
		"Projects": dataProject,
	}

	return tmpl.Execute(c.Response(), projects)
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

	var ProjectDetail = Project{}

	for i, data := range dataProject {
		if id == i {
			ProjectDetail = Project{
				Title: data.Title,
				Content: data.Content,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
			}
		}
	}

	data := map[string]interface{} {
		"Project": ProjectDetail,
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
	duration := calculateDuration(startDate, endDate)

	fmt.Println("Title :", title)
	fmt.Println("Start Date :", startDate)
	fmt.Println("End Date :", endDate)
	fmt.Println("Content :", content)

	// Menampilkan durasi
	switch {
	case duration >= 365:
		duration /= 365
	case duration >= 30:
		duration /= 30
	case duration >= 7:
		duration /= 7
	default:
		duration /= 1
	}

	var newProject = Project{
		Title: title,
		Content: content,
		StartDate: startDate,
		EndDate: endDate,
		Duration: duration,
	}

	dataProject = append(dataProject, newProject)

	fmt.Println(dataProject)
	
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func calculateDuration(startDate, endDate string) int {
	// Mengubah string tanggal menjadi tipe time.Time
	startTime, _ := time.Parse("02-01-2006", startDate)
	endTime, _ := time.Parse("02-01-2006", endDate)

	// Menghitung durasi antara dua tanggal
	duration := int(endTime.Sub(startTime).Hours() / 24)

	return duration
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	fmt.Println("Index :", id)

	dataProject = append(dataProject[:id], dataProject[id+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/")
}