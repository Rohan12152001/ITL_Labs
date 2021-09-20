
// Today's date
var today = new Date();
var dd = String(today.getDate()).padStart(2, '0');
var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
var yyyy = today.getFullYear();

today = dd + '/' + mm + '/' + yyyy;

// Here
try {
    function display_table() {
        contest_table.innerHTML = " "     // Emptying the html before inserting

        fetch("https://corona.lmao.ninja/v2/states?sort&yesterday").then(response => response.json()).then(data => {
            contest_table.setAttribute('border', '2');

            // Modifying the table using DOM
            contest_table.innerHTML = "<tr>" +
                "<th>State</th>" +
                "<th>CasesToday</th>" +
                "<th>ActiveCases</th>" +
                "<th>Recovered</th>" +
                "<th>DeathsToday</th>" +
                "<th>TotalTests</th>" +
                "</tr>"
            console.log(data)

            // display records where contest is running/upcoming
            for (let record = 0; record < data.length; record++) {
                // console.log('1')
                contest_table.innerHTML += "<tr>" +
                    "<td>" + data[record]['state'] + "</td>" +
                    "<td>" + data[record]['todayCases'] + "</td>" +
                    "<td>" + data[record]['active'] + "</td>" +
                    "<td>" + data[record]['recovered'] + "</td>" +
                    "<td>" + data[record]['todayDeaths'] + "</td>" +
                    "<td>" + data[record]['tests'] + "</td>" +
                    "</form></td>" +
                    "</tr>"
            }
        })
    }

    display_table();
}
catch(err){
    console.log("Error in fetching data!")
}
