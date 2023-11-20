// Console output
window.addEventListener('load', async () => {

    let timeout = null

    async function fetchConsoleOutput(){
        const result = await request('GET', '/api/console')
        const content = $('#console #content')
        
        content.innerHTML = result.split("\n").join("<br/>")
        content.scrollTop = content.scrollHeight;

        timeout = window.setTimeout(fetchConsoleOutput, 1000)
    }

    on('#console #clear', 'click', async (event) => {
        event.preventDefault()

        await request('GET', '/api/clear')
        await fetchConsoleOutput()
    })

    await fetchConsoleOutput()

})