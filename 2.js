import fs from 'fs'

const input = fs.readFileSync('input2.txt', {encoding: 'utf8'})

const points = {
    X: 1,
    Y: 2,
    Z: 3
}

const choiceMap = {
    X: 'A',
    Y: 'B',
    Z: 'C'
}
const loseMap = {
    X: 'B',
    Y: 'C',
    Z: 'A'
}
const winMap = {
    X: 'C',
    Y: 'A',
    Z: 'B'
}

const POINTS_FOR_DRAW = 3
const POINTS_FOR_WIN = 6

//part 2
const reverseObject = (obj) => Object.entries(obj).reduce((acc, [key, val]) => ({...acc, [val]: key}), {})

const enemyLoseMap = reverseObject(winMap)
const enemyChoiceMap = reverseObject(choiceMap)
const enemyWinMap = reverseObject(loseMap)

//part 2 function
const choiceChange = (myChoice, enemyChoice) => {
    switch(myChoice){
        case('X'):
            return enemyWinMap[enemyChoice]
        case('Y'):
            return enemyChoiceMap[enemyChoice]
        case('Z'):
            return enemyLoseMap[enemyChoice]
    }
}

const isRoundWon = (myChoice, enemyChoice) => {
    return loseMap[myChoice] !== enemyChoice
}

const pointsForRound = (choice) => {
    const choiceArray = choice.split(' ')
    const [enemyChoice, myBaseChoice] = choiceArray

    const myChoice = choiceChange(myBaseChoice, enemyChoice)

    const pointsForSelection = points[myChoice]
    if(choiceMap[myChoice] === enemyChoice)
        return pointsForSelection + POINTS_FOR_DRAW
    const result = isRoundWon(myChoice, enemyChoice) ? pointsForSelection + POINTS_FOR_WIN : pointsForSelection
    return result
}


const splitted = input.split('\n')

const sumOfPoints = splitted.reduce((acc, item) => acc + pointsForRound(item), 0)

console.log(sumOfPoints)

