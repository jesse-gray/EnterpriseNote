//==============================USER MANAGEMENT=============================

//Creates a new user via RESTful API
function signUp() {
    $('#signupForm').on('submit', function(event) {
        event.preventDefault();

        var firstname = $('#firstname').val()
        var lastname = $('#lastname').val()
        var signuppassword = $('#signUpPassword').val()
        var json = JSON.stringify({ "firstname": firstname, "lastname": lastname, "password": signuppassword })
        if ((firstname.length > 0) && (signuppassword.length > 0)) {
            $.ajax({
                type: 'POST',
                url: 'http://localhost:8000/api/users',
                dataType: 'json',
                data: json,
                contentType: 'application/json'
            });
            location.replace("/api/")
        }
    });
}

//Logs a user in
function login() {
    $('#loginForm').on('submit', function(event) {
        event.preventDefault();

        var userid = $('#userid').val()
        var password = $('#password').val()
        var json = JSON.stringify({ "userid": userid, "password": password })
        $.ajax({
            type: 'POST',
            url: 'http://localhost:8000/api/login',
            dataType: 'json',
            data: json,
            contentType: 'application/json',
            success: function(data) {
                if (data) {
                    location.replace("home")
                }
            }
        });
    });
}

//Logs a user out
function logout() {
    if (window.confirm("Are you sure you want to log out?")) {
        $.ajax({
            type: 'POST',
            url: 'http://localhost:8000/api/logout'
        });
        location.href = '/api/';
    }
}

//============================NOTES============================

//Loads notes from RESTful API
function loadNotes() {
    var url = "http://localhost:8000/api/notes";

    //Declare XMLHttpRequest Object
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var myArr = JSON.parse(this.responseText);
            if (myArr != null) {
                processNotes(myArr);
            }
        }
    };
}

//Formats data and displays it on the webpage
function processNotes(arr) {
    var yourOutput = "<h2>Your Notes</h2>";
    var shareOutput = "<h2>Notes that have been shared with you</h2>";
    if (arr[0] != null) {
        for (var i = 0; i < arr[0].length; i++) {
            //Display extracted info onto the webpage
            yourOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3>Note ID: ' + arr[0][i].noteid + '</h3><p class="card-text">Note Text: ' + arr[0][i].notetext + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="viewBtn" value="' + arr[0][i].noteid + '" onclick="viewNote(this.value)" type="button">View</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="updateBtn" value="' + arr[0][i].noteid + '" onclick="updateNote(this.value)" type="button">Update</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="deleteButton" value="' + arr[0][i].noteid + '" onclick="deleteNote(this.value)" type="button">Delete</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="updatePerms" value="' + arr[0][i].noteid + '" onclick="updatePermsPage(this.value)" type="button">Update Permissions</button></div></div></div></div></div>';
        }
    }
    if (arr[1] != null) {
        for (var i = 0; i < arr[1].length; i++) {
            shareOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3 class="card-title">Note ID: ' + arr[1][i].noteid + '</h3><p class="card-text">Note Text: ' + arr[1][i].notetext + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="viewBtn" value="' + arr[1][i].noteid + '" onclick="viewNote(this.value)" type="button">View</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="updateBtn" value="' + arr[1][i].noteid + '" onclick="updateNote(this.value)" type="button">Update</button></div></div></div></div></div>';
        }
    }
    if (arr[2] != null) {
        for (var i = 0; i < arr[2].length; i++) {
            shareOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3 class="card-title">Note ID: ' + arr[2][i].noteid + '</h3><p class="card-text">Note Text: ' + arr[2][i].notetext + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="viewBtn" value="' + arr[2][i].noteid + '" onclick="viewNote(this.value)" type="button">View</button></div></div></div></div></div>';
        }
    }
    document.getElementById("yourNotes").innerHTML = yourOutput;
    document.getElementById("sharedNotes").innerHTML = shareOutput;
}

//Loads one specific note to the page
function loadNote() {
    var url = "http://localhost:8000/api/note/" + sessionStorage.getItem("noteid");
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var note = JSON.parse(this.responseText);
            var output = '<div class="card text-white bg-secondary mb-3"><div class="card-body"><h3 class="card-title">Note ID: ' + note.noteid + '</h3><p class="card-text">Note Text: ' + note.notetext + '</p><p class="card-text">Author ID: ' + note.authorid + '</p></div></div>';
            document.getElementById("noteDetail").innerHTML = output;
        }
    }
}

//Create a new note via RESTful API
function createNote() {
    const serialize_form = form => JSON.stringify(
        Array.from(new FormData(form).entries())
        .reduce((m, [key, value]) => Object.assign(m, {
            [key]: value
        }), {})
    );

    $('#createForm').on('submit', function(event) {
        event.preventDefault();
        const json = serialize_form(this);

        if ($('#savedOption').is(':checked')) {
            x = "true"
        } else {
            x = "false"
        }
        if ($('#notetext').val() != "") {
            $.ajax({
                type: 'POST',
                url: 'http://localhost:8000/api/notes/' + x,
                dataType: 'json',
                data: json,
                contentType: 'application/json',
                success: function(data) {}
            });
            location.replace("home")
        }
    });
}

