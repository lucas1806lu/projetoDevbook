$('#login').on('submit', fazerLogin);

   function fazerLogin(evento) {
    evento.preventDefault();
       console.log("funcionando o butoon");


    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function() {

        Swal.fire("Ops.....","usuário ou senha invalidos!", "error");
               
    });
}