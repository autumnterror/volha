package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"time"
)

const maxDumps = 20

var dumpRegexp = regexp.MustCompile(`^dump\.\d{8}_\d{6}\.sql$`)

func cleanupDumps(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Println("Error reading dumps dir:", err)
		return
	}
	var dumpFiles []string
	for _, f := range files {
		if !f.IsDir() && dumpRegexp.MatchString(f.Name()) {
			dumpFiles = append(dumpFiles, f.Name())
		}
	}

	if len(dumpFiles) <= maxDumps {
		return // нечего чистить
	}

	sort.Strings(dumpFiles)

	toDelete := len(dumpFiles) - maxDumps

	log.Printf("Too many dumps (%d). Deleting %d oldest...\n", len(dumpFiles), toDelete)

	for i := 0; i < toDelete; i++ {
		path := filepath.Join(dir, dumpFiles[i])
		if err := os.Remove(path); err != nil {
			log.Println("Error deleting", path, ":", err)
		} else {
			log.Println("Deleted old dump:", dumpFiles[i])
		}
	}
}

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

	cmd.Stdout = f
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)

	if err := cmd.Run(); err != nil {
		log.Println("Error dumping database:", err)
	} else {
		log.Println("Dump success:", filename)
	}

	// очищаем старые дампы
	cleanupDumps("/dumps")
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

	//dump(user, password, host, db, port)
	log.Printf("next dump in %s", time.Now().Local().Add(time.Hour).Format("20060102_150405"))
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	for range ticker.C {
		dump(user, password, host, db, port)
		log.Printf("next dump in %s", time.Now().Local().Add(time.Hour).Format("20060102_150405"))
	}
}
