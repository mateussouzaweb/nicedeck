// Console output
window.addEventListener('load', async () => {

    /**
     * Run callback and capture console output
     * @param {Boolean} focus
     * @param {Function} callback
     */
    async function runAndCaptureConsole(focus, callback) {
        if (focus) {
            $('#console').scrollIntoView({
                behavior: 'smooth',
                block: 'end',
                inline: 'nearest'
            })
        }

        let interval = null
        try {
            await request('POST', '/api/console/capture')
            interval = window.setInterval(fetchConsoleOutput, 2000)
            await callback()
        } finally {
            window.clearInterval(interval)
            await request('POST', '/api/console/release')
            await fetchConsoleOutput()
        }
    }

    /**
     * Fetch the console output and write it
     */
    async function fetchConsoleOutput() {
        const result = await request('GET', '/api/console/output')
        writeConsoleOutput(result)
    }

    /**
     * Write console output to target location
     * @param {String} text 
     */
    async function writeConsoleOutput(text) {
        const content = $('#console #content')
        content.innerHTML = text.split("\n").join("<br/>")
        content.scrollTop = content.scrollHeight
    }

    on('#console #clear', 'click', async (event) => {
        event.preventDefault()

        await request('POST', '/api/console/clear')
        await fetchConsoleOutput()
    })

    window.runAndCaptureConsole = runAndCaptureConsole

})