// Today.js

// Today's date //
var today = new Date();
var dd = String(today.getDate()).padStart(2, '0');
var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
var yyyy = today.getFullYear();

today = dd + '/' + mm + '/' + yyyy;

// Get cookie  //
function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) return parts.pop().split(';').shift();
}

// loadTop && use cookie if exists //
function loadTop() {
    if (true){
        // Get cookies
        console.log("Name is:" + getCookie("Name"))
        nameFromCookie = getCookie("Name")
        if (nameFromCookie==undefined){
            nameFromCookie="#"
        }
    }
    document.getElementById('Top-div').innerHTML = `<h2><u>Date: ${today}</u>&nbsp;&nbsp;
                                                    <u>Country:USA</u>&nbsp;&nbsp;
                                                    <u>(Your name: ${nameFromCookie})</u></h2>`
}
loadTop()

try {
    // loadCases //
    function loadCases() {
        contest_table.innerHTML = " "     // Emptying the html before inserting

        fetch("https://corona.lmao.ninja/v2/countries/USA?yesterday=true&strict=true&query")
            .then(response => response.json())
            .then(data => {
            contest_table.setAttribute('border', '2');

            console.log(data)
            contest_table.innerHTML += "<tr>" +
                "<td>" + "<b>" + "Confirmed Cases" + "</b>" + "</td>" +
                "<td>" + data['todayCases'] + "</td>" +
                "</tr>" +
                "<tr>" +
                "<td>" + "<b>" + "Active" + "</b>" + "</td>" +
                "<td>" + data['active'] + "</td>" +
                "</tr>" +
                "<tr>" +
                "<td>" + "<b>" + "Recovered" + "</b>" + "</td>" +
                "<td>" + data['todayRecovered'] + "</td>" +
                "</tr>" +
                "<tr>" +
                "<td>" + "<b>" + "Deaths" + "</b>" + "</td>" +
                "<td>" + data['todayDeaths'] + "</td>" +
                "</tr>" +
                "<tr>" +
                "<td>" + "<b>" + "Total Tests" + "</b>" + "</td>" +
                "<td>" + data['tests'] + "</td>" +
                "</tr>"
        })
    }

    loadCases()
}
catch (err) {
    console.log("Error in fetching data!")
}