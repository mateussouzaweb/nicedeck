// Console output
window.addEventListener('load', async () => {

    let timeout = null

    async function fetchConsoleOutput(){
        const result = await request('GET', '/api/console')
        writeConsoleOutput(result)
    }

    async function writeConsoleOutput(text){
        const content = $('#console #content')
        content.innerHTML = text.split("\n").join("<br/>")
        content.scrollTop = content.scrollHeight;
    }

    async function watchConsoleOutput(){
        await fetchConsoleOutput()
        timeout = window.setTimeout(watchConsoleOutput, 2000)
    }

    async function stopConsoleOutput() {
        window.clearTimeout(timeout)
    }

    on('#console #clear', 'click', async (event) => {
        event.preventDefault()

        await request('GET', '/api/clear')
        await fetchConsoleOutput()
    })

    window.fetchConsoleOutput = fetchConsoleOutput
    window.writeConsoleOutput = writeConsoleOutput
    window.watchConsoleOutput = watchConsoleOutput
    window.stopConsoleOutput = stopConsoleOutput

})