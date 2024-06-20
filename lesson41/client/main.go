package main

import (
	"bytes"
	"client/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	GetEmployee(client)
	CreateEmployee(client)
	UpdateEmployee(client)
	DeleteEmployee(client)
}

func GetEmployee(client *http.Client) {
	req, err := http.NewRequest("GET", "http://localhost:8080/employees/21", nil)
	if err != nil {
		log.Fatal("GET so'rovini yuborishda xatolik: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("GET so'rovini yuborishda xatolik: ", err)
		return
	}
	defer resp.Body.Close()

	var emp models.Employee
	err = json.NewDecoder(resp.Body).Decode(&emp)
	if err != nil {
		log.Fatal("HTTP javobini o'qishda xatolik: ", err)
	}

	fmt.Println(emp)
}


func CreateEmployee(client *http.Client) {
	employee := models.Employee{
		ID:       21,
		Name:     "Eldor Shomurodov",
		Position: "Software Developer",
		Salary:   58000,
	}
	jsonEmp, err := json.Marshal(employee)
	if err != nil {
		log.Fatal("Jsonni enkod qilishda xatolik yuz berdi: ", err)
		return
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/employees", bytes.NewBuffer(jsonEmp))
	if err != nil {
		log.Fatal("POST so'rovini yuborishda xatolik: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("POST so'rovini yuborishda xatolik: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("POST so'rovi javobi: ", resp.Status)
}

func UpdateEmployee(client *http.Client) {
	employee := models.Employee{
		ID:       21,
		Name:     "Abduqodir Husanov",
		Position: "Senior Software Developer",
		Salary:   75000,
	}
	jsonEmp, err := json.Marshal(employee)
	if err != nil {
		log.Fatal("Jsonni encode qilishda xatolik yuz berdi: ", err)
		return
	}

	req, err := http.NewRequest("PUT", "http://localhost:8080/employees/21", bytes.NewBuffer(jsonEmp))
	if err != nil {
		log.Fatal("PUT so'rovini yuborishda xatolik: ", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("PUT so'rovini yuborishda xatolik: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("PUT so'rovi javobi: ", resp.Status)
}

func DeleteEmployee(client *http.Client) {
	req, err := http.NewRequest("DELETE", "http://localhost:8080/employees/21", nil)
	if err != nil {
		log.Fatal("DELETE so'rovini yuborishda xatolik: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("DELETE so'rovini yuborishda xatolik: ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("DELETE so'rovi javobi: ", resp.Status)
}
