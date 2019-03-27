import('path').then(obj => {
    console.log("success")
}).catch(error => {
    console.log("error")
})
const afunc = () => {
    console.log("ld")
}


afunc()
