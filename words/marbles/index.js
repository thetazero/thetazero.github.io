const n = 15


let Pivots = []

function makePivot(x, y) {
  let width = 50
  let height = 10
  let pivot = Bodies.fromVertices(x, y, [
    { x: -width, y: -height / 2 },
    { x: -width, y: height / 2 },
    { x: -20, y: height / 2 },
    { x: 0, y: 30 },
    { x: 20, y: height / 2 },
    { x: width, y: height / 2 },
    { x: width, y: -height / 2 },
  ], {
    collisionFilter: { group },
    render: {
      fillStyle: "hsl(300, 0%, 20%)",
    }
  });
  Matter.Body.rotate(pivot, Math.PI - 0.25)
  let padding = 50
  let bottom = Bodies.rectangle(x, y + 30, width * 2, 10, { isStatic: true })
  let left = Bodies.rectangle(x - width - padding, y, 10, width * 2, { isStatic: true })
  let right = Bodies.rectangle(x + width + padding, y, 10, width * 2, { isStatic: true })
  let con = Constraint.create({
    bodyA: pivot,
    // pointB: Vector.clone(pivot.position),
    pointB: { x, y },
    stiffness: 1,
    length: 0
  })
  Pivots.push(pivot)
  return [pivot, con, bottom, left, right]
}

function dropBall(offset = 0) {
  let circle = Bodies.circle(x + 1 + offset * dx, y - 180 + offset * dy, 20, {
    render: {
      fillStyle: "#00cccc",
    },
  })
  World.add(world, circle,)
  let int = setInterval(() => {
    // console.log(circle.position.y)
    if (circle.position.y > 1800) {
      Matter.Composite.remove(world, circle)
    }
  }, 100)
}

const numberElem = document.querySelector("#number")
const cvalElem = document.querySelector("#cval")

function dropNumber() {
  let num = - (-numberElem.value)
  if (num > 2 ** n) {
    alert(`pick a number between 1 and ${2 ** (n - 1)}`)
    return
  }
  let cn = n - 1
  let int = setInterval(() => {
    console.log(2 ** cn, num)
    for (; 2 ** cn > num; cn--) {
      console.log(cn)
      if (cn < 0) {
        break
      }
    }
    dropBall(cn)
    num -= 2 ** cn
    cn--
    if (num <= 0) {
      clearInterval(int)
    }
  }, 100);
}


var Engine = Matter.Engine,
  Render = Matter.Render,
  Runner = Matter.Runner,
  Composites = Matter.Composites,
  Constraint = Matter.Constraint,
  MouseConstraint = Matter.MouseConstraint,
  Mouse = Matter.Mouse,
  World = Matter.World,
  Bodies = Matter.Bodies,
  Body = Matter.Body,
  Vector = Matter.Vector;

// create engine
var engine = Engine.create(),
  world = engine.world;

// create renderer
let simContainer = document.querySelector('.simulation')
var render = Render.create({
  element: simContainer,
  engine: engine,
  options: {
    width: simContainer.offsetWidth,
    height: simContainer.offsetHeight,
    wireframes: false,
    // showAngleIndicator: true,
    // showCollisions: true,
    // showVelocity: true
  }
});

Render.run(render);

// create runner
var runner = Runner.create();
Runner.run(runner, engine);

// add bodies
var group = Body.nextGroup(true);



World.add(world, [
  // Bodies.rectangle(400*3, 600*3, 800*3, 50.5, { isStatic: true }),
]);

// add mouse control
var mouse = Mouse.create(render.canvas),
  mouseConstraint = MouseConstraint.create(engine, {
    mouse: mouse,
    constraint: {
      stiffness: 0.2,
      render: {
        visible: false
      }
    }
  });

World.add(world, mouseConstraint);

// keep the mouse in sync with rendering
render.mouse = mouse;

// fit the render viewport to the scene
Render.lookAt(render, {
  min: { x: 0, y: 0 },
  max: { x: 800 * 3, y: 600 * 3 }
});


let dx = -75
let dy = 110
let x = 2300
let y = 100

for (let i = 0; i < n; i++) {
  World.add(world, makePivot(x + i * dx, y + i * dy))
}

Matter.Events.on(runner, "beforeTick", () => {
  let sum = 0
  Pivots.forEach((p, i) => {
    if (p.angle > Math.PI) {
      p.parts.forEach(part => {
        part.render.fillStyle = "hsl(300, 100%, 20%)"
      });
      sum += 2 ** i
    } else {
      p.parts.forEach(part => {
        part.render.fillStyle = "hsl(300, 0%, 20%)"
      });
    }
  })
  cvalElem.innerText = sum
})
