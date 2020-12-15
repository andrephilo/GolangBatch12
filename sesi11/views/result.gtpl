<html>
<head>
<style>
table, th, td {
  border: 1px solid black;
}
</style>
</head>
<body>

<h2>Users</h2>

<table style="width:100%">
  <tr>
    <th>name</th>
    <th>age</th> 
    <th>gender</th>
    <th>email</th>
    <th>password</th>
  </tr>
  <tr>
    <td>{{.resname}}</td>
    <td>{{.resage}}</td>
    <td>{{.resgender}}</td>
    <td>{{.resemail}}</td>
    <td>{{.respassword}}</td>
  </tr>
</table>

</body>
</html>