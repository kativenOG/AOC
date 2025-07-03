import { fileContent, inputFileName } from "../utils.js";
import intcodeVM from "../intcode.js";
import _ from "lodash";

async function star(input, verbose = true, star="one") {
  let intcode = _.map(input[0].toString().split(","), (val) => parseInt(val));
  // You have to give it 1 as input:
  var res = [];
  if (verbose) console.log(`STAR ${star.toUpperCase()}:`);
  for await (let val of intcodeVM(intcode)) {
    if (verbose) console.log(val);
    res.push(val);
  }
  return res;
}

async function main() {
  // input 0 -> output 0 otherwhise 1 
  // console.log(await intcodeVM([3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9]).next()) 

  // input 0 -> output 0 otherwhise 1 
  // console.log(await intcodeVM([3,3,1105,-1,9,1101,0,0,12,4,12,99,1]).next()) 

  // 999 if input below 8 
  // 1000 if input equals 8 
  // 1001 if input greater than 8 
  // console.log(await intcodeVM([3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
  // 1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
  // 999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99]).next()) 

  let fileName = inputFileName();
  let input = fileContent(fileName);
  //  await star(input); // input 1 
  await star(input, star="two"); // input 5 
}

await main();
