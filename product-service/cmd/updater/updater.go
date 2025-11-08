package main

import (
	"database/sql"
	"fmt"
	"github.com/rs/xid"
	"log"
	"productService/config"
	"productService/internal/pkg/psql"
)

func main() {
	cfg := config.MustSetup()
	db := psql.MustConnect(cfg)
	defer db.Disconnect()

	tx, err := db.Driver.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else {
			_ = tx.Commit()
		}
	}()

	// отображение старых -> новых id
	idMap := make(map[string]string)

	// 1. Сначала генерим новые id для всех простых таблиц (без апдейтов)
	genIDMap(tx, "brands", idMap)
	genIDMap(tx, "categories", idMap)
	genIDMap(tx, "countries", idMap)
	genIDMap(tx, "materials", idMap)
	genIDMap(tx, "colors", idMap)
	genIDMap(tx, "products", idMap)

	// 2. Обновляем ссылки в products (FK)
	// сначала обновляем ID во всех "простых" таблицах
	updateIDs(tx, "brands", idMap)
	updateIDs(tx, "categories", idMap)
	updateIDs(tx, "countries", idMap)
	updateIDs(tx, "materials", idMap)
	updateIDs(tx, "colors", idMap)
	updateIDs(tx, "products", idMap)

	// потом обновляем ссылки
	updateFK(tx, "products", "brand_id", idMap)
	updateFK(tx, "products", "category_id", idMap)
	updateFK(tx, "products", "country_id", idMap)

	updateRelationTable(tx, "product_materials", "product_id", "material_id", idMap)
	updateRelationTable(tx, "product_colors", "product_id", "color_id", idMap)
	updateRelationTable(tx, "product_seems", "product_id", "similar_product_id", idMap)
	updateRelationTable(tx, "product_color_photos", "product_id", "color_id", idMap)
	log.Println("✅ Все ID и ссылки успешно заменены")
}

// генерим idMap без апдейтов
func genIDMap(tx *sql.Tx, table string, idMap map[string]string) {
	rows, err := tx.Query(fmt.Sprintf(`SELECT id FROM %s`, table))
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Panic(err)
		}
		if _, ok := idMap[id]; !ok {
			idMap[id] = xid.New().String()
		}
	}
}

// обновляем ссылки на новые id
func updateFK(tx *sql.Tx, table, col string, idMap map[string]string) {
	rows, err := tx.Query(fmt.Sprintf(`SELECT id, %s FROM %s`, col, table))
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	type rec struct {
		id string
		fk string
	}
	var list []rec
	for rows.Next() {
		var r rec
		if err := rows.Scan(&r.id, &r.fk); err != nil {
			log.Panic(err)
		}
		list = append(list, r)
	}

	for _, r := range list {
		if newFK, ok := idMap[r.fk]; ok {
			_, err = tx.Exec(fmt.Sprintf(
				`UPDATE %s SET %s=$1 WHERE id=$2`,
				table, col,
			), newFK, r.id)
			if err != nil {
				log.Panic(err)
			}
			log.Printf("[%s.%s] %s → %s\n", table, col, r.fk, newFK)
		}
	}
}

// обновляем id на новые
func updateIDs(tx *sql.Tx, table string, idMap map[string]string) {
	rows, err := tx.Query(fmt.Sprintf(`SELECT id FROM %s`, table))
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			log.Panic(err)
		}
		ids = append(ids, id)
	}

	for _, id := range ids {
		newID := idMap[id]
		if newID != "" && newID != id {
			_, err = tx.Exec(fmt.Sprintf(`UPDATE %s SET id=$1 WHERE id=$2`, table), newID, id)
			if err != nil {
				log.Panic(err)
			}
			log.Printf("[%s] %s → %s\n", table, id, newID)
		}
	}
}

// обновляем связующие таблицы
func updateRelationTable(tx *sql.Tx, table, col1, col2 string, idMap map[string]string) {
	rows, err := tx.Query(fmt.Sprintf(`SELECT %s, %s FROM %s`, col1, col2, table))
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	type pair struct {
		v1, v2 string
	}
	var pairs []pair
	for rows.Next() {
		var v1, v2 string
		if err := rows.Scan(&v1, &v2); err != nil {
			log.Panic(err)
		}
		pairs = append(pairs, pair{v1, v2})
	}

	for _, p := range pairs {
		newV1 := idMap[p.v1]
		if newV1 == "" {
			newV1 = p.v1
		}
		newV2 := idMap[p.v2]
		if newV2 == "" {
			newV2 = p.v2
		}

		_, err = tx.Exec(fmt.Sprintf(
			`UPDATE %s SET %s=$1, %s=$2 WHERE %s=$3 AND %s=$4`,
			table, col1, col2, col1, col2,
		), newV1, newV2, p.v1, p.v2)
		if err != nil {
			log.Panic(err)
		}
		log.Printf("[%s] (%s,%s) → (%s,%s)\n", table, p.v1, p.v2, newV1, newV2)
	}
}
