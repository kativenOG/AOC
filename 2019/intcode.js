import { supportNegativeIndexes, askForInput } from "./utils.js"
import _ from "lodash"

const Modes = {
    POSITION: Symbol("position"),
    IMMEDIATE: Symbol("immediate"),
}
Object.freeze(Modes)

var integerToMode = {
    "0": Modes.POSITION,
    "1": Modes.IMMEDIATE
};


// Node has the error messages immaginable when forgetting the * in the declaration for a generator function lol 
export default async function* intcodeVM(memory){ 
    let commandLen = 0;
    for (let pointer=0; pointer< memory.length; pointer+=commandLen){
        let modes = memory[pointer].toString().split("");
        
        _.map(_.range(5-modes.length), () => modes = ["0"].concat(...modes))
        let pointerA = supportNegativeIndexes(pointer+1),
            pointerB = supportNegativeIndexes(pointer+2),
            pointerC = supportNegativeIndexes(pointer+3),
            firstParam = (integerToMode[modes[2]]==Modes.POSITION) ? memory[supportNegativeIndexes(memory[pointerA])] : memory[pointerA],
            secondParam = (integerToMode[modes[1]]==Modes.POSITION) ? memory[supportNegativeIndexes(memory[pointerB])] : memory[pointerB],
            thirdParam = (integerToMode[modes[0]]==Modes.POSITION) ? memory[supportNegativeIndexes(pointerC)] : pointerC,
            opCode = parseInt(_.sum(modes.slice(3)));
        switch (opCode) {
            case 1: // ADD
                memory[thirdParam] = firstParam + secondParam;
                commandLen = 4
                break;
            case 2: // MULTIPLY
                memory[thirdParam] = firstParam * secondParam;
                commandLen = 4
                break;
            case 3: // INPUT
                commandLen = 2
                let inputVal = 0
                try{
                    let res = await askForInput("Input: ");
                    inputVal = parseInt(res);
                    break;
                }catch(e){
                    console.error( "intcode VM only supports integers, retry", e);
                }
                memory[memory[pointerA]] = inputVal
                break;
            case 4: // OUTPUT
                commandLen = 2
                yield memory[memory[pointerA]];
                break;
            case 5: // JUMP IF 
                if (firstParam!=0) pointer= secondParam;
                commandLen = (firstParam!=0) ?  0 : 2;
                break;
            case 6: // JUMP IF NOT 
                if (firstParam==0) pointer= secondParam;
                commandLen = (firstParam==0) ?  0 : 2;
                break;
            case 7: // LESS THAN 
                commandLen = 4;
                memory[memory[pointerC]] = (firstParam<secondParam) ? 1 : 0;
                break;
            case 7: // EQUAL
                commandLen = 4;
                memory[memory[pointerC]] = (firstParam==secondParam) ? 1 : 0;
                break;
            case 99: // EXIT
                commandLen = 1
                yield memory[0];
                return;
        }
    }

    throw new Error("Invalid intcode program with no exit command");
}