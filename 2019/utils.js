import { readFileSync } from "node:fs"
import { argv } from "node:process"


export function fileContent(filepath){
    let content = readFileSync(filepath);
    let lines = content.toString().split("\n");
    if (lines[lines.length-1].length == 0){
        lines = lines.slice(0, lines.length - 1);
    }
    return lines
}

export function inputFileName(){
    if (argv.length>2){
        return argv[1]
    }
    return `${process.env.AOC_DAY}/input.txt`
}

