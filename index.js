
//task 3

// import fs from 'fs'

// //part 1
// //const input = fs.readFileSync('input3.txt', {encoding: 'utf8'}).split('\n')
// //part 2
// const input = fs.readFileSync('input3.txt', {encoding: 'utf8'}).split('\n')

// let dataset = []

// for(let i = 0; i <= input.length; i+=3){
//     if(!input[i] || !input[i + 1] || !input[i + 2]) break
//     dataset = [...dataset, [input[i], input[i + 1], input[i + 2]]]
// }

// console.log(dataset)

// ////

// const getCommonItem = (data) => {
//     const [firstCompartment, secondCompartment] = [data.slice(0, data.length/2).split(''), data.slice(data.length/2, data.length).split('')]
//     const duplicate = firstCompartment.reduce((acc, item) => secondCompartment.includes(item) ? item : acc, '')
//     return duplicate

// }

// const getCommonItemBetweenGroup = (data) => {
//     console.log(data)
//     const [first, second, third] = data

//     const intersection = first.split('').filter(item => second.split('').includes(item)).filter(item => third.split('').includes(item)) 

//     return intersection[0]

// }

// const evaluateItem = (item) => {
//     const UPPERCASE_TOP = 90
//     const UPPERCASE_SUBSTRACT_VALUE = 64
//     const LOWERCASE_SUBSTRACT_VALUE = 96
//     const UPPERCASE_OFFSET = 26

//     const value = item.charCodeAt(0) > UPPERCASE_TOP ? (item.charCodeAt(0) - LOWERCASE_SUBSTRACT_VALUE) : (item.charCodeAt(0) - (UPPERCASE_SUBSTRACT_VALUE) + UPPERCASE_OFFSET)

//     return value 
// }

// //const result = input.reduce((acc, item) => acc + evaluateItem(getCommonItem(item)), 0)

// //evaluateItem(getCommonItem('tlBCxFgZFNpGClFFZGGtFBZrdZQznfdQQrRWZRQdRsVf'))


// const result = dataset.reduce((acc, item) => acc + evaluateItem(getCommonItemBetweenGroup(item)), 0)

// console.log(result)

//task 4