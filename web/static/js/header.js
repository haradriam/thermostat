$(document).ready(function(){
    $.ajax({
        type: "GET",
        url: "/getinfo",
        //contentType: "application/json; charset=utf-8",
        dataType: "json",
        success: function (data, status, jqXHR) {
            $.each(data, function(index, value) {
                if(index == "Env") {
                    $("#header_table_env_temp").text("Temp: " + value.Temp + "º")
                    $("#header_table_env_hum").text("Hum: " + value.Hum + "%")
                } 
                if (index == "Heating") {
                    if(value == "true") {
                        $("#header_table_env_status").css('background-color', 'green');
                    } else {
                        $("#header_table_env_status").css('background-color', 'red');
                    }
                }
            });
        },
        error: function (jqXHR, status) {
            // error handler
        }
    });
});

