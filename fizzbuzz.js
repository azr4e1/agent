/**
 * FizzBuzz implementation in JavaScript
 * Prints numbers from 1 to 15
 * - For multiples of 3, print "Fizz" instead of the number
 * - For multiples of 5, print "Buzz" instead of the number
 * - For multiples of both 3 and 5, print "FizzBuzz"
 */

function fizzBuzz(start = 1, end = 100) {
    for (let i = start; i <= end; i++) {
        let output = '';
        
        if (i % 3 === 0) {
            output += 'Fizz';
        }
        
        if (i % 5 === 0) {
            output += 'Buzz';
        }
        
        console.log(output || i);
    }
}

// Execute the FizzBuzz function
console.log('Starting FizzBuzz (1 to 15):');
fizzBuzz(1, 15);