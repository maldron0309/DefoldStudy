components {
  id: "bullet"
  component: "/script/bullet.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bulletAnim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/animation/bullet.tilesource\"\n"
  "}\n"
  ""
  position {
    y: 2.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"bullets\"\n"
  "mask: \"enemies\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -1.0\n"
  "      y: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.34202015\n"
  "      w: 0.9396926\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 2.7157724\n"
  "  data: 10.374827\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
