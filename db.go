package main

import (
	"database/sql"
	"log"
	"os"
	"time"
)

var (
	db *sql.DB
)

type Appointment struct {
	username  string
	date      time.Time
	timeOfDay string
}

const Morning = "Morning"
const Afternoon = "Afternoon"
const Evening = "Evening"

func timeIsValid(value string) bool {
	return value == Morning || value == Afternoon || value == Evening
}

func getUserAppointment(username string) Appointment {
	var appointment Appointment
	rows, err := db.Query("SELECT day, time FROM appointments WHERE username = $1", username)
	if err != nil {
		log.Panicf("Database error: %q", err)
	}
	defer rows.Close()
	for rows.Next() {
		appointment = Appointment{}
		appointment.username = username
		err := rows.Scan(&appointment.date, &appointment.timeOfDay)
		if err != nil {
			log.Panicf("Database error: %q", err)
		}
	}
	return appointment
}

func createUserAppointment(appointment Appointment) {
	deleteUserAppointment(appointment.username)
	_, err := db.Query("INSERT INTO appointments VALUES ($1, $2, $3)", appointment.username, appointment.date, appointment.timeOfDay)
	if err != nil {
		log.Panicf("Database error: %q", err)
	}
}

func deleteUserAppointment(username string) {
	_, err := db.Query("DELETE FROM appointments WHERE username = $1", username)
	if err != nil {
		log.Panicf("Database error: %q", err)
	}
}

func setupDB() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("GO_INSURANCE_TEST_PGURL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	_, err = db.Exec(`
		DO $$
			BEGIN
			  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'time_of_day') THEN
			    CREATE TYPE time_of_day AS ENUM ('Morning', 'Afternoon', 'Evening');
			  END IF;
				CREATE TABLE IF NOT EXISTS appointments (username TEXT UNIQUE, day date, time time_of_day);
			END
		$$`)
	if err != nil {
		log.Fatalf("Error setting up tables: %q", err)
	}
}
