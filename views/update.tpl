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
 <div class="principal2">
    <h4>Actualizar Usuario</h4>
  <form role="form" id="usuario" method="POST">
    <div class="mdl-textfield mdl-js-textfield">
    <input id="nombre" class="mdl-textfield__input" name="nombre" value="{{.Nom}}" type="text" />
    <label class="mdl-textfield__label" for="nombre">Nombre</label>
    </div>
    <br/>
    <div class="mdl-textfield mdl-js-textfield">
    <input id="app" class="mdl-textfield__input"name="app" value="{{.App}}" type="text" />
    <label class="mdl-textfield__label" for="app">Apellido Paterno</label>
    </div>
    <br/>
    <div class="mdl-textfield mdl-js-textfield">
    <input id="apm" class="mdl-textfield__input" name="apm" value="{{.Apm}}" type="text" />
    <label class="mdl-textfield__label" for="apm">Apellido Materno</label>
    </div>
    <br/>
    <div class="mdl-textfield mdl-js-textfield">
    <input id="correo" class="mdl-textfield__input" name="correo" pattern="[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$" value="{{.Correo}}" type="text" />
    <label class="mdl-textfield__label" for="correo">Correo Electrónico</label>
    <span class="mdl-textfield__error">Ingresa un correo válido</span>
    </div>
    <br/>
    <div class="mdl-textfield mdl-js-textfield">
    <input id="edad" class="mdl-textfield__input" pattern="[0-9]*?" name="edad" value="{{.Edad}}" type="text" />
    <label class="mdl-textfield__label" for="edad">Edad</label>
    <span class="mdl-textfield__error">Sólo se aceptan números enteros positivos</span>
    </div>
    <br/>
    <input class="mdl-button mdl-js-button mdl-button--raised" type="submit" value="Aceptar" />
  </form>
  </div>
</body>
</html>


