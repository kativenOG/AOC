import { fileContent, inputFileName } from "../utils.js";
import _ from "lodash";

function wireToSetOfPositions(wire) {
  let currentPos = [0, 0],
    positions = new Set(),
    t = 0,
    timestamps = {};

  for (let movement of wire) {
    let direction = movement[0],
      val = parseInt(movement.slice(1).toString()),
      dir = direction == "D" || direction == "L" ? -1 : 1,
      newPos = 0,
      increase = dir * val;

    switch (direction) {
      case "U":
      case "D":
        for (let inner = 1; inner <= val; ++inner) {
          newPos = JSON.stringify([currentPos[0] + inner * dir, currentPos[1]]);
          positions.add(newPos);
          timestamps[newPos] = t + inner;
        }
        currentPos[0] += increase;
        break;
      case "L":
      case "R":
        for (let inner = 1; inner <= val; ++inner) {
          newPos = JSON.stringify([currentPos[0], currentPos[1] + inner * dir]);
          positions.add(newPos);
          timestamps[newPos] = t + inner;
        }
        currentPos[1] += increase;
        break;
    }
    t += val;
  }

  return [positions, timestamps];
}

var parseWire = (wire) => wire.split(",");

function starOne(input, verbose = true) {
  let wireA = wireToSetOfPositions(parseWire(input[0])),
    wireB = wireToSetOfPositions(parseWire(input[1]));
  let ints = [...wireA[0].intersection(wireB[0]).keys()];
  let res = _.min(
    _.map(ints, (point) =>
      _.sum(_.map(JSON.parse(point), (val) => Math.abs(val))),
    ),
  );

  if (verbose) console.log(`STAR ONE: ${res}`);
  return res;
}

function starTwo(input, verbose = true) {
  let parseWire = (wire) => wire.split(",");
  let wireA = wireToSetOfPositions(parseWire(input[0])),
    wireB = wireToSetOfPositions(parseWire(input[1]));
  let ints = [...wireA[0].intersection(wireB[0]).keys()];
  let res = _.min(
    _.map(ints, (intersec) => wireA[1][intersec] + wireB[1][intersec]),
  );

  if (verbose) console.log(`STAR TWO: ${res}`);
  return res;
}

function main() {
  let fileName = inputFileName();

  let input = fileContent(fileName);
  starOne(input);
  starTwo(input);
}

main();
