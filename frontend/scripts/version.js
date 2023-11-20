// Version display
window.addEventListener('load', async () => {

    const result = await request('GET', '/api/version')
    const version = $('header #version')
    version.innerHTML = result
    
})
