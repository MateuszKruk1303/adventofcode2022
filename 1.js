import fs from 'fs'

const input = fs.readFileSync('input.txt', {encoding: 'utf8'})

const sums = input.split('\n\n').map((item) => {
    return item.split('\n').length <= 1 ? 
    parseInt(item) : 
    item.split('\n').reduce((acc, value) => parseInt(acc) + parseInt(value))
})

const findMax = (arr) => arr.reduce((acc, item) => acc >= item ? acc : item)

const find3Max = (arr, highestValues = []) => {
    if(highestValues.length === 3) return highestValues
    const highestValue = findMax(arr)
    const index = arr.indexOf(highestValue)
    const newArray = [...arr.slice(0, index), ...arr.slice(index + 1, arr.length)]

    return find3Max(newArray, [...highestValues, highestValue])
}

const sum = (arr) => arr.reduce((acc, item) => acc + item)

console.log(sum(find3Max(sums)))
