import { fileContent, inputFileName } from "../utils.js";

function criteria(psw) {
  let prev = 0,
    foundValid = false,
    strPsw = psw.toString().split("");
  for (var digit of strPsw) {
    digit = parseInt(digit);
    if (prev > digit) {
      foundValid = false;
      // Maybe here increase psw depending on the digit :)
      break;
    }
    if (prev == digit) foundValid = true;
    prev = digit;
  }
  return foundValid;
}

function starOne(input, verbose = true) {
  let res = 0,
    range = input[0].split("-");
  for (var psw = parseInt(range[0]); psw <= parseInt(range[1]); psw++)
    if (criteria(psw)) res++;
  if (verbose) console.log(`STAR ONE: ${res}`);
  return res;
}

function stricterCriteria(psw) {
  let prev = 0,
    matchingCounter = 0,
    foundValid = false,
    strPsw = psw.toString().split("");
  for (var digit of strPsw) {
    digit = parseInt(digit);
    if (prev > digit) {
      matchingCounter = 0;
      foundValid = false;
      break;
    } else if (!foundValid) {
      if (prev == digit) {
        matchingCounter++;
      } else {
        foundValid = matchingCounter == 1 ? true : false;
        matchingCounter = 0;
      }
    }
    prev = digit;
  }
  return foundValid || matchingCounter == 1 ? true : false;
}

function starTwo(input, verbose = true) {
  let res = 0,
    range = input[0].split("-");
  for (var psw = parseInt(range[0]); psw <= parseInt(range[1]); psw++)
    if (stricterCriteria(psw)) res++;
  if (verbose) console.log(`STAR ONE: ${res}`);
  return res;
}

function main() {
  let fileName = inputFileName();
  let input = fileContent(fileName);

  console.log(
    stricterCriteria(112233),
    stricterCriteria(123444),
    stricterCriteria(111122),
  );
  starOne(input);
  starTwo(input);
}

main();