//Uses a SQL query to find all notes matching a text pattern
function findNote() {
    const serialize_form = form => JSON.stringify(
        Array.from(new FormData(form).entries())
        .reduce((m, [key, value]) => Object.assign(m, {
            [key]: value
        }), {})
    );

    $('#findNoteForm').on('submit', function(event) {
        event.preventDefault();
        const json = serialize_form(this);

        if ($('#searchterm').val() != "") {
            $.ajax({
                type: 'GET',
                url: 'http://localhost:8000/api/notes/' + $('#searchterm').val(),
                dataType: 'json',
                data: json,
                contentType: 'application/json',
                success: function(data) {
                    processNotes(data)
                }
            });
        }
    });
}

//Counts number of times a text pattern appears in a given note
function analyseNote() {
    const serialize_form = form => JSON.stringify(
        Array.from(new FormData(form).entries())
        .reduce((m, [key, value]) => Object.assign(m, {
            [key]: value
        }), {})
    );

    $('#findNoteForm').on('submit', function(event) {
        event.preventDefault();
        const json = serialize_form(this);

        if ($('#searchterm').val() != "") {
            $.ajax({
                type: 'GET',
                url: 'http://localhost:8000/api/notes/' + $('#formnoteid').val() + "/" + $('#searchterm').val(),
                dataType: 'json',
                data: json,
                contentType: 'application/json',
                success: function(data) {
                    document.getElementById("notesOutput").innerHTML = '<h2>Results:</h2><p class="card-text">Note ' + $('#formnoteid').val() + ' has ' + data + ' occuurances of "' + $('#searchterm').val() + '"</p>';
                },
                error: function(data) {
                    document.getElementById("notesOutput").innerHTML = '<h2>Results:</h2><p class="card-text">Permission to this note has been denied, or note doesnt exist</p>';
                }
            });
        }
    });
}

//Deletes a note via RESTful API
function deleteNote(noteID) {
    if (window.confirm("Are you sure you want to delete note " + noteID + "?")) {
        $.ajax({
            type: 'DELETE',
            url: 'http://localhost:8000/api/notes/' + noteID
        });
        location.href = 'home';
    }
}

// //Updates a note via RESTful API
// function updateNote() {
//     const serialize_form = form => JSON.stringify(
//         Array.from(new FormData(form).entries())
//         .reduce((m, [key, value]) => Object.assign(m, {
//             [key]: value
//         }), {})
//     );

//     $('#updateNoteForm').on('submit', function(event) {
//         event.preventDefault();
//         const json = serialize_form(this);
//         $.ajax({
//             type: 'PUT',
//             url: 'http://localhost:8000/api/notes/' + $('#formnoteid').val(),
//             dataType: 'json',
//             data: json,
//             contentType: 'application/json',
//             success: function(data) {}
//         });
//         location.href = 'viewNotes';
//     });
// }


//============================USERS============================

//Load all users from RESTful API
function loadUsers() {
    var url = "http://localhost:8000/api/users";

    //Declare XMLHttpRequest Object
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var myArr = JSON.parse(this.responseText);
            if (myArr != null) {
                processUsers(myArr);
            }
        }
    };
}

//Formats data and displays it on the webpage
function processUsers(arr) {
    var profileOutput = "<h2>Your Profile:</h2>";
    var userOutput = "<h2>All Other Users:</h2>";
    for (var i = 0; i < arr.length; i++) {
        //Display extracted info onto the webpage
        if (arr[i].cookieid == document.cookie.substring(document.cookie.indexOf("=") + 1)) {
            profileOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3 class="card-title">User ID: ' + arr[i].userid + '</h3><p class="card-text">Name: ' + arr[i].firstname + ' ' + arr[i].lastname + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="updateBtn" value="' + arr[i].userid + '" onclick="location.href=\'updateUser\'" type="button">Update</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="deleteButton" onclick="deleteUser()" type="button">Delete</button></div></div></div></div>';
        } else {
            userOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><h5 class="card-title">User ID: ' + arr[i].userid + '</h5><p class="card-text">Name: ' + arr[i].firstname + ' ' + arr[i].lastname + '</p></div></div></div>';
        }
    }
    document.getElementById("userProfile").innerHTML = profileOutput;
    document.getElementById("userList").innerHTML = userOutput;
}

//Deletes user via RESTful API
function deleteUser() {
    if (window.confirm("Are you sure you want to delete this user?")) {
        $.ajax({
            type: 'DELETE',
            url: 'http://localhost:8000/api/users'
        });
        sessionStorage.setItem("currentlyloggedin", 0);
        location.replace('/api/');
    }
}

