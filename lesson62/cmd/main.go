package main

import (
	"log"
	"students/api"
	"students/api/handler"
	"students/storage/redis"

	"github.com/casbin/casbin/v2"
)

func main() {
	// Casbin enforcer yaratish va konfiguratsiya fayllarini yuklash
	enforcer, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		log.Fatalf("Casbin enforcer yaratishda xatolik: %v", err)
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		log.Fatalf("Policy yuklashda xatolik: %v", err)
	}

	// Redis bilan ulanish
	db := redis.ConnectRDB()
	if err != nil {
		log.Fatalf("Redis bilan ulanishda xatolik: %v", err)
	}

	// Handler yarating
	h := handler.NewHandler(db, enforcer)

	// Routerni sozlang va handlerni o'rnating
	router := api.NewRouter(h)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Serverni ishga tushirishda xatolik: %v", err)
	}
}
