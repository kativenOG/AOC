
export default function intcodeVM(memory){
    let commandLen = 0
    for (let pointer=0; pointer< memory.length; pointer+=commandLen){
        switch (memory[pointer]) {
            case 1: // ADD
                memory[memory[pointer+3]] = memory[memory[pointer+1]] + memory[memory[pointer+2]];
                commandLen = 4
                break;
            case 2: // MULTIPLY
                memory[memory[pointer+3]] = memory[memory[pointer+1]] * memory[memory[pointer+2]];
                commandLen = 4
                break;
            case 99: // EXIT
                commandLen = 1
                return memory[0];
        }
    }

    throw new Error("Invalid intcode program with no exit command");
}
 