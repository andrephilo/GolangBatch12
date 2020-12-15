<html>
    <head>
        <title>
            Login
        </title>
    </head>
    <body>
        <form action="/login" method="post" id="login">
            Username: <input type ="text" name="username" id ="username"><br>
            Password: <input type ="text" name="password" id ="password"><br>
            Age: <input type ="number" name="age" id ="age"><br>
            Gender: <select name="gender" id="gender">
                <option value="laki-laki">Laki-laki</option>
                <option value="perempuan">Perempuan</option>
            </select>
            Email : <input type ="text" name="email" id ="email"><br>
            <input type="submit" value="Reg & Login">
        </form>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.5.1/jquery.min.js"> </script>
        <script>
            $("#login").on("submit", function (e) {
                e.preventDefault();

                var $self = $(this);
                var payload = JSON.stringify({
                    username: $('[name="username"]').val(),
                    password: $('[name="password"]').val(),
                    age: parseInt($('[name="age"]').val(), 10),
                    gender: $('[name="gender"]').val(),
                    email: $('[name="email"]').val(),
                });

                $.ajax({
                    url: $self.attr("action"),
                    type: $self.attr("method"),
                    data: payload,
                    contentType: 'application/json',
                }).then(function (res) {
                    var data = JSON.parse(res)
                    alert(data.message)
                    //$(".message").text(res);
                }).catch(function (a) {
                    alert("ERROR: " + a.responseText);
                });
            });
        </script>
    </body>
</html>