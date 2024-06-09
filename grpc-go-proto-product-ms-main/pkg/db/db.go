package db

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Handler struct {
    DB *gorm.DB
}

func Init(host, port, user, password, dbname string) Handler {
    // Crear la cadena de conexión
    connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
    if err != nil {
        log.Fatalln("Failed to connect to database:", err)
    }

    return Handler{DB: db}
}


