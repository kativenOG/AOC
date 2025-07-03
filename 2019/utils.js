import { readFileSync } from "node:fs";
import { argv } from "node:process";
import readline from "node:readline";
import { stdin, stdout } from "node:process";

export var supportNegativeIndexes = (val, arr) =>
  (val > arr.length) ? val%arr.length : (val < 0) ? arr.length - val : val;

export async function askForInput(questionText) {
  console.log(questionText);
  let rl = readline.createInterface(stdin, stdout),
    res = await new Promise((resolve) => rl.question("", resolve));
  rl.close();
  return res;
}
export function fileContent(filepath) {
  let content = readFileSync(filepath);
  let lines = content.toString().split("\n");
  if (lines[lines.length - 1].length == 0) {
    lines = lines.slice(0, lines.length - 1);
  }
  return lines;
}

export function inputFileName() {
  if (argv.length > 2) {
    return argv[1];
  }
  return `${process.env.AOC_DAY}/input.txt`;
}
