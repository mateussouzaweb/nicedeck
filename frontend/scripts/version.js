// Version display
window.addEventListener('load', async () => {

    async function loadVersion(){
        const result = await request('GET', '/api/version')
        const version = $('header #version a')
        version.innerHTML = result
    }
    
    await loadVersion()

})
