// Setup
window.addEventListener('load', () => {

    on('#setup form', 'submit', async (event) => {
        event.preventDefault();

        const form = $('#setup form')
        const data = new FormData(form)
        const button = $('button[type="submit"]', form)

        try {
            button.disabled = true
            window.watchConsoleOutput()
            await request('POST', '/api/setup', data)
        } finally {
            window.stopConsoleOutput()
            button.disabled = false
        }
    })

})