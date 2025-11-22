package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func dump(user, password, host, db string, port int) {
	ts := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("/dumps/dump.%s.sql", ts)

	f, err := os.Create(filename)
	if err != nil {
		log.Println("Error creating dump file:", err)
		return
	}
	defer f.Close()

	cmd := exec.Command(
		"pg_dump",
		"-h", host,
		"-p", strconv.Itoa(port),
		"-U", user,
		db,
	)

	// пишем дамп в файл
	cmd.Stdout = f
	cmd.Stderr = os.Stderr

	cmd.Env = append(os.Environ(),
		"PGPASSWORD="+password,
	)

	if err := cmd.Run(); err != nil {
		log.Println("Error dumping database:", err)
	} else {
		log.Println("Dump success:", filename)
	}
}

func main() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		host = "localhost"
	}
	portStr := os.Getenv("POSTGRES_PORT")
	if portStr == "" {
		portStr = "5432"
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalln("Invalid POSTGRES_PORT:", err)
	}

	if err := os.MkdirAll("/dumps", 0755); err != nil {
		log.Fatalln("Cannot create /dumps:", err)
	}

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		dump(user, password, host, db, port)
	}
}
