import { fileContent, inputFileName } from "../utils.js";
import intcodeVM from "../intcode.js";
import _ from "lodash";

async function starOne(input, verbose = true) {
  let intcode = _.map(input[0].toString().split(","), (val) => parseInt(val));
  // You have to give it 1 as input:
  var res = [];
  if (verbose) console.log(`STAR ONE:`);
  for await (let val of intcodeVM(intcode)) {
    if (verbose) console.log(val);
    res.push(val);
  }
  return res;
}

async function main() {
  let fileName = inputFileName();
  let input = fileContent(fileName);
  await starOne(input);
}

await main();
