// Version display
window.addEventListener('load', async () => {

    /**
     * Load and display current software version
     */
    async function loadVersion() {
        const result = await request('GET', '/api/version')
        const version = $('header #version a')
        version.innerHTML = result
    }

    await loadVersion()

})
