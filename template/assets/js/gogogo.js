var host = window.location.origin;
//console.log(host);
var socket = io.connect();
var localWorld = "";
var moveTo = "";
var GameOver = false;

//PlayerType , tipo Jugador
var PlayerType = 1;
//SentraPointType , tipo Moneda Sentra
var SentraPointType = 2;
//EnemyType , tipo Enemigo en el juego
var EnemyType = 3;

function PlayerMove(keyCodeNumber) {
  (LEFT = 37), (UP = 38), (RIGHT = 39), (DOWN = 40);
  switch (keyCodeNumber) {
    case LEFT:
      moveTo = "Left";
      break;
    case UP:
      moveTo = "Up";
      break;
    case RIGHT:
      moveTo = "Right";
      break;
    case DOWN:
      moveTo = "Down";
      break;

    default:
      moveTo = "";
      break;
  }
  var player = {
    name: $("#player").val(),
    move: moveTo
  };
  var playerJson = JSON.stringify(player);

  socket.emit("move", playerJson, function(callback) {
    console.log("callback : " + callback);
    socket.emit("updateWorld", "");
  });
}

function renderObject(object) {
  var sprite = moveTo == "" ? "Right" : moveTo;
  var htmlPlayer = "<center>";
  //Coordenadas Player

  if (object.Type == PlayerType) {
    if (object.Who == $("#player").val()) {
      htmlPlayer += `<smal><strong style="color:#00F">${
        object.Who
      }</strong></smal>`;
      htmlPlayer += `<br/><img width="32px" src="/assets/img/player${sprite}.gif"/>`;
    } else {
      htmlPlayer += `<smal>${object.Who}</smal>`;
      htmlPlayer += `<br/><img width="32px" src="/assets/img/playerRight.gif"/>`;
    }
  }

  if (object.Type == EnemyType) {
    htmlPlayer += `<smal><strong style="color:#F00">${
      object.Who
    }</strong></smal>`;
    htmlPlayer += `<br/><img width="32px" src="/assets/img/enemy.gif"/>`;
    if ($("#player").val() == object.Who && !GameOver) {
      GameOver = true;
      document.onkeydown = null;
      alert("GAME OVER!!!");
    }
  }
  htmlPlayer += "</center>";
  $(`.pos_${object.X}_${object.Y}`).html(htmlPlayer);
}

socket.on("renderWorld", function(remoteWorld) {
  if (localWorld != remoteWorld) {
    localWorld = remoteWorld;
    var mundo = JSON.parse(remoteWorld);
    console.log(remoteWorld);
    $(".tg-0lax").html("");
    for (var p in mundo.Objects) {
      var player = mundo.Objects[p];
      renderObject(player);
    }
  }
});

function checkKeycode(event) {
  var keyDownEvent = event || window.event,
    keycode = keyDownEvent.which ? keyDownEvent.which : keyDownEvent.keyCode;
  PlayerMove(keycode);
  return false;
}

function createWorld() {
  document.onkeydown = checkKeycode;
  var totalX = 10;
  var totalY = 20;
  var html = "";
  for (var x = 1; x <= totalX; x++) {
    html += "<tr>";
    for (var y = 0; y <= totalY; y++) {
      html += `<td class="tg-0lax pos_${x}_${y}"></td>`;
    }
    html += "</tr>";
  }
  document.getElementById("world").innerHTML = html;
  ping();
}

function ping() {
  console.log("PING!");
  setTimeout(() => {
    socket.emit("updateWorld", localWorld);
    ping();
  }, 300);
}
document.onkeydown = checkKeycode;
