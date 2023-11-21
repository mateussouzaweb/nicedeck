// Console output
window.addEventListener('load', async () => {

    async function runAndCaptureConsole(callback){
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

    async function fetchConsoleOutput(){
        const result = await request('GET', '/api/console/output')
        writeConsoleOutput(result)
    }

    async function writeConsoleOutput(text){
        const content = $('#console #content')
        content.innerHTML = text.split("\n").join("<br/>")
        content.scrollTop = content.scrollHeight;
    }

    on('#console #clear', 'click', async (event) => {
        event.preventDefault()

        await request('POST', '/api/console/clear')
        await fetchConsoleOutput()
    })

    window.runAndCaptureConsole = runAndCaptureConsole

})