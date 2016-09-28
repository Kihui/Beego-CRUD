<!DOCTYPE html>

<html>
<head>
  <title>Beego-CRUD</title>
  <link rel="stylesheet" href="/static/css/material.css">
  <link rel="stylesheet" href="/static/css/main.css">    
  <script src="/static/js/material.min.js"></script>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="https://fonts.googleapis.com/icon?family=Material+Icons"
      rel="stylesheet">
  </head>

<body>

  <div id="header">
    <div class="mdl-grid">
      <div class="mdl-cell mdl-cell--6-col">
	<h2>Kihui-DEV</h2>
	</div>
	<div class="mdl-cell mdl-cell--6-col">
	  <h4 class="beego">CRUD de Beego</h4>
	  </div>
    </div>
  </div>
  
  <div class="principal mdl-grid">

    <div class="mdl-cell mdl-cell--8-col">
      <h4>Usuarios</h4>
  <table class="mdl-data-table mdl-js-data-table">
      <thead>
        <tr>
          <th class="mdl-data-table__cell--non-numeric">Nombre</th>
          <th>Correo</th>
	  <th>Edad</th>
	  <th></th>
	  <th></th>
        </tr>
      </thead>
      <tbody>
        {{range $record := .records}}
        <tr>
          <td class="mdl-data-table__cell--non-numeric">{{$record.Nombre}} {{$record.App}} {{$record.Apm}}</td>
	  <td class="mdl-data-table__cell--non-numeric">{{$record.Correo}}</td>
	  <td class="mdl-data-table__cell--non-numeric">{{$record.Edad}}</td>
          <td class="mdl-data-table__cell--non-numeric"><a href="/delete?id={{$record.Id}}"><i class="material-icons">delete_forever</i></a></td>
	  <td class="mdl-data-table__cell--non-numeric"><a href="/update?id={{$record.Id}}"><i class="material-icons">edit</i></a></td>
        </tr>
        {{end}}
      </tbody>
  </table>
  </div>

  <div class="mdl-cell mdl-cell--4-col">
    <h4>Agregar Usuario</h4>
  <form role="form" id="usuario" method="POST">
    <div class="mdl-textfield mdl-js-textfield">
    <input id="nombre" class="mdl-textfield__input" name="nombre" value="{{.Usuario.Nombre}}" type="text" />
    <label class="mdl-textfield__label" for="nombre">Nombre</label>
    </div>
    <br/>
    <div class="mdl-textfield mdl-js-textfield">
    <input id="app" class="mdl-textfield__input"name="app" value="{{.Usuario.App}}" type="text" />
    <label class="mdl-textfield__label" for="app">Apellido Paterno</label>
    </div>
    <br/>
    <div class="mdl-textfield mdl-js-textfield">
    <input id="apm" class="mdl-textfield__input" name="apm" value="{{.Usuario.Apm}}" type="text" />
    <label class="mdl-textfield__label" for="apm">Apellido Materno</label>
    </div>
    <br/>
    <div class="mdl-textfield mdl-js-textfield">
    <input id="correo" class="mdl-textfield__input" name="correo" pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$" value="{{.Usuario.Correo}}" type="text" />
    <label class="mdl-textfield__label" for="correo">Correo Electrónico</label>
    <span class="mdl-textfield__error">Ingresa un correo válido</span>
    </div>
    <br/>
    <div class="mdl-textfield mdl-js-textfield">
    <input id="edad" class="mdl-textfield__input" pattern="[0-9]*?" name="edad" value="{{.Usuario.Edad}}" type="text" />
    <label class="mdl-textfield__label" for="edad">Edad</label>
    <span class="mdl-textfield__error">Sólo se aceptan números enteros positivos</span>
    </div>
    <br/>
    <input class="mdl-button mdl-js-button mdl-button--raised" type="submit" value="Aceptar" />
  </form>
  </div>
</body>
</html>


