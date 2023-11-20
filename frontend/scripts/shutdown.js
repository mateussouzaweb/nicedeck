// Shutdown support
window.addEventListener('load', async () => {

    on('#shutdown', 'click', async (event) => {
        event.preventDefault()
        await request('POST', '/app/shutdown')

        window.setTimeout(() =>{
            window.close()
        }, 100) 
    })

})

// document.addEventListener('visibilitychange', function() {
//     if( document.visibilityState == 'hidden' ){
//         navigator.sendBeacon('/app/shutdown', { method: 'POST' })
//     }
// })