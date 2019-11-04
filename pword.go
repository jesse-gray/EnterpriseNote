package main

// type Pword struct {
// 	PwordID 	int 	`json:"pwordID"`
// 	UserID 		int 	`json:"userID"`
// 	Pword		string	`json:"pword"`
// }

// func checkPassword(password string) bool {
// 	var pword string

// 	db := opendb()
// 	defer db.Close()

// 	// using sqlStatement to match use through code

// 	// sqlStatement := db.Prepare("SELECT pword FROM pasword WHERE userID = $1;")
// 	sqlStatement := db.Prepare("SELECT user_password FROM 'user' WHERE user_id = $1;")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = sqlStatement.QueryRow(password).Scan(&pword)

// 	// if err == null ie nothing is returned from query, there is no matching password so yjrm er report false
// 	if err == sql.ErrNoRows {
// 		return false
// 	}

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// we have a match =)

// 	return true
// }

// // function to add a new password to the table need to add if statement to check if password already exists for statement
// func createPassword(w http.ResponseWriter, r *http.Request) {

// 	var newPassword Pword
// 	_ = json.NewDecoder(r.Body).Decode(&Pword)
// 	db := opendb()
// 	defer db.Close()
// 	sqlStatement, err := db.Prepare("INSERT INTO pasword VALUES ($1,$2,$3);")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = sqlStatement.Exec(newPassword.PwordID, newPassword.UserID, newPassword.PWord)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return "Password successfully created for the account"
// }

// func updatePassword(w http.ResponseWriter, r *http.Request) string {
// 	params := mux.Vars(r)
// 	var newPassword Pword
// 	_ = json.NewDecoder(r.Body).Decode(&Pword)
// 	db := opendb()
// 	defer db.Close()
// 	sqlStatement, err := db.Prepare("UPDATE pasword SET newPassword = $1 WHERE newPassword.user_id")
// 	_, err = sqlStatement.Exec(newPassword.Pword, newPassword.PwordID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return "new password added"
// }

// adding in a log in function
