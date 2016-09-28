<!DOCTYPE html>

<html>
<head>
  <title>Beego-CRUD</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  </head>

<body>
  
  <h3>Actualizar usuario</h3>
  <form role="form" id="usuario" method="POST">
    Nombre：<input name="nombre" value="{{.Nom}}" type="text" />
    Apellido paterno：<input name="app" value="{{.App}}" type="text" />
    Apellido materno：<input name="apm" value="{{.Apm}}" type="text" />
    Correo electrónico：<input name="correo" value="{{.Correo}}" type="text" />
    Edad：<input name="edad" value="{{.Edad}}" type="text" />
    <input type="submit" value="Aceptar" />
  </form>
</body>
</html>


