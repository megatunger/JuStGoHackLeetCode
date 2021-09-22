function jumpingOnClouds(c) {
    let steps = 0;
    let i=0;
    while (i<=c.length) {
        if (i+2 < c.length && c[i+2] === 0) {
            steps = steps + 1
            i = i+2
            console.log(steps, i)
            continue
        } else {
            if (c[i+1] === 0) {
                steps = steps + 1
                i = i+1
                console.log(steps, i)
                continue
            }
        }
        i++
    }
    return steps
}

console.log('result: ', jumpingOnClouds([0,0,1,0,0,1,0]))
console.log('result: ', jumpingOnClouds([0,0,0,0,1,0]))