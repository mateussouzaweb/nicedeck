// Version display
window.addEventListener('load', async () => {

    /**
     * Load and display current software version
     */
    async function loadVersion() {
        const result = await request('GET', '/api/version')
        const version = $('#version a')
        version.innerHTML = `version: ${result.trim()}`
    }

    await loadVersion()

})
