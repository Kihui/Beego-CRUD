<!DOCTYPE html>

<html>
<head>
  <title>Beego-CRUD - Crear</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  </head>

<body>
  <h3>Agregar usuario</h3>
  <form role="form" id="usuario" method="POST">
    Nombre：<input name="nombre" value="{{.Usuario.Nombre}}" type="text" />
    Apellido paterno：<input name="app" value="{{.Usuario.App}}" type="text" />
    Apellido materno：<input name="apm" value="{{.Usuario.Apm}}" type="text" />
    Correo electrónico：<input name="correo" value="{{.Usuario.Correo}}" type="text" />
    Edad：<input name="edad" value="{{.Usuario.Edad}}" type="text" />
    <input type="submit" value="Aceptar" />
  </form>
</body>
</html>

