import {fileContent, inputFileName} from "../utils.js" 
import intcodeVM from "../intcode.js";
import _ from "lodash" 

   
function starOne(input, verbose=true){
    let intcode = _.map(input[0].toString().split(","), (val)=>parseInt(val))
    let mappings = [[1,12], [2,2]];
    for (let pair of mappings){
        intcode[pair[0]] = pair[1];
    }
    let res = intcodeVM(intcode)
    if (verbose) console.log(`STAR ONE: ${res}`);
    return res
}

function testSolution(originalMemory, noun, verb){
    let intcode = originalMemory.slice();
    intcode[1]= noun;
    intcode[2]= verb;
    return intcodeVM(intcode)
}


function starTwo(input, verbose=true){
    let originalMemory = _.map(input[0].toString().split(","), (val)=>parseInt(val))
    let target = 19690720;

    for (let noun=0; noun<100; ++noun){
        for (let verb=0; verb<100; ++verb){
            let res = testSolution(originalMemory, noun, verb);
            if (res === target){
                res = 100*noun + verb;
                if (verbose) console.log(`STAR TWO: ${res}`);
                return res;
            } 
        }
    }

    throw new Error("no combination of noun and verb can solve this problem");

}

function main() {
    let fileName = inputFileName();

    let input = fileContent(fileName);
    starOne(input) 
    starTwo(input);
} 

main()