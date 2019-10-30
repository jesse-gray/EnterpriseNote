package main

//Note Struct (Model)
type Note struct {
	NoteID   string `json:"noteid"`
	NoteText string `json:"notetext"`
	AuthorID int    `json:"authorid"`
}

//User Struct
type User struct {
	UserID    string `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Permission Struct
type Permission struct {
	NoteID          int  `json:"noteid"`
	UserID          int  `json:"userid"`
	ReadPermission  bool `json:"readpermission"`
	WritePermission bool `json:"writepermission"`
}
