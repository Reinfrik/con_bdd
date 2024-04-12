package conexion

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

func Conexion() {
	dsn := `user=SYSTEM
	        password=123456
			connectString=localhost:1521/XE`
	db, err := sql.Open("godror", dsn)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select count(*) from empleados")
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)

	defer rows.Close()
	var strVal string
	for rows.Next() {
		err := rows.Scan(&strVal)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strVal)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
