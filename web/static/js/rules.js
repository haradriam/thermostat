$(document).ready(function(){
    readEvents();
});

function readEvents() {
    $.ajax({
        type: "GET",
        url: "/getevents",
        dataType: "json",
        success: function (data, status, jqXHR) {
            $.each(data, function(index, value) {
                var resStart = data[index].StartTime.split(":");
                var resEnd = data[index].EndTime.split(":");
                var StartTime = resStart[0] + ":" + resStart[1];
                var EndTime = resEnd[0] + ":" + resEnd[1];

                var newRow = "<tr id=\"row" + index + "\"><td>" +
                            data[index].StartTemp +
                            "</td><td>" +
                            data[index].Periodic +
                            "</td><td>" +
                            StartTime +
                            "</td><td>" +
                            EndTime +
                            "</td><td style=\"display:none;\">" +
                            data[index].Active +
                            "</td><td>";
                if(data[index].Active == 1) {
                newRow = newRow + "<a><img src=/static/images/enabled.png class=\"event_table_img\" id=\"enable" + index + "\"></a>"
                }
                else {
                    newRow = newRow + "<img src=/static/images/disabled.png class=\"event_table_img\" id=\"enable" + index + "\">"
                }

                newRow = newRow + "</td><td>" +
                                "<img src=/static/images/edit.png class=\"event_table_img\" id=\"edit" + index + "\">" +
                                "</td><td>" +
                                "<img src=/static/images/delete.png class=\"event_table_img\" id=\"delete" + index + "\">" +
                                "</td></tr>"

                $("#event_table tr:last").after(newRow);
            });

            $(".event_table_img").click(function () {
                if (this.id.substring(0, this.id.length - 1) == "enable") {click_enable(parseInt(this.id[this.id.length - 1]) + 1);}
                if (this.id.substring(0, this.id.length - 1) == "edit") {click_edit(parseInt(this.id[this.id.length - 1]));}
                if (this.id.substring(0, this.id.length - 1) == "delete") {click_delete(this.id[this.id.length - 1]);}
            });
        },
        error: function (jqXHR, status) {
            // error handler
        }
    });
}

function click_enable(row) {
    if($("#event_table tr:eq(" + row +") td:eq(4)").text() == "1") {
        $("#event_table tr:eq(" + row +") td:eq(4)").text("0");
    }
    else {
        $("#event_table tr:eq(" + row +") td:eq(4)").text("1");
    }
    save_events();
    cleanTable();
    readEvents();
}

function click_edit(row) {
    var id = "#row" + row;
    $(id).attr("contentEditable", "true");
    $(id).attr("class", "event_table_editing");
    $("#event_table tr:eq(" + (row + 1)  +") td:eq(6)").html("<img src=/static/images/save.png class=\"event_table_img\" id=\"save\">");

    var size = $("#event_table tr").length;
    for(i = 0; i < size; i++) {
        id = "#row" + i;
        if(id != "#row" + row) {
            $(id).attr("class", "event_table_notediting");
        }
    }

    $("#save").click(function () {
        $("#event_table tr:eq(" + (row + 1)  +") td:eq(6)").html("<img src=/static/images/edit.png class=\"event_table_img\" id=\"edit" + row + "\">");
        for(i = 0; i < size; i++) {
            id = "#row" + i;
            $(id).attr("class", "none");
        }

        save_events();
    });
}

function click_delete(row) {
    var id = "#row" + row;
    $(id).remove();
    save_events();
}

function save_events() {
    var index = 0;
    var eventEntry = new Object();
    var eventList = "[";

    $("#event_table tr").each(function() {
        if(index != 0) {
            eventEntry.StartTemp = parseInt($("#event_table tr:eq(" + index +") td:eq(0)").text());
            eventEntry.Periodic = $("#event_table tr:eq(" + index +") td:eq(1)").text();
            eventEntry.StartTime = $("#event_table tr:eq(" + index +") td:eq(2)").text() + ":00";
            eventEntry.EndTime = $("#event_table tr:eq(" + index +") td:eq(3)").text() + ":00";
            eventEntry.Active = parseInt($("#event_table tr:eq(" + index +") td:eq(4)").text());

            eventList = eventList + JSON.stringify(eventEntry) + ",";
        }
        index++;
    });

    eventList = eventList.substring(0, eventList.length - 1) + "]";

    $.ajax({
        type: "POST",
        url: "/setevents",
        dataType: "json",
        data: eventList,
        success: function() {
            alert("HOLA");
        }
    });
}

function cleanTable() {
    var id;
    var size = $("#event_table tr").length;
    for(i = 0; i < size; i++) {
        id = "#row" + i;
        $(id).remove();
    }
}