//Updates a user via RESTful API
function updateUser() {
    const serialize_form = form => JSON.stringify(
        Array.from(new FormData(form).entries())
        .reduce((m, [key, value]) => Object.assign(m, {
            [key]: value
        }), {})
    );

    $('#updateUserForm').on('submit', function(event) {
        event.preventDefault();
        const json = serialize_form(this);
        $.ajax({
            type: 'PUT',
            url: 'http://localhost:8000/api/users',
            dataType: 'json',
            data: json,
            contentType: 'application/json',
            success: function(data) {
                alert(data)
            }
        });
        location.href = 'viewUsers';
    });
}


//=========================FAVOURITES==========================

//Loads a users favourites list from RESTful API
function loadFavourites() {
    var url = "http://localhost:8000/api/favourite";

    //Declare XMLHttpRequest Object
    var xmlhttp = new XMLHttpRequest();
    xmlhttp.open("GET", url, true);
    xmlhttp.send();
    xmlhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            var myArr = JSON.parse(this.responseText);
            if (myArr != null) {
                processFavourites(myArr);
            }
        }
    };
}

//Formats data and displays it on the webpage
function processFavourites(arr) {
    var output = "<h2>Your Favourites:</h2>";
    for (var i = 0; i < arr.length; i++) {
        //Display extracted info into the divs
        output += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3 class="card-title">User ID: ' + arr[i].userid + '</h3><p class="card-text">Name: ' + arr[i].firstname + ' ' + arr[i].lastname + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="removeButton" value="' + arr[i].userid + '" onclick="removeFave(this.value)" type="button">Remove</button></div></div></div></div></div>';
    }
    document.getElementById("userList").innerHTML = output;
}

//Creates a new favourite via RESTful API
function createFavourite() {
    $('#addFaveForm').on('submit', function(event) {
        event.preventDefault();

        var userid = $('#userid').val()
        var readpermission = $('input[name="readPerm"]:checked').val()
        var writepermission = $('input[name="writePerm"]:checked').val()

        var json = JSON.stringify({ "userid": parseInt(userid), "readpermission": (readpermission === "true"), "writepermission": (writepermission === "true") })
        console.log(json)
        $.ajax({
            type: 'POST',
            url: 'http://localhost:8000/api/favourite',
            dataType: 'json',
            data: json,
            contentType: 'application/json',
            success: function(data) {
                alert(data)
            }
        });
        location.href = 'viewFavourites';
    });
}

//Deletes favourite via RESTful API
function removeFave(userID) {
    if (window.confirm("Are you sure you want to remove UserID: " + userID + " from favourites?")) {
        $.ajax({
            type: 'DELETE',
            url: 'http://localhost:8000/api/favourite/' + userID
        });
        location.href = 'viewFavourites';
    }
}


//=========================PERMISSIONS==========================

//Updates a notes permissions via RESTful API
function updatePerms(noteID) {
    $('#updateForm').on('submit', function(event) {
        event.preventDefault();

        var noteid = $('#formnoteid').val()
        var userid = $('#userid').val()
        var readpermission = $('input[name="readPerm"]:checked').val()
        var writepermission = $('input[name="writePerm"]:checked').val()

        var json = JSON.stringify({ "noteid": parseInt(noteid), "userid": parseInt(userid), "readpermission": (readpermission === "true"), "writepermission": (writepermission === "true") })
        $.ajax({
            type: 'PUT',
            url: 'http://localhost:8000/api/permission',
            dataType: 'json',
            data: json,
            contentType: 'application/json',
            success: function(data) {
                alert(data)
            }
        });
        location.href = 'viewNotes';
    });
}


//=========================LOAD PAGES==========================

//Redirects from index page if already logged in
// function splashPageLoad() {
//     if (document.cookie != undefined) {
//         location.replace("/api/home")
//     }
// }

//Sets up updateNote page
function loadUpdateNotePage() {
    document.getElementById("formnoteid").value = sessionStorage.getItem("noteid");
}

//Sets up updateUser page
function loadUpdateUserPage() {
    document.getElementById("formuserid").value = sessionStorage.getItem("userid");
}

//Sets up updatePerms page
function loadUpdatePermissionPage() {
    document.getElementById("formnoteid").value = sessionStorage.getItem("noteid");
}

//==========================REDIRECTS==========================

//Sets up viewNote page
function viewNote(noteID) {
    sessionStorage.setItem("noteid", noteID)
    location.href = 'viewNote';
}

//Sets up updateNote page
function updateNote(noteID) {
    sessionStorage.setItem("noteid", noteID)
    location.href = 'updateNote';
}

//Sets up updatePerms page
function updatePermsPage(noteID) {
    sessionStorage.setItem("noteid", noteID)
    location.href = 'updatePerms';
}