function softmax(vec) {
    sum_val = vec.reduce((acc, val) => {
        return acc + (Math.E ** val)
    })
    ret = vec.map((val) => {
        return (Math.E ** val) / sum_val;
    })

    return ret;
}