package main

import (
    "log"
)

func main() {
    cfg := config.LoadConfig()

    // MariaDB 연결 설정
    mariaDB, err := database.NewMariaDB(cfg.MariaDBURL)
    if err != nil {
        log.Fatalf("Failed to connect to MariaDB: %v", err)
    }
    defer mariaDB.Close()

    // S3 클라이언트 설정
    s3Client, err := storage.NewS3Client(cfg.S3Bucket, cfg.S3Region)
    if err != nil {
        log.Fatalf("Failed to connect to S3: %v", err)
    }

    // 서비스 및 핸들러 설정
    videoService := services.NewVideoService(s3Client, mariaDB)
    videoHandler := rest.NewVideoHandler(videoService)

    // HTTP 서버 설정
    http.HandleFunc("/upload", videoHandler.UploadVideoHandler)
    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
