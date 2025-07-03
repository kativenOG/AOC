import { fileContent, inputFileName } from "../utils.js";
import _ from "lodash";

function recursiveFuelConsuption(inputFuel) {
  let fuelForFuel = Math.floor(inputFuel / 3) - 2;
  if (fuelForFuel <= 0) return 0;
  return fuelForFuel + recursiveFuelConsuption(fuelForFuel);
}

function starOne(input, verbose = true) {
  let res = _.sum(
    _.map(input, (shipModule) => {
      return parseInt(Math.floor(shipModule / 3) - 2);
    }),
  );
  if (verbose) console.log(`STAR ONE: ${res}`);
  return res;
}

function starTwo(input, verbose = true) {
  let res = _.sum(
    _.map(input, (shipModule) => {
      return recursiveFuelConsuption(shipModule);
    }),
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
