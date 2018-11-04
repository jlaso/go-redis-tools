var current = window.location.pathname.replace(new RegExp("^\/"), "");

$.get("/top-menu?current=" + current, function (data) {
    $("body").prepend(data);
});