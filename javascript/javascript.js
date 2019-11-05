//Loads notes
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

//Finds relevant data in document and displays it on the webpage
function processNotes(arr) {
    var yourOutput = "<h2>Your Notes</h2>";
    var shareOutput = "<h2>Notes that have been shared with you</h2>";
    for (var i = 0; i < arr.length; i++) {
        //Display extracted info into the divs
        if (arr[i].cookieid == document.cookie.substring(document.cookie.indexOf("=") + 1)) {
            yourOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3>Note ID: ' + arr[i].noteid + '</h3><p class="card-text">Note Text: ' + arr[i].notetext + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="viewBtn" value="' + arr[i].noteid + '" onclick="viewNote(this.value)" type="button">View</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="updateBtn" value="' + arr[i].noteid + '" onclick="updateNote(this.value)" type="button">Update</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="deleteButton" value="' + arr[i].noteid + '" onclick="deleteNote(this.value)" type="button">Delete</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="updatePerms" value="' + arr[i].noteid + '" onclick="updatePerms(this.value)" type="button">Update Permissions</button></div></div></div></div></div>';
        } else {
            shareOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3 class="card-title">Note ID: ' + arr[i].noteid + '</h3><p class="card-text">Note Text: ' + arr[i].notetext + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="viewBtn" value="' + arr[i].noteid + '" onclick="viewNote(this.value)" type="button">View</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="updateBtn" value="' + arr[i].noteid + '" onclick="updateNote(this.value)" type="button">Update</button></div></div></div></div></div>';
        }
    }
    document.getElementById("yourNotes").innerHTML = yourOutput;
    document.getElementById("sharedNotes").innerHTML = shareOutput;
}

//Load users
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

//Finds relevant data in document and displays it on the webpage
function processUsers(arr) {
    var profileOutput = "<h2>Your Profile:</h2>";
    var userOutput = "<h2>All Other Users:</h2>";
    for (var i = 0; i < arr.length; i++) {
        //Display extracted info into the divs
        if (arr[i].cookieid == document.cookie.substring(document.cookie.indexOf("=") + 1)) {
            profileOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3 class="card-title">User ID: ' + arr[i].userid + '</h3><p class="card-text">Name: ' + arr[i].firstname + ' ' + arr[i].lastname + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="updateBtn" value="' + arr[i].userid + '" onclick="location.href=\'updateUser\'" type="button">Update</button><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="deleteButton" onclick="deleteUser()" type="button">Delete</button></div></div></div></div>';
        } else {
            userOutput += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><h5 class="card-title">User ID: ' + arr[i].userid + '</h5><p class="card-text">Name: ' + arr[i].firstname + ' ' + arr[i].lastname + '</p></div></div></div>';
        }
    }
    document.getElementById("userProfile").innerHTML = profileOutput;
    document.getElementById("userList").innerHTML = userOutput;
}

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

function processFavourites(arr) {
    var output = "<h2>Your Favourites:</h2>";
    for (var i = 0; i < arr.length; i++) {
        //Display extracted info into the divs
        output += '<div class="container-fluid"><div class="card text-white bg-secondary mb-3"><div class="card-body"><div class="row"><div class="col-sm-8"><h3 class="card-title">User ID: ' + arr[i].userid + '</h3><p class="card-text">Name: ' + arr[i].firstname + ' ' + arr[i].lastname + '</p></div><div class="col-sm-4"><button class="btn btn-light mr-1 mx-auto d-block btn-block" id="removeButton" value="' + arr[i].userid + '" onclick="removeFave(this.value)" type="button">Remove</button></div></div></div></div></div>';
    }
    document.getElementById("userList").innerHTML = output;
}

function removeFave(userID) {
    if (window.confirm("Are you sure you want to remove UserID: " + userID + " from favourites?")) {
        $.ajax({
            type: 'DELETE',
            url: 'http://localhost:8000/api/favourite/' + userID
        });
        location.href = 'viewFavourites';
    }
}

//Page loads
// function splashPageLoad() {
//     if (document.cookie != undefined) {
//         //location.replace("/api/home")
//     }
// }

function loadUpdateNotePage() {
    document.getElementById("formnoteid").value = sessionStorage.getItem("noteid");
}

function loadUpdateUserPage() {
    document.getElementById("formuserid").value = sessionStorage.getItem("userid");
}

function loadUpdatePermissionPage() {
    document.getElementById("formnoteid").value = sessionStorage.getItem("noteid");
}

//Notes
function viewNote(noteID) {
    sessionStorage.setItem("noteid", noteID)
    location.href = 'viewNote';
}

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

function findNote() {
    //Create SQL statement

    //Send to API
    var url = "http://localhost:8000/api/note/" + sqlStatement;
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

function updateNote(noteID) {
    sessionStorage.setItem("noteid", noteID)
    location.href = 'updateNote';
}

function deleteNote(noteID) {
    if (window.confirm("Are you sure you want to delete note " + noteID + "?")) {
        $.ajax({
            type: 'DELETE',
            url: 'http://localhost:8000/api/notes/' + noteID
        });
        location.href = 'home';
    }
}

function updatePerms(noteID) {
    sessionStorage.setItem("noteid", noteID)
    location.href = 'updatePerms';
}

//User managmenent functions
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

function logout() {
    if (window.confirm("Are you sure you want to log out?")) {
        $.ajax({
            type: 'POST',
            url: 'http://localhost:8000/api/logout'
        });
        location.href = '/api/';
    }

}