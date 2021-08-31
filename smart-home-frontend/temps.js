function GetTemps() {
    $.get("/api/v1/sensors/temperatures", function (data) {
        data.forEach(SetTemp);
    });
}

function SetTemp(item, index) {
    console.log( item["id"] + item["raw_value"]);
    h2 = document.getElementById(item["id"]).getElementsByClassName("temp_text")[0].innerHTML = item["raw_value"] + " °C";
    console.log(h2);
}

function UpdateFooter() {
    var dt = new Date();
    document.getElementById('time').innerHTML = dt.toLocaleTimeString()
}

window.onload = function() {
    GetTemps();
    setInterval(GetTemps, 1000);
    setInterval(UpdateFooter, 1000)
};